package gosdk

import (
	"fmt"
	"testing"
)

func TestMapInt(t *testing.T) {
	// 测试 map string 零值
	m1 := make(map[string]string, 10)
	fmt.Printf("step1:map:%v\n", m1)
	m1["test"] = "test-desc"
	m1["test2"] = "test-desc2"
	fmt.Printf("step1:map:%v\n", m1)
	fmt.Println(m1["test"])
	fmt.Println(m1["test2"])
	fmt.Println(m1["hello"])
	fmt.Println(m1["hello"] == "")

	fmt.Println("====================================")

	// 测试 map int 零值
	m0 := make(map[int]int, 10)
	fmt.Printf("step1:map:%v\n", m0)
	m0[0] = 0
	m0[1] = 1
	fmt.Printf("step2:map:%v\n", m0)
	fmt.Println(m0[0])
	fmt.Println(m0[1])
	fmt.Println(m0[10])
}
