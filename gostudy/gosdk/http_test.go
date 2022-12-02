package gosdk

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"testing"
)

// 测试 health 方法是否可以被 80 8080 同时监听
func TestHttpHandler(t *testing.T) {
	// service server
	svc := &http.Server{
		Addr: ":80",
	}
	// health check server
	hc := &http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "ok")
	})

	// service server
	go func() {
		if err := svc.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("service server exit error: %v", err)
		}
	}()

	// health check server
	go func() {
		if err := hc.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("health check server exit error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	svc.Close()
	hc.Close()
}
