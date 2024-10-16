package accounts

import (
	"context"
	"fmt"
	"locgame-mini-server/internal/blockchain"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/internal/service/accounts/cognito"
	awsses "locgame-mini-server/internal/service/accounts/ses"
	"locgame-mini-server/internal/service/accounts/siwe"
	"locgame-mini-server/internal/service/shared"
	"locgame-mini-server/internal/sessions"
	"locgame-mini-server/internal/store"
	"locgame-mini-server/pkg/dto/accounts"
	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/dto/errors"
	"locgame-mini-server/pkg/dto/player"
	"locgame-mini-server/pkg/log"
	"locgame-mini-server/pkg/pubsub"
	"locgame-mini-server/pkg/stime"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang-jwt/jwt/v4"
)

type Service struct {
	store           *store.Store
	config          *config.Config
	sessions        *sessions.SessionStore
	cognitoInstance *cognito.Cognito
	sesInstance     *awsses.SES

	jwt JWT

	accountInfoChanged []shared.AccountInfoUpdateListener

	OnLogout          func(ctx context.Context)
	UserLoginListener []LoginListener
}

// New creates a new instance of the account services.
func New(
	config *config.Config,
	sessions *sessions.SessionStore,
	store *store.Store,
) *Service {
	s := new(Service)
	s.config = config
	s.store = store
	s.sessions = sessions
	s.cognitoInstance = cognito.NewCognito(config.CognitoEnv)
	s.sesInstance = awsses.NewSES(config.Ses)

	prvKey, err := os.ReadFile("configs/cert/cert.key")
	if err != nil {
		log.Fatal(err)
	}
	pubKey, err := os.ReadFile("configs/cert/cert.pub")
	if err != nil {
		log.Fatal(err)
	}

	s.jwt = NewJWT(prvKey, pubKey)

	// pubsub.RegisterPlayerHandler(&DuplicateLoginHandler{sessions: sessions})
	// pubsub.RegisterPlayerHandler(
	// 	&PlayerDataChangedHandler{sessions: sessions, cards: cards, store: store},
	// )

	s.RegisterListener(s)

	return s
}

func (s *Service) SubscribeAccountInfoChange(listener shared.AccountInfoUpdateListener) {
	s.accountInfoChanged = append(s.accountInfoChanged, listener)
}

func (s *Service) loginWithAuthToken(
	ctx context.Context,
	refreshToken string,
) (*accounts.LoginResponse, *sessions.Session, error) {
	userData, err := s.GetUserInfo(refreshToken)
	if err != nil {
		return nil, nil, errors.ErrInvalidToken
	}

	session, err := s.NewSessionBuilder(ctx, s.config).
		SetWallet(userData.ActiveWallet).
		CreateSession().Build()
	if err != nil {
		return nil, nil, err
	}

	return s.createLoginResponse(ctx, session)
}

func (s *Service) createLoginResponse(
	ctx context.Context,
	session *sessions.Session,
) (*accounts.LoginResponse, *sessions.Session, error) {
	userData := &accounts.UserData{
		ID:           session.PlayerData.ID,
		Name:         session.PlayerData.Name,
		ActiveWallet: session.PlayerData.ActiveWallet,
	}

	token, err := s.jwt.Create(48*time.Hour, userData)
	if err != nil {
		return nil, nil, err
	}

	return &accounts.LoginResponse{
		RefreshToken:   token,
		UserData:       userData,
		UnixServerTime: &base.Timestamp{Seconds: stime.Now(ctx).Unix()},
	}, session, nil
}

func (s *Service) GetUserInfo(idToken string) (*accounts.UserData, error) {
	mapClaims, err := s.getClaims(idToken)
	if err != nil {
		return nil, err
	}

	data := &accounts.UserData{
		ID:           &base.ObjectID{Value: s.tryGetStringValueFromClaims("sub", mapClaims)},
		Name:         s.tryGetStringValueFromClaims("given_name", mapClaims),
		ActiveWallet: s.tryGetStringValueFromClaims("wallet", mapClaims),
	}

	return data, nil
}

func (s *Service) getClaims(idToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(idToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", token.Header["alg"])
		}

		key, err := jwt.ParseRSAPublicKeyFromPEM(s.jwt.publicKey)
		if err != nil {
			return "", fmt.Errorf("validate: parse key: %w", err)
		}
		return key, nil
	}, jwt.WithoutClaimsValidation())
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.ErrInvalidToken
	}

	vErr := new(jwt.ValidationError)
	now := time.Now().UTC().Unix()

	if !mapClaims.VerifyExpiresAt(now, false) {
		vErr.Inner = jwt.ErrTokenExpired
		vErr.Errors |= jwt.ValidationErrorExpired
	}

	if !mapClaims.VerifyNotBefore(now, false) {
		vErr.Inner = jwt.ErrTokenNotValidYet
		vErr.Errors |= jwt.ValidationErrorNotValidYet
	}

	if vErr.Errors != 0 {
		return nil, vErr
	}

	return mapClaims, nil
}

