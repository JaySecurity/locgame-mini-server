package blockchain

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"
	"time"

	"locgame-mini-server/internal/blockchain/contracts"
	"locgame-mini-server/internal/config"
	storeDto "locgame-mini-server/pkg/dto/store"
	"locgame-mini-server/pkg/log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/onrik/ethrpc"
)

type Blockchain struct {
	nftContract      *contracts.LOCGamePlayNFT
	locgBaseContract *contracts.LOCGBridged
	config           *config.Blockchain

	PolygonClient  *ethclient.Client
	EthereumClient *ethclient.Client
	BaseClient     *ethclient.Client
}

func Connect(config *config.Blockchain) (*Blockchain, error) {
	polygonClient, err := ethclient.Dial("https://" + config.RpcAddresses.Polygon)
	if err != nil {
		return nil, err
	}

	ethereumClient, err := ethclient.Dial("https://" + config.RpcAddresses.Ethereum)
	if err != nil {
		return nil, err
	}

	baseClient, err := ethclient.Dial("https://" + config.RpcAddresses.Base)
	if err != nil {
		return nil, err
	}

	nftContract, err := contracts.NewLOCGamePlayNFT(common.HexToAddress(config.Contracts.NFT), polygonClient)
	if err != nil {
		return nil, err
	}

	locgBaseContract, err := contracts.NewLOCGBridged(common.HexToAddress(config.Contracts.BaseLOCG), baseClient)
	if err != nil {
		return nil, err
	}

	return &Blockchain{
		config:           config,
		nftContract:      nftContract,
		locgBaseContract: locgBaseContract,
		PolygonClient:    polygonClient,
		EthereumClient:   ethereumClient,
		BaseClient:       baseClient,
	}, nil
}

func (b *Blockchain) getAccount(ctx context.Context, privKey string, client *ethclient.Client) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		log.Fatal("Invalid private key. Error:", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("Invalid private key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(2000000) // in units

	auth.Context = ctx

	return auth
}

func (b *Blockchain) NftMint(ctx context.Context, address string, tokensInfo []*storeDto.TokenInfo) (string, error) {
	var tokens []*big.Int
	for _, token := range tokensInfo {
		if token.Status == storeDto.TokenStatus_TokenWaitingForMint {
			id := new(big.Int)
			id.SetString(token.Token, 10)
			tokens = append(tokens, id)
		}
	}
	if len(tokens) == 0 {
		return "", nil
	}

	account := b.getAccount(ctx, b.config.MinterPrivateKey, b.PolygonClient)
	account.GasLimit = uint64(200000 * len(tokensInfo))

	var (
		transaction *types.Transaction
		err         error
	)

	if len(tokens) == 1 {
		transaction, err = b.nftContract.Mint(account, common.HexToAddress(address), tokens[0])
	} else {
		transaction, err = b.nftContract.MintBatch(account, common.HexToAddress(address), tokens)
	}

	if err != nil {
		return "", err
	}
	return transaction.Hash().Hex(), err
}

func (b *Blockchain) GetTransactionResult(ctx context.Context, client *ethclient.Client, transaction common.Hash, attempt int) (bool, error) {
	if attempt > 120 {
		return false, ethereum.NotFound
	}
	receipt, err := client.TransactionReceipt(ctx, transaction)
	if err != nil {
		time.Sleep(3000 * time.Millisecond)
		return b.GetTransactionResult(ctx, client, transaction, attempt+1)
	}
	if err != nil {
		return false, err
	}
	return receipt.Status == 1, nil
}

var ZeroAddress = "0x0000000000000000000000000000000000000000"

func (b *Blockchain) GetTokenOwner(tokenID *big.Int) (string, error) {
	return Retry[string](context.Background(), func() (string, error) {
		addr, err := b.nftContract.OwnerOf(&bind.CallOpts{}, tokenID)
		if err != nil {
			if strings.Contains(err.Error(), "owner query for nonexistent token") {
				return ZeroAddress, nil
			}
			return "", err
		}
		return addr.Hex(), nil
	})
}

func (b *Blockchain) GetTokens(address string) ([]*big.Int, error) {
	return Retry[[]*big.Int](context.Background(), func() ([]*big.Int, error) {
		return b.nftContract.TokensOfOwner(&bind.CallOpts{}, common.HexToAddress(address))
	})
}

func (b *Blockchain) WithdrawAndBurn(ctx context.Context, wallet common.Address, amount *big.Int, burnAmount *big.Int) (string, error) {
	account := b.getAccount(ctx, b.config.MinterPrivateKey, b.BaseClient)
	transfer, err := b.locgBaseContract.Transfer(account, wallet, amount)
	if err != nil {
		return "", err
	}
	nonce := account.Nonce

	account = b.getAccount(ctx, b.config.MinterPrivateKey, b.BaseClient)
	if account.Nonce.Cmp(nonce) == 0 {
		account.Nonce.Add(account.Nonce, big.NewInt(1))
	}
	_, err = b.locgBaseContract.Burn(account, burnAmount)
	if err != nil {
		log.Error("Burn failed:", err)
	}

	return transfer.Hash().String(), nil
}

func (b *Blockchain) WatchNativeEthereumTransactions(tokenName string, rpcAddress string, recipient common.Address, isReconnect bool, onTransferReceived func(tx *contracts.ERC20Transfer)) {
	client, err := ethclient.Dial(`wss://` + rpcAddress)
	if err != nil {
		if !isReconnect {
			log.Warning("first")
			log.Warning(err)
		}
		return
	}
	defer client.Close()

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		if !isReconnect {
			log.Warning("second")
			log.Warning(err)
		}
		return
	}
	defer sub.Unsubscribe()

	if isReconnect {
		log.Warning("The monitoring of native token transactions (" + tokenName + ") in the blockchain has been restored...")
	} else {
		log.Info("The monitoring of native token transactions (" + tokenName + ") in the blockchain has been started...")
	}

	for {
		select {

		case err := <-sub.Err():
			if err != nil {
				if !isReconnect {
					log.Warning("third")
					log.Warning(err)
				}
				return
			}
		case header := <-headers:
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Errorf("%s - Block By Hash Error: %v\n", tokenName, err)
				log.Debugf("Block number: %v\nBlock Hash: %v", header.Number.String(), header.Hash().Hex())
				continue
			}
			for _, tx := range block.Transactions() {
				// Get the signer for the transaction
				signer := types.LatestSignerForChainID(tx.ChainId())
				// Derive the sender's address
				sender, err := types.Sender(signer, tx)
				if err != nil {
					log.Fatalf("Failed to get sender from transaction: %v", err)
				}
				if tx.To() != nil && *tx.To() == recipient {
					log.Debugf("Native token transaction detected:\nTx hash: %s\n", tx.Hash().Hex())

					if tx.Value().Cmp(big.NewInt(0)) > 0 {
						transfer := &contracts.ERC20Transfer{
							From:  sender,
							To:    *tx.To(),
							Value: tx.Value(),
							Raw:   types.Log{TxHash: tx.Hash()},
						}
						go onTransferReceived(transfer)
					}

				}
			}
		}
	}
}

