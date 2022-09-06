package goutils

import (
	"fmt"
	"testing"
)

func TestGenRandomStr(t *testing.T) {
	str := GenRandomStr(10)
	fmt.Println(str)
	fmt.Println("======================")
	str2 := GenTimeRandStr(TestType)
	fmt.Println(str2)
}
