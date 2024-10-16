package blockchain

import (
	"context"
	"fmt"
	"locgame-mini-server/internal/blockchain/contracts"
	"locgame-mini-server/pkg/log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractInstance interface {
	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)
}
type BalanceChecker struct {
	client   *ethclient.Client
	contract ContractInstance
}

func NewBalanceChecker(network string, tokenContractAAddress string) (*BalanceChecker, error) {
	network = fmt.Sprint("https://", network)
	client, err := ethclient.Dial(network)
	if err != nil {
		return nil, err
	}
	var contract ContractInstance = nil
	tokenAddress := common.HexToAddress(tokenContractAAddress)
	if tokenContractAAddress != "" {
		contract, err = contracts.NewERC20(tokenAddress, client)
		if err != nil {
			return nil, err
		}
	}
	return &BalanceChecker{
		client:   client,
		contract: contract,
	}, nil
}

func (b *BalanceChecker) IsTransferable(amount *big.Int, from common.Address, toAddress common.Address) error {
	eth, err := b.GetEthBalance(from)
	if err != nil {
		return err
	}
	// log.Debug("getEthBalance", eth)
	gasFee, err := b.estimateGasFee(amount, from, toAddress)
	if err != nil {
		return err
	}
	// log.Debug("estimateGasFee", gasFee)
	if b.contract != nil {
		tokenBalance, err := b.GetTokenBalance(from)
		if err != nil {
			return err
		}

		if amount.Cmp(tokenBalance) >= 0 {
			return fmt.Errorf("not enough token")
		}

		// log.Debug("getTokenBalance", tokenBalance)

		if gasFee.Cmp(eth) >= 0 {
			return fmt.Errorf("not enough gas fee")
		}

		if big.NewInt(0).Cmp(tokenBalance) >= 0 {
			return fmt.Errorf("no tokens in wallet")
		}
		return nil
	} else {
		total := new(big.Int).Add(amount, gasFee)
		if total.Cmp(eth) >= 0 {
			return fmt.Errorf("not enough gas fee")
		}
		return nil
	}
}

func (b *BalanceChecker) estimateGasFee(value *big.Int, from common.Address, toAddress common.Address) (*big.Int, error) {
	gasPrice, err := b.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Debug("Failed to get gas price: ", err)
		return nil, err
	}
	var data []byte
	var ethValue *big.Int
	if b.contract != nil {
		transferFnSignature := []byte("transfer(address,uint256)")
		hash := sha3.NewLegacyKeccak256()
		hash.Write(transferFnSignature)
		methodID := hash.Sum(nil)[:4]

		paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

		paddedAmount := common.LeftPadBytes(value.Bytes(), 32)

		data = append(data, methodID...)
		data = append(data, paddedAddress...)
		data = append(data, paddedAmount...)
		//---------------------
	} else {
		data = nil
		ethValue = value
	}
	// Estimate gas for the transaction
	estimatedGas, err := b.client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  from,
		To:    &toAddress,
		Data:  data,
		Value: ethValue,
	})

	gasLimit := int(float64(estimatedGas) * 1.0)

	if err != nil {
		return nil, err
	}

	// Convert gasLimit to a *big\.Int
	gasLimitBigInt := big.NewInt(int64(gasLimit))

	// Calculate the total gas cost
	totalGasCost := new(big.Int).Mul(gasPrice, gasLimitBigInt)

	// log.Debug("gas price is ", gasPrice)
	//log.Debug(hexutil.Encode(methodID))
	//log.Debug(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d
	//log.Debug(hexutil.Encode(paddedAmount))  // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000
	// log.Debug("Estimated gas limit:", gasLimit)
	//log.Debug("gas limit big is ", gasLimitBigInt)
	// log.Debug("Estimated gas fee:", totalGasCost)
	return totalGasCost, nil
}

func (b *BalanceChecker) GetEthBalance(walletAddress common.Address) (*big.Int, error) {
	balance, err := b.client.BalanceAt(context.Background(), walletAddress, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (b *BalanceChecker) GetTokenBalance(address common.Address) (*big.Int, error) {

	tokenBalance, err := b.contract.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		return nil, err
	}

	return tokenBalance, nil
}
