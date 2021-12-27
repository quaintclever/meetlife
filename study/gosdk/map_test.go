package gosdk

import (
	"fmt"
	"testing"
)

func TestMapInt(t *testing.T) {
	// 测试 map int 零值
	m := make(map[int]int, 10)
	fmt.Printf("step1:map:%v\n", m)
	m[0] = 0
	m[1] = 1
	fmt.Printf("step2:map:%v\n", m)
	fmt.Println(m[0])
	fmt.Println(m[1])
	fmt.Println(m[10])
}
