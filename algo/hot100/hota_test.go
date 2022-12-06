package hot

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	ans := twoSum([]int{1, 2, 3}, 3)
	if ans[0] == 0 && ans[1] == 1 {
		fmt.Println("success!")
	} else {
		fmt.Println("fail!")
	}
}
