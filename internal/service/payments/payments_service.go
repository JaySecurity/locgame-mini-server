package payments

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"locgame-mini-server/internal/blockchain"
	"locgame-mini-server/internal/blockchain/contracts"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/internal/sessions"
	"locgame-mini-server/internal/store"
	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/log"
	"math/big"
	"net/http"
	"runtime"
	"time"

	storeDto "locgame-mini-server/pkg/dto/store"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/shopspring/decimal"
)

type Service struct {
	store      *store.Store
	config     *config.Config
	sessions   *sessions.SessionStore
	blockchain *blockchain.Blockchain

	locgConvertRate *storeDto.LoCGConvertRate
	ethConvertRate  *storeDto.EthConvertRate

	OnPaymentSuccess func(ctx context.Context, order *storeDto.Order)
}

// New creates a new instance of the friends services.
func New(cfg *config.Config, sessions *sessions.SessionStore, store *store.Store, blockchain *blockchain.Blockchain) *Service {
	s := new(Service)
	s.config = cfg
	s.store = store
	s.sessions = sessions
	s.blockchain = blockchain

	s.locgConvertRate = new(storeDto.LoCGConvertRate)
	s.ethConvertRate = new(storeDto.EthConvertRate)

	if s.config.Environment != config.Development && s.config.Environment != config.Production && s.config.Environment != config.Staging {
		return s
	}

	// go s.watchNativeTransfer(
	// 	"ETH",
	// 	storeDto.PaymentMethod_ETH,
	// 	common.HexToAddress(s.config.Blockchain.PaymentRecipients.USDC),
	// 	s.config.Blockchain.RpcAddresses.Ethereum)

	// go s.watchNativeTransfer(
	// 	"ETH Base",
	// 	storeDto.PaymentMethod_ETHBase,
	// 	common.HexToAddress(s.config.Blockchain.PaymentRecipients.USDC),
	// 	s.config.Blockchain.RpcAddresses.Base)

	// go s.watchTransfer(
	// 	"USDC",
	// 	storeDto.PaymentMethod_USDC,
	// 	common.HexToAddress(s.config.Blockchain.Contracts.USDC),
	// 	common.HexToAddress(s.config.Blockchain.PaymentRecipients.USDC),
	// 	s.config.Blockchain.RpcAddresses.Ethereum,
	// )
	// go s.watchTransfer(
	// 	"USDT",
	// 	storeDto.PaymentMethod_USDT,
	// 	common.HexToAddress(s.config.Blockchain.Contracts.USDT),
	// 	common.HexToAddress(s.config.Blockchain.PaymentRecipients.USDT),
	// 	s.config.Blockchain.RpcAddresses.Ethereum,
	// )

	// go s.watchTransfer(
	// 	"LOCG",
	// 	storeDto.PaymentMethod_LOCG,
	// 	common.HexToAddress(s.config.Blockchain.Contracts.LOCG),
	// 	common.HexToAddress(s.config.Blockchain.PaymentRecipients.LOCG),
	// 	s.config.Blockchain.RpcAddresses.Ethereum,
	// )

	// go s.watchTransfer(
	// 	"BASE - LOCG",
	// 	storeDto.PaymentMethod_LOCGBase,
	// 	common.HexToAddress(s.config.Blockchain.Contracts.BaseLOCG),
	// 	common.HexToAddress(s.config.Blockchain.PaymentRecipients.LOCG),
	// 	s.config.Blockchain.RpcAddresses.Base,
	// )
	// go s.watchTransfer(
	// 	"BASE - USDC",
	// 	storeDto.PaymentMethod_USDCBase,
	// 	common.HexToAddress(s.config.Blockchain.Contracts.BaseUSDC),
	// 	common.HexToAddress(s.config.Blockchain.PaymentRecipients.USDC),
	// 	s.config.Blockchain.RpcAddresses.Base,
	// )

	return s
}