func (b *Blockchain) WatchNativeBaseTransactions(tokenName string, rpcAddress string, recipient common.Address, isReconnect bool, onTransferReceived func(tx *contracts.ERC20Transfer)) {
	wsclient, err := ethclient.Dial(`wss://` + rpcAddress)
	if err != nil {
		if !isReconnect {
			log.Warning("first")
			log.Warning(err)
		}
		return
	}
	defer wsclient.Close()

	headers := make(chan *types.Header)
	sub, err := wsclient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		if !isReconnect {
			log.Warning("second")
			log.Warning(err)
		}
		return
	}
	defer sub.Unsubscribe()

	if isReconnect {
		log.Warning("The monitoring of native token transactions (" + tokenName + ") in the blockchain has been restored...")
	} else {
		log.Info("The monitoring of native token transactions (" + tokenName + ") in the blockchain has been started...")
	}

	for {
		select {

		case err := <-sub.Err():
			if err != nil {
				if !isReconnect {
					log.Warning("third")
					log.Warning(err)
				}
				return
			}
		case header := <-headers:
			client := ethrpc.New(`https://` + rpcAddress)
			block, err := client.EthGetBlockByHash(header.Hash().Hex(), true)
			if err != nil {
				log.Errorf("%s - Block By Hash Error: %v\n", tokenName, err)
				log.Debugf("Block number: %v\nBlock Hash: %v", header.Number.String(), header.Hash().Hex())
				continue
			}
			for _, tx := range block.Transactions {
				if strings.EqualFold(tx.To, recipient.Hex()) {
					// log.Debugf("Native token transaction detected:\nTx hash: %s\nValue: %v", tx.Hash, tx.Value.String())
					if tx.Value.Cmp(big.NewInt(0)) > 0 {
						transfer := &contracts.ERC20Transfer{
							From:        common.Address(common.FromHex(tx.From)),
							To:          common.Address(common.FromHex(tx.To)),
							StringValue: tx.Value.String(),
							Raw:         types.Log{TxHash: common.HexToHash(tx.Hash)},
						}
						go onTransferReceived(transfer)
					}

				}
			}
		}
	}
}

func (b *Blockchain) WatchERC20Transactions(tokenName string, rpcAddress string, contractAddress common.Address, recipient common.Address, isReconnect bool, onTransferReceived func(transfer *contracts.ERC20Transfer)) {
	client, err := ethclient.Dial("wss://" + rpcAddress)
	if err != nil {
		if !isReconnect {
			log.Warning("first")
			log.Warning(err)
		}
		return
	}
	defer client.Close()

	recipientContract, err := contracts.NewERC20Filterer(contractAddress, client)
	if err != nil {
		if !isReconnect {
			log.Warning("second")
			log.Warning(err)
		}
		return
	}

	ch := make(chan *contracts.ERC20Transfer, 100)

	sub, err := recipientContract.WatchTransfer(nil, ch, nil, []common.Address{recipient})
	if err != nil {
		if !isReconnect {
			log.Warning("third")
			log.Warning(err)
		}
		return
	}
	defer sub.Unsubscribe()

	if isReconnect {
		log.Warning("The monitoring of payment transactions (" + tokenName + ") in the blockchain has been restored...")
	} else {
		log.Info("The monitoring of payment transactions (" + tokenName + ") in the blockchain has been started...")
	}
	for {
		select {
		case err := <-sub.Err():
			if err != nil {
				log.Error("Error while receiving transaction ("+tokenName+"):", err)
			}
			log.Warning("The monitoring of payment transactions (" + tokenName + ") in the blockchain has been interrupted. Retrying...")
			return
		case transfer := <-ch:
			go onTransferReceived(transfer)
		}
	}
}

func (b *Blockchain) GetFailingMessage(hash common.Hash) (string, error) {
	tx, _, err := b.PolygonClient.TransactionByHash(context.Background(), hash)
	if err != nil {
		return "", err
	}

	from, err := types.Sender(types.NewLondonSigner(tx.ChainId()), tx)
	if err != nil {
		return "", err
	}

	msg := ethereum.CallMsg{
		From:     from,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}

	res, err := b.PolygonClient.CallContract(context.Background(), msg, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