func (s *Service) tryGetStringValueFromClaims(key string, mapClaims jwt.MapClaims) string {
	if value, ok := mapClaims[key]; ok {
		return value.(string)
	}
	return ""
}

func (s *Service) Logout(ctx context.Context) error {
	if s.OnLogout != nil {
		s.OnLogout(ctx)
	}
	return nil
}

// func (s *Service) AuthToken(
// 	ctx context.Context,
// 	in *accounts.RefreshTokenRequest,
// ) (*accounts.LoginResponse, error) {
// 	return s.loginWithAuthToken(ctx, in.RefreshToken)
// }

func (s *Service) SetInfo(id string, in *accounts.SetInfoRequest) error {
	session := s.sessions.Get(id)
	session.PlayerData.Name = in.Name
	session.PlayerData.AvatarID = in.AvatarID
	ctx := context.Background()
	// for _, listener := range s.accountInfoChanged {
	// 	listener.OnAccountInfoChanged(ctx, &accounts.AccountInfo{
	// 		ID:       session.PlayerData.ID,
	// 		Name:     in.Name,
	// 		AvatarID: in.AvatarID,
	// 	})
	// }
	return s.store.Players.SetData(
		ctx,
		session.AccountID,
		&player.PlayerData{Name: in.Name, AvatarID: in.AvatarID},
	)
}

func (s *Service) SetOnlineState(id string, isOnline bool) {
	ctx := context.Background()
	if session := s.sessions.Get(id); session != nil {
		_ = s.store.Players.SetOnlineState(ctx, session.AccountID, isOnline)
	}
}

func (s *Service) RequestChallenge(
	in *accounts.Web3AuthRequest,
) (*accounts.Web3ChallengeResponse, error) {
	address := strings.ToLower(in.Address)
	if in.Address == common.BigToAddress(common.Big0).Hex() {
		return nil, errors.ErrZeroWalletAddress
	}
	client := ""
	switch in.Client {
	case "sale":
		client = "sale"
	default:
		client = "game"

	}

	// Lookup Account by Wallet Address
	ctx := context.Background()
	_, err := s.store.Players.GetAccountIDByWallet(ctx, address)
	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	nonce := siwe.GenerateNonce()
	message, err := siwe.InitMessage(
		address,
		nonce,
		time.Now().UTC(),
		time.Now().UTC().Add(48*time.Hour),
		client,
	)
	if err != nil {
		return nil, err
	}

	err = s.store.Auth.AddChallenge(ctx, address, message)
	if err != nil {
		return nil, err
	}

	return &accounts.Web3ChallengeResponse{Challenge: message.String()}, err
}

func (s *Service) Web3Authorize(
	in *accounts.Web3Signature,
) (*accounts.LoginResponse, *sessions.Session, error) {
	ctx := context.Background()
	address := strings.ToLower(in.Address)
	authMessage, err := s.store.Auth.GetChallenge(ctx, address)
	// log.Debug("Auth Message: " + authMessage.String())
	if err != nil {
		return nil, nil, err
	}

	publicKey, err := authMessage.Verify(in.Signature, authMessage.Nonce)
	if err != nil {
		log.Error(err)
		return nil, nil, errors.ErrInvalidSignature
	}

	if strings.ToLower(crypto.PubkeyToAddress(*publicKey).Hex()) != address {
		return nil, nil, errors.ErrInvalidSignature
	}

	s.store.Auth.DeleteChallenge(ctx, address)

	session, err := s.NewSessionBuilder(ctx, s.config).
		SetWallet(address).
		CreateSession().Build()
	if err != nil {
		return nil, nil, err
	}

	return s.createLoginResponse(ctx, session)
}

func (s *Service) FakeWeb3Authorize(
	in *accounts.Web3AuthRequest,
) (*accounts.LoginResponse, *sessions.Session, error) {
	if s.config.Environment == config.Production {
		return nil, nil, errors.ErrNotAuthorized
	}

	ctx := context.Background()

	address := strings.ToLower(in.Address)

	// Lookup Account by Wallet Address
	_, err := s.store.Players.GetAccountIDByWallet(ctx, address)
	if err != nil {
		return nil, nil, errors.ErrUserNotFound
	}

	session, err := s.NewSessionBuilder(ctx, s.config).
		SetWallet(address).
		CreateSession().Build()
	if err != nil {
		return nil, nil, err
	}

	return s.createLoginResponse(ctx, session)
}

