package net_test

import (
	"net/http"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

func TestServer(t *testing.T) {
	http.HandleFunc("/socket", socketHandler)
	http.HandleFunc("/", home)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		logrus.Fatalf("listen err:%s", err.Error())
	}
}

func TestClient(t *testing.T) {
	done = make(chan interface{})
	interrupt = make(chan os.Signal)

	signal.Notify(interrupt, os.Interrupt)

	socketUrl := "ws://localhost:8080" + "/socket"
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		logrus.Fatal("error conn to server:", err)
	}
	defer conn.Close()
	go receiveHandler(conn)

	for {
		select {
		case <-time.After(time.Duration(3) * time.Second):
			err := conn.WriteMessage(websocket.TextMessage, []byte("ping! hello from client"))
			if err != nil {
				logrus.Info("error during writing to websocket:", err)
				return
			}
		case <-interrupt:
			logrus.Info("[client] received interrupt signal. closing all pending connections.")
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				logrus.Info("error during closing websocket:", err)
				return
			}

			select {
			case <-done:
				logrus.Info("[client] receiver channel close! exiting...")
			case <-time.After(time.Duration(1) * time.Second):
				logrus.Info("[client] timeout in closing receiving channel. exiting...")
			}
			return
		}
	}
}
