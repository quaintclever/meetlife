package sruntime

import (
	"fmt"
	"runtime"
	"testing"
)

func TestRuntimeGetVersion(t *testing.T) {
	fmt.Printf("version:%s\n", runtime.Version())
}
