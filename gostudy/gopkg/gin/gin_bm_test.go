package gin

import (
	"net/http"
	"testing"
)

func BenchmarkRouterCall(b *testing.B) {
	b.SetParallelism(100)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_, _ = http.Get("http://localhost:8080")
		}
	})
}
