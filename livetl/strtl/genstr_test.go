package strtl

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGenRandomStr(t *testing.T) {
	str := GenRandomStr(10)
	fmt.Println(str)
	fmt.Println("======================")
	str2 := GenTimeRandStr(TestType)
	fmt.Println(str2)
}

type FlagType string

const (
	TestType FlagType = "T"
)

func GenTimeRandStr(flag FlagType) string {
	// 运维单 session 生成
	now := time.Now()
	rand.Seed(now.UnixNano())
	rn := rand.Intn(1000)
	mixSession := fmt.Sprintf("%s%s%03d", flag, now.Format("20060102150405"), rn)
	return mixSession
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().Unix())
}

// 生成随机字符串
func GenRandomStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