func (s *Service) GetLoCGRate() (*storeDto.LoCGConvertRate, error) {
	if s.locgConvertRate.NextUpdateTime == nil || s.locgConvertRate.NextUpdateTime.Seconds < time.Now().UTC().Unix() {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", fmt.Sprintf("https://pro-api.coinmarketcap.com/v1/tools/price-conversion?symbol=USDC&convert=LOCG&amount=1"), nil)
		req.Header.Set("X-CMC_PRO_API_KEY", s.config.Blockchain.CoinMarketCapApiKey)
		req.Header.Set("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var result convertData
		err = json.Unmarshal(data, &result)
		if err != nil {
			return nil, err
		}

		s.locgConvertRate.Price = result.Data.Quote.LOCG.Price
		s.locgConvertRate.NextUpdateTime = &base.Timestamp{Seconds: time.Now().UTC().Add(5 * time.Minute).Unix()}
	}

	return s.locgConvertRate, nil
}

func (s *Service) GetEthRate() (*storeDto.EthConvertRate, error) {
	if s.locgConvertRate.NextUpdateTime == nil || s.locgConvertRate.NextUpdateTime.Seconds < time.Now().UTC().Unix() {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", fmt.Sprintf("https://pro-api.coinmarketcap.com/v1/tools/price-conversion?symbol=USDC&convert=ETH&amount=1"), nil)
		req.Header.Set("X-CMC_PRO_API_KEY", s.config.Blockchain.CoinMarketCapApiKey)
		req.Header.Set("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var result convertData
		err = json.Unmarshal(data, &result)
		if err != nil {
			return nil, err
		}

		s.ethConvertRate.Price = result.Data.Quote.ETH.Price
		s.ethConvertRate.NextUpdateTime = &base.Timestamp{Seconds: time.Now().UTC().Add(5 * time.Minute).Unix()}
	}

	return s.ethConvertRate, nil
}

func (s *Service) watchTransfer(tokenName string, method storeDto.PaymentMethod, ERC20Contract common.Address, paymentRecipient common.Address, rpcAddress string) {
	isReconnect := false
	for {
		s.blockchain.WatchERC20Transactions(tokenName, rpcAddress, ERC20Contract, paymentRecipient, isReconnect, func(transfer *contracts.ERC20Transfer) {
			s.onTransferReceived(method, transfer)
		})
		time.Sleep(1 * time.Second)
		isReconnect = true
	}
}

func (s *Service) watchNativeTransfer(tokenName string, method storeDto.PaymentMethod, paymentRecipient common.Address, rpcAddress string) {
	isReconnect := false
	for {
		if method == storeDto.PaymentMethod_ETH {
			s.blockchain.WatchNativeEthereumTransactions(tokenName, rpcAddress, paymentRecipient, isReconnect, func(transfer *contracts.ERC20Transfer) {
				s.onTransferReceived(method, transfer)
			})
			time.Sleep(1 * time.Second)
			isReconnect = true
		} else if method == storeDto.PaymentMethod_ETHBase {
			s.blockchain.WatchNativeBaseTransactions(tokenName, rpcAddress, paymentRecipient, isReconnect, func(transfer *contracts.ERC20Transfer) {
				s.onTransferReceived(method, transfer)
			})
			time.Sleep(1 * time.Second)
			isReconnect = true
		}
	}
}

func (s *Service) onTransferReceived(method storeDto.PaymentMethod, transfer *contracts.ERC20Transfer) {
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Errorf("Panic at %v: %v\n%s", transfer.Raw.TxHash, err, buf)
		}
	}()
	ctx := context.Background()

	log.Debugf("\n\nPayment received. From: %v\nTxHash: %v\n\n", transfer.From.Hex(), transfer.Raw.TxHash.Hex())
	time.Sleep(30 * time.Second)

	// TODO Add temporary storage in Redis with TTL,
	//  if suddenly the client did not have time to send the transaction hash.

	mutex := s.store.DistributedLocks.NewLock("order:"+transfer.Raw.TxHash.Hex(), 1*time.Minute)
	if err := mutex.LockContext(ctx); err != nil {
		log.Debug("Duplicate")
		return
	}
	// Unlock will happen automatically after TTL expires

	order, err := s.store.Orders.GetByHash(context.Background(), transfer.Raw.TxHash.String())
	if err != nil {
		log.Error("Order not found")
		return
	}

	if order.Status != storeDto.OrderStatus_WaitingForPayment {
		log.Debug("Already completed")
		return
	}

	if order.PaymentMethod != method {
		log.Errorf("PaymentMethod != %v", method)
		return
	}

	order.Status = storeDto.OrderStatus_PaymentReceived

	err = s.store.Orders.Update(ctx, order)
	if err != nil {
		log.Error("Unable to update order in database:", err)
	}

	log.Debugf("Processing payment: %v", transfer.Raw.TxHash.Hex())

	if transfer.Value == nil {
		log.Debug("Nil Value using StringValue:", transfer.StringValue)
		value, ok := new(big.Int).SetString(transfer.StringValue, 10)
		if ok {
			transfer.Value = value
		} else {
			// handle the error case
			log.Error("Invalid value:", transfer.StringValue)
		}
	}

	price := new(big.Int)
	price.SetString(order.Price, 10)
	comparePrice := price.Cmp(transfer.Value)
	if comparePrice <= 0 {
		s.OnPaymentSuccess(ctx, order)
	} else {
		log.Error("Invalid price. Got:", transfer.Value, "Want:", price)
	}
}

