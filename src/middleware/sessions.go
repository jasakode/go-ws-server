package middleware

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WSSessions struct {
	ConnectinIn    uint
	ConnectionTime uint
	ExpiredIn      uint
	User           *interface{}
	Socket         *websocket.Conn
}

var Sessions []WSSessions

func AddSession(s *WSSessions) {
	Sessions = append(Sessions, *s)
}
