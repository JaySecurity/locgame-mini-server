package store

type FriendsData struct {
	Friends         map[string]struct{}
	IncomingInvites map[string]struct{}
	OutgoingInvites map[string]struct{}
}
