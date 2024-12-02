package dto

// Message represents the chat details model

// type Message struct {
// 	UserID     string `json:"userId"` // Use string if the JSON has string representation for userId
// 	ReceiverID string `json:"receiverId"`
// 	Message    string `json:"content"`
// 	CallType   string `json:"callType,omitempty"`
// 	RoomID     string `json:"roomId,omitempty"`
// 	Duration   int    `json:"duration,omitempty"`
// }

type Message struct {
	UserID     uint   `json:"userId"`
	ReceiverID uint   `json:"receiverId"`
	Message    string `json:"content"`
	CallType   string `json:"callType,omitempty"`
	RoomID     string `json:"roomId,omitempty"`
	Duration   int    `json:"duration,omitempty"`
}
