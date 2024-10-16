package middleware

// type CheckSessionMiddleware struct {
// 	handler  network.ServerHandler
// 	sessions *sessions.SessionStore
// }

// func NewCheckSessionMiddleware(handler network.ServerHandler, sessions *sessions.SessionStore) *CheckSessionMiddleware {
// 	return &CheckSessionMiddleware{
// 		handler:  handler,
// 		sessions: sessions,
// 	}
// }

// func (h *CheckSessionMiddleware) Serve(client *network.Client, packet *network.Packet, arg interface{}) {
// 	/*
// 		Do not require a session for methods:
// 			10000: "GetConfigs"
// 			10001: "AuthToken"
// 			10002: "Web3ChallengeRequest"
// 			10003: "Web3Authorize"
// 			10051  "SendLoginEmail"
// 			10052  "VerifyLoginEmail"
// 			10053  "SocialLogin"
// 			10054  "GetUserBalances"
// 	*/

// 	if (packet.MethodID >= 10000 && packet.MethodID <= 10004) || (packet.MethodID >= 10051 && packet.MethodID <= 10054) {
// 		h.handler.Serve(client, packet, arg)
// 	} else if session := h.sessions.TryGet(client.Context()); session == nil {
// 		err := errors.ErrSessionNotFound
// 		packet.Payload = nil
// 		packet.Error = &network.Error{ErrorCode: errors.ErrorsByCode[err], Description: err.Error()}
// 		err = client.Stream.WritePacket(packet)
// 		if err != nil {
// 			log.Error("Failed to send error to client:", err)
// 			client.Stream.Close()
// 		}
// 	} else {
// 		h.handler.Serve(client, packet, arg)
// 	}
// }

// func (h *CheckSessionMiddleware) GetArgTypeByMethodID(methodID uint16) reflect.Type {
// 	return h.handler.GetArgTypeByMethodID(methodID)
// }

// func (h *CheckSessionMiddleware) GetMethodNameByID(methodID uint16) string {
// 	return h.handler.GetMethodNameByID(methodID)
// }

// func (h *CheckSessionMiddleware) Validate(methodID uint16) bool {
// 	return h.handler.Validate(methodID)
// }
