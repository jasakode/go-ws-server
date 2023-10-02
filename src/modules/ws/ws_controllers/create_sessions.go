package wscontrollers

import (
	"encoding/json"
	"fmt"
	"go-ws-server/src/config"
	"go-ws-server/src/middleware"
	"go-ws-server/src/middleware/req"
	"go-ws-server/src/middleware/res"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type RXMessage struct {
	Action        string `json:"action"`
	Message       string `json:"message"`
	Authorization string `json:"authorization"`
}

var Connections = make(map[*websocket.Conn]bool)

var CreateSessions = func(c *gin.Context) {
	w := c.Writer
	r := c.Request
	authorization := r.Header.Get("authorization")
	if authorization == "" {
		fmt.Println("Error : Unauthorized")
		resp := res.Message(false, "Unauthorized")
		res.Response(w, resp)
		return
	}
	dataBody, errGenBosy := json.Marshal(map[string]interface{}{
		"token": authorization,
	})
	if errGenBosy != nil {
		fmt.Println(errGenBosy.Error())
		resp := res.Message(false, errGenBosy.Error())
		res.Response(w, resp)
		return
	}
	respon, errFetch := req.FetchApi("POST", "http://127.0.0.1:5000/api/v1/auth",
		[]req.Headers{}, dataBody)
	if errFetch != nil {
		fmt.Println(errFetch.Error())
		resp := res.Message(false, errFetch.Error())
		res.Response(w, resp)
		return
	}
	fmt.Println("TEST")
	HandleWebSocket(c.Writer, c.Request, respon)
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request, user interface{}) {
	conn, err := middleware.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	middleware.AddSession(&middleware.WSSessions{
		ConnectinIn:    uint(time.Now().Unix()),
		ConnectionTime: 0,
		ExpiredIn:      uint(time.Now().Unix()) + config.Config.EXPIRED_CONNECTIONS,
		User:           &user,
		Socket:         conn,
	})

}
