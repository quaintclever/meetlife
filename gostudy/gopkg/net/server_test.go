package net

import (
	"net/http"
	"testing"

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
