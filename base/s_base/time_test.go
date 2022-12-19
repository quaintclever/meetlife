package sbase

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestZoneFmt(t *testing.T) {
	val := "2006-12-19T14:09:02+0800"
	// ti := time.Now()
	// fmt.Println(ti.Local().Format("2006-01-02T15:04:05Z0700"))
	res, err := time.ParseInLocation("2006-01-02T15:04:05Z0700", val, time.Local)
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		fmt.Println(res)
	}
}

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
