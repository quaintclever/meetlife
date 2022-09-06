package gosdk

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestTimeFmt(t *testing.T) {
	fmt.Println("hello time fmt")
	now := time.Now()
	rand.Seed(now.UnixNano())
	var randStr string
	if rn := rand.Intn(100); rn < 10 {
		randStr = fmt.Sprintf("0%d", rn)
	} else {
		randStr = fmt.Sprintf("%d", rn)
	}
	sprintf := fmt.Sprintf("R%s%s", now.Format("20060102150405"), randStr)

	fmt.Println("==============================")
	fmt.Println(sprintf)
}

func TestTimeAddTime(t *testing.T) {
	now := time.Now()
	dd, _ := time.ParseDuration("168h")
	dd1 := now.Add(dd)
	fmt.Println(dd1)
}
