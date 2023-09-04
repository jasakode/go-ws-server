package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type WSSessions struct {
	Authorization  string
	ConnectionDate int64
	Socket         *websocket.Conn
}

type RXMessage struct {
	Action        string `json:"action"`
	Message       string `json:"message"`
	Authorization string `json:"authorization"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Set to true to allow all origins, adjust as needed
	},
}

// SESSION MEMORY

var Connections = make(map[*websocket.Conn]bool)
var Sessions []WSSessions

// ...

func GetSessions(Authorization string) []WSSessions {
	var respon []WSSessions
	for _, SC := range Sessions {
		if SC.Authorization == Authorization {
			respon = append(respon, SC)
		}
	}
	return respon
}
func DeletedAction(Authorization string) {
	var respon []WSSessions
	for _, SC := range Sessions {
		if SC.Authorization != Authorization {
			respon = append(respon, SC)
		}
	}
	Sessions = respon
}
func AddSessions(Authorization string, s websocket.Conn) {
	ss := GetSessions(Authorization)
	for i := 0; i < len(ss); i++ {
		// lastTime := time.Now().Unix() - 259200000
		// if ss[i].ConnectionDate < lastTime {
		// 	DeletedAction(Authorization)
		// }
		if ss[i].Authorization == Authorization {
			DeletedAction(Authorization)
		}
	}
	Sessions = append(Sessions, WSSessions{
		Authorization:  Authorization,
		Socket:         &s,
		ConnectionDate: time.Now().Unix(),
	})
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	Connections[conn] = true

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			delete(Connections, conn)
			return
		}
		var rxmsg RXMessage
		errd := json.Unmarshal(p, &rxmsg)
		if errd != nil {
			fmt.Println(errd.Error())
		}

		// VALIDASI TOKEN TERHUBUNG DENGAN USER YANG ADA DI DB ATAU TIDAK -- todo
		switch rxmsg.Action {
		case "authentication":
			AddSessions(rxmsg.Authorization, *conn)
		}
	}
}
