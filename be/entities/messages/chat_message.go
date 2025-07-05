package messages

type ChatMessage struct {
	SenderID uint64
	RoomID   uint64
	Message  string
}
