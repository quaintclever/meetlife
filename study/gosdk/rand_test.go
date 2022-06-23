package gosdk

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRand(t *testing.T) {
	for i := 0; i < 20; i++ {
		n := rand.Int63n(1000)
		fmt.Println(n)
	}
}
