package gin

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var ati atomic.Int64

// 当第一个没有返回的时候，第二个，第三个也会阻塞。
func TestGinRouterSleep(t *testing.T) {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		ati.Add(1)
		if ati.Load()%2 == 0 {
			time.Sleep(time.Millisecond * 3000)
		}
		c.JSON(
			200,
			map[string]interface{}{
				"currentNum": ati.Load(),
			},
		)
	})
	r.Run()
}

// 测试 gin 和 http 是否可以同时监听到 health 方法
// 请参考 http.server.go  => func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request)
// http://localhost:8080/  => Hello Gin
// http://localhost:8079/health  => ok
// http://localhost:8080/health  => 404 page not found
func TestGinAndHttpHandler(t *testing.T) {
	// service server
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Gin")
	})

	// health check server
	hc := &http.Server{
		Addr:    ":8079",
		Handler: nil,
	}
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "ok")
	})

	// health check server
	go func() {
		if err := hc.ListenAndServe(); err != nil && errors.Cause(err) != http.ErrServerClosed {
			logrus.Fatalf("health check server exit error: %v", err)
		}
	}()

	r.Run(":8080")
	hc.Close()
}

// 测试 gin handler 最大数量
// Group的 handler + Route 的 handler 最大不超过63 (测试是 不能超过61个)
// 因为有2个，在这里 engine.Use(Logger(), Recovery())
func TestGinMixGroupHandler(t *testing.T) {
	// service server
	r := gin.Default()

	ng := r.Group("/v1",
		ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler,
		ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler,
		ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler,
		ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler,
		ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler,
		ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler,
		ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler,
		ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler)

	// http://localhost:8080/v1
	// ng.GET("/",
	// 	// ginHelloHandler,
	// 	ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler,
	// 	ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler,
	// 	ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler,
	// 	ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler,
	// 	ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler,
	// 	ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler,
	// 	ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler)

	// http://localhost:8080/v1/hello/1/test/test/list
	ng.GET("/hello/:id/test/:test/list",
		// ginHelloHandler,
		ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler,
		ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler,
		ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler,
		ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler,
		ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler,
		ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler,
		ginHelloHandler, ginHelloHandler, ginHelloHandler, ginHelloHandler)

	r.Run()
}

func ginHelloHandler(c *gin.Context) {
	c.String(200, "hello ")
}
