package gosdk

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	a := 1
	switch a {
	case 1:
		fmt.Println("1")
		// error: Fallthrough statement out of place
		//if true {
		//	fallthrough
		//}
		fmt.Println("????")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}
}