// func (s *Service) Withdraw(ctx context.Context, in *resources.WithdrawRequest) (*resources.WithdrawResponse, error) {
// 	session := s.sessions.Get(id)
// 	adjustment := &resources.ResourceAdjustment{
// 		ResourceID: 1,
// 		Quantity:   in.LC,
// 	}
//
// 	withdrawalMinTime := time.Now().UTC().Add(-6 * 24 * time.Hour).Unix()
// 	if session.PlayerData.CreatedAt.Seconds > withdrawalMinTime || (session.PlayerData.PlayerStoreData.LastWithdrawalAt != nil && session.PlayerData.PlayerStoreData.LastWithdrawalAt.Seconds > withdrawalMinTime) {
// 		return nil, errors.ErrWithdrawalIsNotAvailableYet
// 	}
//
// 	}
//
// 	if in.LC < s.config.Withdraw.Min {
// 		return nil, errors.ErrWithdrawalAmountIsLowerThanRequired
// 	}
//
// 	if in.LC > s.config.Withdraw.Max {
// 		return nil, errors.ErrWithdrawalLimitExceeded
// 	}
//
// 	rate, err := s.GetLoCGRate()
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	withdrawAmount, burnAmount := s.calculateWithdraw(in.LC, rate.Price)
// 	tx, err := s.blockchain.WithdrawAndBurn(ctx, common.HexToAddress(session.PlayerData.ActiveWallet), withdrawAmount, burnAmount)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	adjustments, err := s.inventory.InverseAdjust(session, "withdraw", adjustment)
// 	if err != nil {
// 		log.Error("Failed to spent the LC from the player:", session.AccountID, "Error:", err)
// 	}
//
// 	session.PlayerData.PlayerStoreData.LastWithdrawalAt = &base.Timestamp{Seconds: time.Now().UTC().Unix()}
// 	err = s.store.InGameStore.SetLastWithdrawalAt(ctx, session.PlayerData.ID, session.PlayerData.PlayerStoreData.LastWithdrawalAt)
// 	if err != nil {
// 		log.Error("Failed to set withdrawal time for the player:", session.AccountID, "Error:", err)
// 	}
//
// 	amount, _ := WeiToEther(withdrawAmount).Float64()
// 	return &resources.WithdrawResponse{
// 		TransactionHash: tx,
// 		LOCG:            amount,
// 		Adjustments:     adjustments,
// 	}, nil
// }
//
// func (s *Service) calculateWithdraw(lcAmount int32, locgRate float64) (*big.Int, *big.Int) {
// 	locgAmountWithoutFee := (float64(lcAmount) * 0.001) * locgRate
//
// 	maxFeeWithdraw := s.config.Withdraw.Min * 10
//
// 	var fee float64
// 	if lcAmount <= maxFeeWithdraw {
// 		x := float64(lcAmount-s.config.Withdraw.Min) / float64(maxFeeWithdraw-s.config.Withdraw.Min)
// 		fee = (s.config.Withdraw.MaxFee-s.config.Withdraw.MinFee)*x + s.config.Withdraw.MinFee
// 	} else {
// 		fee = s.config.Withdraw.MaxFee
// 	}
//
// 	locgAmount := ToWei(locgAmountWithoutFee-(locgAmountWithoutFee*fee), 18)
// 	burnAmount := ToWei(locgAmountWithoutFee*0.05, 18)
//
// 	return locgAmount, burnAmount
// }

func WeiToEther(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), big.NewFloat(params.Ether))
}

func ToWei(iAmount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iAmount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}
