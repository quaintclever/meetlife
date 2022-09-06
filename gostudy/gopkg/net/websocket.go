package net

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// =================== server ===================
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Fatal("error during connection upgradation", err)
	}
	defer conn.Close()

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			logrus.Error("error during msg reading:", err)
			break
		}
		logrus.Infof("[server] received: %s", msg)

		rmsg := "pong! hello from server"
		err = conn.WriteMessage(mt, []byte(rmsg))
		if err != nil {
			logrus.Error("error during msg write:", err)
			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index Page")
}

// =================== client ===================
var done chan interface{}
var interrupt chan os.Signal

func receiveHandler(conn *websocket.Conn) {
	defer conn.Close()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			logrus.Fatal("conn read msg err:", err)
		}
		logrus.Infof("[client] received msg:%s", msg)
	}
}
