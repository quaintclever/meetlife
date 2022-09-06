package gosdk

import (
	"fmt"
	"testing"
)

func TestIntPoint(t *testing.T) {
	// 定义 int变量
	intP := 3
	// 定义 int 指针
	var replicas *int
	replicas = &intP
	// 改变指针地址里的值
	*replicas = 5
	fmt.Println(*replicas)
}
