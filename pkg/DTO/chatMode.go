package dto

type Message struct {
	UserID     uint   `json:"userId"`
	ReceiverID uint   `json:"receiverId"`
	Message    string `json:"content"`
	CallType   string `json:"callType,omitempty"`
	RoomID     string `json:"roomId,omitempty"`
	Duration   int    `json:"duration,omitempty"`
}
