package str

import (
	"fmt"
	"testing"
)

func TestGenRandomStr(t *testing.T) {
	str := genRandomStr(10)
	fmt.Println(str)
}
