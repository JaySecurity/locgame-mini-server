package router

import (
	"encoding/json"
	"io"
	"locgame-mini-server/internal/middleware"
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
	cookie := http.Cookie{Name: "SessionID", Value: session.SessionID, Path: "/", SameSite: http.SameSiteNoneMode, Secure: true, HttpOnly: true, MaxAge: int(3600)}
	http.SetCookie(w, &cookie)
	cookie = http.Cookie{Name: "IdToken", Value: authResponse.RefreshToken, Path: "/", SameSite: http.SameSiteNoneMode, Secure: true, HttpOnly: true, MaxAge: int(3600)}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (r *Router) SendLoginEmail(w http.ResponseWriter, req *http.Request) {

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

// func (r *Router) GetUserBalances(_ *network.Client, in *accounts.AccountBalanceRequest) (*accounts.AccountBalanceResponse, error) {
// 	return r.Accounts.GetUserBalances(in.Wallet)
// }

func (r *Router) HandleAccountRoutes() {
	// Get Store Data
	r.Mux.HandleFunc("/account", func(w http.ResponseWriter, req *http.Request) {
		_, _ = w.Write([]byte("Accounts"))
	})
	r.Mux.HandleFunc("/account/login/email", middleware.Log(r.SendLoginEmail))
	r.Mux.HandleFunc("/account/login/verifyemail", middleware.Log(r.VerifyLoginEmail))
	r.Mux.HandleFunc("/account/login/wallet", middleware.Log(r.Web3ChallengeRequest))
	r.Mux.HandleFunc("/account/login/fakewallet", middleware.Log(r.FakeWeb3Authorize))
	r.Mux.HandleFunc("/account/login/verifywallet", middleware.Log(r.Web3Authorize))
}
