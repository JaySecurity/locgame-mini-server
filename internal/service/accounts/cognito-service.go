package accounts

// "fmt"
// "locgame-mini-server/pkg/dto/errors"

// func (s *Service) SendLoginEmail(ctx context.Context, request *accounts.LoginEmailRequest) (*accounts.LoginEmailResponse, error) {
// 	valid := common.ValidateEmail(request.Email)
// 	if !valid {
// 		return nil, errors.ErrUserNotConfirmed
// 	}

// 	err := s.cognitoInstance.GetUserByEmail(ctx, request.Email)
// 	if err != nil {
// 		log.Error("GetUserByEmail cognito error: ", err)
// 		err := s.cognitoInstance.Signup(request.Email)
// 		if err != nil {
// 			log.Error("SendLoginEmail cognito signup error: ", err)
// 			return nil, errors.ErrUnexpectedError
// 		}
// 	}

// 	challenge, sessionStr := s.cognitoInstance.SendEmail(request.Email)

// 	return &accounts.LoginEmailResponse{
// 		ChallengeName: challenge,
// 		Session:       sessionStr,
// 	}, nil
// }

// func (s *Service) VerifyLoginEmail(ctx context.Context, request *accounts.VerifyLoginEmailRequest) (*accounts.LoginResponse, error) {
// 	accessToken, IdToken, err := s.cognitoInstance.VerifyEmail(request.Email, request.Code, request.ChallengeName, request.Session)
// 	if err != nil {
// 		log.Error(fmt.Sprintf("failed to exchange token: %v", err))
// 		return nil, err
// 	}

// 	return s.createCognitoResponse(ctx, accessToken, IdToken, request.Email)
// }

// func (s *Service) SocialLogin(ctx context.Context, request *accounts.VerifySocialLoginRequest) (*accounts.LoginResponse, error) {
// 	var l common.LoginRequest
// 	var u common.UserData

// 	common.ProcessMessage(request.Code, &l, &u)

// 	err := s.cognitoInstance.HandleSocialAuth(&l, &u)
// 	if err != nil {
// 		log.Error("SocialLogin: ", err)
// 		return nil, err
// 	}

// 	return s.createCognitoResponse(ctx, u.AccessToken, u.IdToken, "social")
// }

// func (s *Service) createCognitoResponse(ctx context.Context, accessToken, idToken string, email string) (*accounts.LoginResponse, error) {
// 	cognitoUsername, err := s.cognitoInstance.GetUserByToken(ctx, accessToken)
// 	if err != nil {
// 		log.Error("createCognitoResponse s.cognitoInstance.getUserByToken: ", err)
// 		return nil, err
// 	}

// 	session, err := s.NewSessionBuilder(ctx, s.config).SetCognitoUsername(cognitoUsername).SetEmailAddress(email).CreateSession().Build()
// 	if err != nil {
// 		log.Error("createCognitoResponse session: ", err)
// 		return nil, err
// 	}

// 	sd, err := s.createLoginResponse(ctx, session)
// 	if err != nil {
// 		log.Error("createCognitoResponse createLoginResponse: ", err)
// 		return nil, err
// 	}
// 	sd.IdToken = idToken

// 	return sd, err
// }
