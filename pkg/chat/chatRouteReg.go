package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/chat/handler"
)

// Chat handles the websocket connection in the chat page.
func (c *Chat) Chat(ctx *gin.Context) {
	handler.HandleWebSocketConnection(ctx, c.client, c.userClient)
}

// VideoCall handles initiating a video call.
func (c *Chat) VideoCall(ctx *gin.Context) {
	handler.StartVideoCall(ctx, c.client)
}

// ChatPage handles to load the chat page
func (c *Chat) ChatPage(ctx *gin.Context) {
	handler.ChatPage(ctx, c.client)
}