func (s *Service) SetActiveWallet(id string, in *accounts.SetActiveWalletRequest) error {
	walletAddress := strings.ToLower(in.Wallet)
	if session := s.sessions.Get(id); session != nil {
		ctx := context.Background()

		if walletAddress == session.PlayerData.ActiveWallet {
			return nil
		} else if session.PlayerData.ActiveWallet == "" {
			err := s.store.Players.SetActiveWallet(ctx, session.AccountID, in)
			if err == nil {
				session.PlayerData.ActiveWallet = strings.ToLower(in.Wallet)

				data := &player.PlayerData{
					ID:                  session.PlayerData.ID,
					ActiveWallet:        session.PlayerData.ActiveWallet,
					Name:                session.PlayerData.Name,
					AvatarID:            session.PlayerData.AvatarID,
					LastActivity:        session.PlayerData.LastActivity,
					Online:              session.PlayerData.Online,
					Decks:               session.PlayerData.Decks,
					Resources:           session.PlayerData.Resources,
					ResettableResources: session.PlayerData.ResettableResources,
					StoryMode:           session.PlayerData.StoryMode,
					ArenaData:           session.PlayerData.ArenaData,
					PlayerStoreData:     session.PlayerData.PlayerStoreData,
					TutorialData:        session.PlayerData.TutorialData,
					CreatedAt:           session.PlayerData.CreatedAt,
					Status:              session.PlayerData.Status,
					CognitoUsername:     session.PlayerData.CognitoUsername,
					Email:               session.PlayerData.Email,
					ParticleWallet:      session.PlayerData.ParticleWallet,
					ExternalWallet:      session.PlayerData.ExternalWallet,
				}
				err = pubsub.SendToPlayer(session.AccountID.Hex(), data)
				if err != nil {
					log.Error(ctx, "Failed to send player data", err)
				}
				return err
			}
		} else {
			return errors.ErrWalletAlreadyAttched
		}
		return errors.ErrUserNotFound
	}
	return nil
}

func (s *Service) GetUserBalances(wallet string) (*accounts.AccountBalanceResponse, error) {
	balances := make(map[string]string)
	address := common.HexToAddress(wallet)
	balanceChecker, err := blockchain.NewBalanceChecker(s.config.Blockchain.RpcAddresses.Ethereum, s.config.Blockchain.Contracts.LOCG)
	if err != nil {
		return nil, err
	}
	bal, err := balanceChecker.GetEthBalance(address)
	if err != nil {
		return nil, err
	}
	balances["ETH"] = bal.String()

	bal, err = balanceChecker.GetTokenBalance(address)
	if err != nil {
		return nil, err
	}

	balances["LOCG"] = bal.String()

	balanceChecker, err = blockchain.NewBalanceChecker(s.config.Blockchain.RpcAddresses.Ethereum, s.config.Blockchain.Contracts.USDT)
	if err != nil {
		return nil, err
	}
	bal, err = balanceChecker.GetTokenBalance(address)
	if err != nil {
		return nil, err
	}
	balances["USDT"] = bal.String()
	balanceChecker, err = blockchain.NewBalanceChecker(s.config.Blockchain.RpcAddresses.Ethereum, s.config.Blockchain.Contracts.USDC)
	if err != nil {
		return nil, err
	}
	bal, err = balanceChecker.GetTokenBalance(address)
	if err != nil {
		return nil, err
	}
	balances["USDC"] = bal.String()

	balanceChecker, err = blockchain.NewBalanceChecker(s.config.Blockchain.RpcAddresses.Base, s.config.Blockchain.Contracts.BaseLOCG)
	if err != nil {
		return nil, err
	}
	bal, err = balanceChecker.GetEthBalance(address)
	if err != nil {
		return nil, err
	}
	balances["ETHBase"] = bal.String()

	bal, err = balanceChecker.GetTokenBalance(address)
	if err != nil {
		return nil, err
	}

	balances["LOCGBase"] = bal.String()

	balanceChecker, err = blockchain.NewBalanceChecker(s.config.Blockchain.RpcAddresses.Base, s.config.Blockchain.Contracts.BaseUSDC)
	if err != nil {
		return nil, err
	}
	bal, err = balanceChecker.GetTokenBalance(address)
	if err != nil {
		return nil, err
	}
	balances["USDCBase"] = bal.String()

	ctx := context.Background()
	userID, err := s.store.Players.GetAccountIDByWallet(ctx, wallet)
	if err == nil {
		playerData, err := s.store.Players.GetPlayerDataByID(ctx, userID)
		if err == nil {
			for resource, amount := range playerData.Resources {
				balances[fmt.Sprintf("RESOURCE_%d", resource)] = fmt.Sprintf("%d", amount)
			}
		}
	}

	return &accounts.AccountBalanceResponse{
		Balances: balances,
	}, nil
}
