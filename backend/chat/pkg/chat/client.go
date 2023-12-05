package chat

import (
	"encoding/json"
	"foip/chat/pkg/log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	CLIENT_WRITE_DEADLINE = 10 * time.Second
	CLIENT_PONG_WAIT_TIME = 60 * time.Second
	CLIENT_PING_PERIOD    = (CLIENT_PONG_WAIT_TIME * 9) / 10

	MESSAGE_MAX_SIZE = 1024

	TIME_STAMP_LAYOUT = "2006-01-02 15:04:05"
)

type Message struct {
	Table   string `json:"table"`
	Message string `json:"message"`
	Time    string `json:"time,omitempty"`
	User    string `json:"user"`
	Data    []byte `json:"-"`
}

type Client struct {
	manager *BloadcastManager

	conn    *websocket.Conn
	message chan []byte
}

func (c *Client) writer() {
	ticker := time.NewTicker(CLIENT_PING_PERIOD)
	defer func() {
		log.Call.Info("closed connection")
		c.conn.Close()
	}()

	for {
		select {
		case _message, ok := <-c.message:
			if !ok {
				if err := c.conn.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Call.Error(err)
				}
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Call.Error(err)
				return
			}
			if _, err := w.Write(_message); err != nil {
				log.Call.Error(err)
				goto close
			}

			//Add queued chat messages to the current websocket message.
			//ref: https://github.com/gorilla/websocket/issues/639
			for i, l := 0, len(c.message); i < l; i++ {
				if _, err := w.Write(<-c.message); err != nil {
					log.Call.Error(err)
					goto close
				}
			}

		close:
			if err := w.Close(); err != nil {
				log.Call.Error(err)
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(CLIENT_WRITE_DEADLINE))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Call.Error(err)
				return
			}
		}
	}
}

func (c *Client) reader() {
	defer func() {
		c.manager.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(MESSAGE_MAX_SIZE)
	c.conn.SetReadDeadline(time.Now().Add(CLIENT_PONG_WAIT_TIME))
	c.conn.SetPongHandler(
		func(string) error {
			if err := c.conn.SetReadDeadline(time.Now().Add(CLIENT_PONG_WAIT_TIME)); err != nil {
				log.Call.Error(err)
			}
			return nil
		},
	)
	for {
		_, _message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Call.Error(err)
			}
			return
		}

		message := Message{}
		if err := json.Unmarshal(_message, &message); err != nil {
			log.Call.Error(err)
			return
		}

		message.Time = time.Now().Format(TIME_STAMP_LAYOUT)
		if _message, err = json.Marshal(message); err != nil {
			log.Call.Error(err)
			return
		}
		message.Data = _message

		c.manager.broadcast <- message
	}
}
