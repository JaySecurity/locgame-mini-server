package pubsub

var subscribeHandlers map[string]Handler
var queueSubscribeHandlers map[string]Handler

func RegisterHandler(handler Handler) {
	if subscribeHandlers == nil {
		subscribeHandlers = make(map[string]Handler)
	}
	subscribeHandlers[getSubjectNameByHandler(handler)] = handler
}

func getSubjectNameByHandler(handler Handler) string {
	return handler.GetDataType().String()
}

func RegisterQueueHandler(handler UnsubscribableHandler) {
	if queueSubscribeHandlers == nil {
		queueSubscribeHandlers = make(map[string]Handler)
	}
	queueSubscribeHandlers[getSubjectNameByHandler(handler)] = handler
}

func GetHandlers() map[string]Handler {
	return subscribeHandlers
}

func GetQueueHandlers() map[string]Handler {
	return queueSubscribeHandlers
}
