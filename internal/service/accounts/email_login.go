package accounts

import (
	"context"
	"fmt"
	common "locgame-mini-server/internal/service/accounts/lib"
	"locgame-mini-server/internal/sessions"
	"locgame-mini-server/pkg/dto/accounts"
	"locgame-mini-server/pkg/dto/errors"
	"locgame-mini-server/pkg/dto/player"
	"locgame-mini-server/pkg/log"
	"math/rand"
	"strings"
	"time"
)

// SendLoginEmail sends a login code to the user.

func (s *Service) SendLoginEmail(request *accounts.LoginEmailRequest) (*accounts.LoginEmailResponse, error) {
	ctx := context.Background()
	valid := common.ValidateEmail(request.Email)
	if !valid {
		return nil, errors.ErrUserNotFound
	}
	email := strings.ToLower(request.Email)
	wallet := ""
	if request.Wallet != "" {
		wallet = strings.ToLower(request.Wallet)
	}
	// Generate OTP Code, Create Redis Entry send Email via SES

	challengeCode := generateOTP()

	html, text := generateEmail(challengeCode)

	err := s.sesInstance.SendEmail(email, "Legends of Crypto Login Code", html, text)
	if err != nil {
		log.Error(fmt.Sprintf("failed to send email: %v", err))
		return nil, err
	}
	err = s.store.Auth.AddEmailChallenge(ctx, email, challengeCode, wallet)
	if err != nil {
		log.Error(fmt.Sprintf("failed to add email challenge: %v", err))
		return nil, err
	}
	if wallet != "" {
		return &accounts.LoginEmailResponse{
			ChallengeName: "OTP_LOGIN",
			Session:       fmt.Sprintf("%s|%t|%s", email, request.IsMetaMask, wallet),
		}, nil
	}
	return &accounts.LoginEmailResponse{
		ChallengeName: "OTP_LOGIN",
		Session:       fmt.Sprintf("%s|%t", email, request.IsMetaMask),
	}, nil
}

// VerifyLoginEmail verifies the login code sent by the user matches the code sent by the server.

func (s *Service) VerifyLoginEmail(request *accounts.VerifyLoginEmailRequest) (*accounts.LoginResponse, *sessions.Session, error) {
	ctx := context.Background()
	isMetaMask := false
	email := strings.ToLower(request.Email)
	data, err := s.store.Auth.GetEmailChallenge(ctx, email)

	if err != nil {
		log.Error(fmt.Sprintf("failed to get challenge record: %v", err))
		return nil, nil, err
	}

	if data.Code != request.Code {
		return nil, nil, errors.ErrUserNotConfirmed
	}

	sessionInfo := strings.Split(request.Session, "|")
	if sessionInfo[1] == "true" {
		isMetaMask = true
	}
	if len(sessionInfo) > 2 {
		ctx := context.Background()
		wallet := sessionInfo[2]
		accountId, err := s.store.Players.GetAccountIDByWallet(ctx, wallet)
		if err != nil {
			return nil, nil, errors.ErrUserNotFound
		}
		err = s.store.Players.SetData(ctx, accountId, &player.PlayerData{Email: email})
		if err != nil {
			return nil, nil, err
		}
	}

	session, err := s.NewSessionBuilder(ctx, s.config).SetEmailAddress(email).SetIsMetaMaskUser(isMetaMask).CreateSession().Build()
	if err != nil {
		log.Error(fmt.Sprintf("failed to create session: %v", err))
		return nil, nil, err
	}

	return s.createLoginResponse(ctx, session)
}

func generateOTP() string {
	rand.Seed(time.Now().UnixNano()) // Add this line to seed the random number generator.
	min := 100000
	max := 999999
	otp := rand.Intn(max-min+1) + min
	return fmt.Sprintf("%06d", otp)
}

func generateEmail(code string) (string, string) {
	html := "<h1>Welcome!  " +
		"<h3>Your Secret Login Code Is: </h3>" +
		"<h2>" + code + "</h2>" +
		"<h3>Please enter this code to login.</h3>"

	text := "Welcome! " +
		"Your Secret Login Code Is: " +
		code +
		"Please enter this code to login."

	return html, text

}
