package router

import (
	"encoding/json"
	"io"
	storeDto "locgame-mini-server/pkg/dto/store"
	"locgame-mini-server/pkg/log"
	"net/http"
)

type ErrorMsg struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (r *Router) CreateOrder(w http.ResponseWriter, req *http.Request) {
	sessionIdCookie, err := req.Cookie("SessionID")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionId := sessionIdCookie.Value
	session := r.Sessions.Get(sessionId)
	if session == nil {
		log.Debug("Session Not Found")
		errMsg := &ErrorMsg{
			Message: "Session Not Found",
			Code:    "ErrInvalidSession",
		}
		jsondata, _ := json.Marshal(errMsg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsondata)
		return
	}
	in := &storeDto.OrderRequest{}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Error("Error reading request body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	orderResponse, err := r.InGameStore.CreateOrder(session.Context, sessionId, in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jsonData, err := json.Marshal(orderResponse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// func (r *Router) CreateUpgradeOrder(client *network.Client, in *storeDto.UpgradeRequest) (*storeDto.OrderResponse, error) {
// 	return r.InGameStore.CreateUpgradeOrder(client.Context(), in)
// }

func (r *Router) SendPaymentReceipt(w http.ResponseWriter, req *http.Request) {
	sessionIdCookie, err := req.Cookie("SessionID")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionId := sessionIdCookie.Value
	session := r.Sessions.Get(sessionId)
	if session == nil {
		log.Debug("Session Not Found")
		errMsg := &ErrorMsg{
			Message: "Session Not Found",
			Code:    "ErrInvalidSession",
		}
		jsondata, _ := json.Marshal(errMsg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsondata)
		return
	}
	in := &storeDto.Receipt{}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Error("Error reading request body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = r.InGameStore.SetPaymentReceipt(session.Context, in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// func (r *Router) OpenPack(client *network.Client, in *base.ObjectID) (*storeDto.OpenPackResponse, error) {
// 	return r.InGameStore.OpenPack(client.Context(), in)
// }

// func (r *Router) GetUnopenedPacks(client *network.Client, _ *base.Empty) (*storeDto.Orders, error) {
// 	orders, err := r.InGameStore.GetUnopenedPacks(client.Context())
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &storeDto.Orders{Data: orders}, nil
// }

// func (r *Router) GetLoCGRate(_ *network.Client, _ *base.Empty) (*storeDto.LoCGConvertRate, error) {
// 	return r.Payments.GetLoCGRate()
// }

// func (r *Router) GetInGameStoreData(client *network.Client, _ *base.Empty) (*storeDto.StoreData, error) {
// 	return r.InGameStore.GetData(client.Context())
// }

// func (r *Router) SubmitPromoCode(client *network.Client, in *storeDto.PromoCodeSubmitRequest) (*storeDto.PromoCodeSubmitResponse, error) {
// 	return r.InGameStore.SubmitPromoCode(client.Context(), in)
// }

// Store data route
func (r *Router) HandleStoreRoutes() {
	// Get Store Data
	r.Mux.HandleFunc("/store", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		data, err := r.InGameStore.GetData(req.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Debug("Store data:", data.Tokens[0].Available)
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(jsonData)
	})

	r.Mux.HandleFunc("POST /order", r.CreateOrder)
	r.Mux.HandleFunc("PATCH /order", r.SendPaymentReceipt)
}
