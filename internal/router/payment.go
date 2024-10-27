package router

import (
	"encoding/json"
	"fmt"
	"io"
	"locgame-mini-server/internal/middleware"
	"locgame-mini-server/pkg/log"
	"math/big"
	"net/http"
	"os"
	"strings"
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

func (r *Router) HandlePaymentRoutes() {
	m := middleware.NewMiddleWare(r.config)
	r.Mux.HandleFunc("/payment", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Content-Type", "application/text")
		_, _ = w.Write([]byte("Payments"))
	})

	r.Mux.HandleFunc("GET /payment/gas", m.Logger(r.GetGas))
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

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

// func GweiToWei(gwei string) (*big.Float, error) {
// 	gweiFloat, ok := new(big.Float).SetString(gwei)
// 	if !ok {
// 		return nil, fmt.Errorf("Unable to convert to Wei")
// 	}
// 	return new(big.Float).Mul(gweiFloat, big.NewFloat(params.GWei)), nil
// }

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
