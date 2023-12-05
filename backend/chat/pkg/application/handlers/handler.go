package handler

import (
	"foip/chat/pkg/chat"
	"foip/chat/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		//TODO: write core here for check origin.
		return true
	},
}

type Controller struct {
	manager *chat.BloadcastManager
}

func New(manager *chat.BloadcastManager) *Controller {
	return &Controller{
		manager: manager,
	}
}

func (c *Controller) Handler(cxt *gin.Context) {
	//upgrade client connection protocol to WebSocket from HTTP.
	ws, err := upgrader.Upgrade(cxt.Writer, cxt.Request, nil)
	if err != nil {
		log.Call.Error(err)
		return
	}

	c.manager.RegisterNewClient(ws)
}
