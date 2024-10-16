package router

import (
	"encoding/json"
	"io"
	"locgame-mini-server/pkg/dto/accounts"
	"locgame-mini-server/pkg/log"
	"net/http"
)

func (r *Router) Web3ChallengeRequest(w http.ResponseWriter, req *http.Request) {
	in := &accounts.Web3AuthRequest{}
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
	challengeResponse, err := r.Accounts.RequestChallenge(in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jsonData, err := json.Marshal(challengeResponse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (r *Router) FakeWeb3Authorize(w http.ResponseWriter, req *http.Request) {
	in := &accounts.Web3AuthRequest{}
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
	authResponse, session, err := r.Accounts.FakeWeb3Authorize(in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jsonData, err := json.Marshal(authResponse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cookie := http.Cookie{Name: "SessionID", Value: session.SessionID, Path: "/", HttpOnly: true, MaxAge: int(3600)}
	http.SetCookie(w, &cookie)
	cookie = http.Cookie{Name: "IdToken", Value: authResponse.RefreshToken, Path: "/", HttpOnly: true, MaxAge: int(3600)}
	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (r *Router) Web3Authorize(w http.ResponseWriter, req *http.Request) {
	in := &accounts.Web3Signature{}
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
	authResponse, session, err := r.Accounts.Web3Authorize(in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jsonData, err := json.Marshal(authResponse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cookie := http.Cookie{Name: "SessionID", Value: session.SessionID, Path: "/", HttpOnly: true, MaxAge: int(3600)}
	http.SetCookie(w, &cookie)
	cookie = http.Cookie{Name: "IdToken", Value: authResponse.RefreshToken, Path: "/", HttpOnly: true, MaxAge: int(3600)}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// func (r *Router) Logout( _ *base.Empty) (*base.Empty, error) {
// 	return nil, r.Accounts.Logout(client.Context())
// }

// func (r *Router) AuthToken( in *accounts.RefreshTokenRequest) (*accounts.LoginResponse, error) {
// 	return r.Accounts.AuthToken(client.Context(), in)
// }

// func (r *Router) SetAccountInfo( in *accounts.SetInfoRequest) (*base.Empty, error) {
// 	err := r.Accounts.SetInfo(client.Context(), in)
// 	return &base.Empty{}, err
// }

func (r *Router) SendLoginEmail(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/text")
	var loginRequest accounts.LoginEmailRequest
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Error("Error reading request body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &loginRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	loginResponse, err := r.Accounts.SendLoginEmail(&loginRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jsonData, err := json.Marshal(loginResponse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (r *Router) VerifyLoginEmail(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/text")
	verifyRequest := &accounts.VerifyLoginEmailRequest{}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Error("Error reading request body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, verifyRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	loginResponse, session, err := r.Accounts.VerifyLoginEmail(verifyRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jsonData, err := json.Marshal(loginResponse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cookie := http.Cookie{Name: "SessionID", Value: session.SessionID, Path: "/", HttpOnly: true, MaxAge: int(3600)}
	http.SetCookie(w, &cookie)
	cookie = http.Cookie{Name: "IdToken", Value: loginResponse.RefreshToken, Path: "/", HttpOnly: true, MaxAge: int(3600)}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// func (r *Router) SetActiveWallet( in *accounts.SetActiveWalletRequest) (*base.Empty, error) {
// 	err := r.Accounts.SetActiveWallet(client.Context(), in)
// 	return &base.Empty{}, err
// }

// func (r *Router) GetUserBalances(_ *network.Client, in *accounts.AccountBalanceRequest) (*accounts.AccountBalanceResponse, error) {
// 	return r.Accounts.GetUserBalances(in.Wallet)
// }

func (r *Router) HandleAccountRoutes() {
	// Get Store Data
	r.Mux.HandleFunc("/account", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		w.Header().Set("Content-Type", "application/text")
		_, _ = w.Write([]byte("Accounts"))
	})
	r.Mux.HandleFunc("/account/login/email", r.SendLoginEmail)
	r.Mux.HandleFunc("/account/login/verifyemail", r.VerifyLoginEmail)
	r.Mux.HandleFunc("/account/login/wallet", r.Web3ChallengeRequest)
	r.Mux.HandleFunc("/account/login/fakewallet", r.FakeWeb3Authorize)
	r.Mux.HandleFunc("/account/login/verifywallet", r.Web3Authorize)
}
