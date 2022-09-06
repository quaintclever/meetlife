package goutils

import (
	"fmt"
	"testing"
)

func TestArrayEq(t *testing.T) {
	fmt.Println(compareArr([]string{"a", "b"}, []string{"b", "a"}))
	fmt.Println(compareArr([]string{"a", "b", "b"}, []string{"b", "b", "a"}))
	fmt.Println(compareArr([]string{"a", "b", "c"}, []string{"d", "b", "a"}))
	fmt.Println(compareArr([]string{"a", "b", "b", "b"}, []string{"b", "a", "a", "a"}))
}
