package gosdk

import (
	"fmt"
	"testing"
)

type TestIntType int32

const (
	V1 TestIntType = 1
	V2 TestIntType = 2
	V3 TestIntType = 3
)

func TestIf(t *testing.T) {
	if V1 > V2 {
		fmt.Println("v1 > v2")
	}
	if V3 >= V2 {
		fmt.Println("v3 >= v2")
	}
	if V2 >= V2 {
		fmt.Println("v2 >= v2")
	}
	if 4 >= V2 {
		fmt.Println("4 >= v2")
	}
}
