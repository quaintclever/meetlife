package goutils

import (
	"fmt"
	"testing"
)

func TestGetGid(t *testing.T) {
	gid := GetGid()
	fmt.Println(gid)
}
