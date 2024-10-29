package router

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"locgame-mini-server/internal/blockchain/contracts"
	"locgame-mini-server/pkg/dto/store"
	"locgame-mini-server/pkg/log"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type Result struct {
	LastBlock       string
	SafeGasPrice    string
	ProposeGasPrice string
	FastGasPrice    string
	SuggestBaseFee  string
	GasUsedRatio    string
}
type Response struct {
	Status  string
	Message string
	Result  *Result
}

type GasResponse struct {
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
	MaxFeePerGas         string `json:"maxFeePerGas"`
}

type SubmitHashRequest struct {
	TransactionHash string `json:"txHash"`
	BuyerID         string `json:"buyer_id"`
	OrderID         string `json:"order_id"`
	Wallet          string `json:"wallet"`
}

func (r *Router) HandlePaymentRoutes() {
	m := r.middleware
	r.Mux.HandleFunc("/payment", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Content-Type", "application/text")
		_, _ = w.Write([]byte("Payments"))
	})

	r.Mux.HandleFunc("GET /payment/gas", m.Logger(r.GetGas))
	r.Mux.HandleFunc("/payment/submit", m.Logger(r.SubmitTx))
}

func (r *Router) GetGas(w http.ResponseWriter, req *http.Request) {
	log.Info("get gas")
	ETHERSCAN_API_KEY := os.Getenv("ETHERSCAN_API_KEY")
	url := fmt.Sprintf("https://api.etherscan.io/api?module=gastracker&action=gasoracle&apikey=%s", ETHERSCAN_API_KEY)
	res, err := http.Get(url)
	if err != nil {
		log.Errorf("Error fetching gas: %v", err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error("Error reading response body", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	response := &Response{}
	err = json.Unmarshal(body, response)
	if err != nil {
		log.Error("Error parsing response body", err)

		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	result := response.Result

	maxFeePerGas, err := GweiToWei(result.SafeGasPrice)
	if err != nil {
		log.Error("Error parsing response body", err)

		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	maxPriorityFeePerGas, err := GweiToWei("1.5")
	if err != nil {
		log.Error("Error parsing response body", err)

		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	gas := &GasResponse{
		MaxPriorityFeePerGas: BigIntToHex(maxPriorityFeePerGas),
		MaxFeePerGas:         BigIntToHex(maxFeePerGas),
	}

	jsonData, err := json.Marshal(gas)
	if err != nil {
		log.Error("Error reading gas response body", err)
		errMsg := &ErrorMsg{
			Message: "Error reading gas response body",
			Code:    "",
		}
		jsondata, _ := json.Marshal(errMsg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsondata)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func (r *Router) SubmitTx(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	addressMap := map[store.PaymentMethod]string{
		store.PaymentMethod_ETH:      r.config.Blockchain.PaymentRecipients.LOCG,
		store.PaymentMethod_ETHBase:  r.config.Blockchain.PaymentRecipients.LOCG,
		store.PaymentMethod_LOCGBase: r.config.Blockchain.Contracts.BaseLOCG,
		store.PaymentMethod_USDCBase: r.config.Blockchain.Contracts.BaseUSDC,
		store.PaymentMethod_LOCG:     r.config.Blockchain.Contracts.LOCG,
		store.PaymentMethod_USDT:     r.config.Blockchain.Contracts.USDT,
	}
	in := &SubmitHashRequest{}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Error("Error reading request body", err)
		errMsg := &ErrorMsg{
			Message: "Error reading request body",
			Code:    "",
		}
		jsondata, _ := json.Marshal(errMsg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsondata)
		return
	}
	err = json.Unmarshal(body, in)
	if err != nil {
		log.Error("Error parsing request body", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	order, err := r.store.Orders.Get(ctx, in.OrderID)
	if err != nil {
		log.Error("Error Fetching Order", err)
		errMsg := &ErrorMsg{
			Message: "Error fetching order",
			Code:    mongo.ErrNoDocuments.Error(),
		}
		jsondata, _ := json.Marshal(errMsg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsondata)
		return
	}
	if order.Status != store.OrderStatus_WaitingForPayment {
		log.Error("Invalid Order Status")
		errMsg := &ErrorMsg{
			Message: "Invalid order status",
			Code:    order.Status.String(),
		}
		jsondata, _ := json.Marshal(errMsg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsondata)
		return
	}
	if order.BuyerID.Value != in.BuyerID {
		log.Error("Buyer Id Mismatch")
		errMsg := &ErrorMsg{
			Message: "Buyer Id Mismatch",
			Code:    "",
		}
		jsondata, _ := json.Marshal(errMsg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsondata)
		return
	}
	if order.PaymentHash != "" {
		log.Error("Transaction Hash Already Submitted")
		errMsg := &ErrorMsg{
			Message: "Transaction Hash Already Submitted",
			Code:    order.PaymentHash,
		}
		jsondata, _ := json.Marshal(errMsg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsondata)
		return
	}
	order.PaymentHash = in.TransactionHash
	err = r.store.Orders.Update(ctx, order)
	if err != nil {
		log.Error("Error Updating Order")
		errMsg := &ErrorMsg{
			Message: "Error Updating Order",
			Code:    err.Error(),
		}
		jsondata, _ := json.Marshal(errMsg)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsondata)
		return
	}

	value, ok := new(big.Int).SetString(order.Price, 10)
	if !ok {
		log.Error("Error setting value")
		errMsg := &ErrorMsg{
			Message: "Error setting value",
			Code:    "",
		}
		jsondata, _ := json.Marshal(errMsg)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsondata)
		return
	}
	transfer := &contracts.ERC20Transfer{
		From:  common.HexToAddress(in.Wallet),
		To:    common.HexToAddress(addressMap[order.PaymentMethod]),
		Value: value,
		Raw:   types.Log{TxHash: common.HexToHash(order.PaymentHash)},
	}
	go r.Payments.OnTransferReceived(order.PaymentMethod, transfer)
	w.WriteHeader(http.StatusNoContent)
}

func GweiToWei(gweiStr string) (*big.Int, error) {
	gweiStr = strings.TrimSpace(gweiStr)
	gweiFloat, ok := new(big.Float).SetString(gweiStr)
	if !ok {
		return nil, fmt.Errorf("invalid Gwei value")
	}
	gweiToWeiFactor := big.NewFloat(1e9)
	weiFloat := new(big.Float).Mul(gweiFloat, gweiToWeiFactor)
	weiInt := new(big.Int)
	weiFloat.Int(weiInt)
	return weiInt, nil
}

func BigIntToHex(b *big.Int) string {
	return "0x" + b.Text(16)
}

// func CheckTransaction(tokenName string, rpcAddress string, recipient common.Address, isReconnect bool, onTransferReceived func(tx *contracts.ERC20Transfer)) {
// 	client, err := ethclient.Dial(`wss://` + rpcAddress)
// 	if err != nil {
// 		if !isReconnect {
// 			log.Warning("first")
// 			log.Warning(err)
// 		}
// 		return
// 	}
// 	defer client.Close()

// 	headers := make(chan *types.Header)
// 	sub, err := client.SubscribeNewHead(context.Background(), headers)
// 	if err != nil {
// 		if !isReconnect {
// 			log.Warning("second")
// 			log.Warning(err)
// 		}
// 		return
// 	}
// 	defer sub.Unsubscribe()

// 	if isReconnect {
// 		log.Warning("The monitoring of native token transactions (" + tokenName + ") in the blockchain has been restored...")
// 	} else {
// 		log.Info("The monitoring of native token transactions (" + tokenName + ") in the blockchain has been started...")
// 	}

// 	for {
// 		select {

// 		case err := <-sub.Err():
// 			if err != nil {
// 				if !isReconnect {
// 					log.Warning("third")
// 					log.Warning(err)
// 				}
// 				return
// 			}
// 		case header := <-headers:
// 			block, err := client.BlockByHash(context.Background(), header.Hash())
// 			if err != nil {
// 				log.Errorf("%s - Block By Hash Error: %v\n", tokenName, err)
// 				log.Debugf("Block number: %v\nBlock Hash: %v", header.Number.String(), header.Hash().Hex())
// 				continue
// 			}
// 			for _, tx := range block.Transactions() {
// 				// Get the signer for the transaction
// 				signer := types.LatestSignerForChainID(tx.ChainId())
// 				// Derive the sender's address
// 				sender, err := types.Sender(signer, tx)
// 				if err != nil {
// 					log.Fatalf("Failed to get sender from transaction: %v", err)
// 				}
// 				if tx.To() != nil && *tx.To() == recipient {
// 					log.Debugf("Native token transaction detected:\nTx hash: %s\n", tx.Hash().Hex())

// 					if tx.Value().Cmp(big.NewInt(0)) > 0 {
// 						transfer := &contracts.ERC20Transfer{
// 							From:  sender,
// 							To:    *tx.To(),
// 							Value: tx.Value(),
// 							Raw:   types.Log{TxHash: tx.Hash()},
// 						}
// 						go onTransferReceived(transfer)
// 					}

// 				}
// 			}
// 		}
// 	}
// }
