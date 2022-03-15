package gosdk

import (
	"fmt"
	"strconv"
	"testing"
)

type StringType string

const (
	St1 StringType = "test1"
	St2 StringType = "test2"
)

func TestCaseString(t *testing.T) {
	s1 := "test1"
	s2 := "test2"
	s3 := "test3"
	fmt.Println(St1 == StringType(s1))
	fmt.Println(St1 == StringType(s2))
	fmt.Println(St2 == StringType(s3))
	fmt.Println(StringType(s1))
	fmt.Println(StringType(s2))
	fmt.Println(StringType(s3))
}

func TestInt64ToString(t *testing.T) {
	var i = 100
	// conversion from int64 to string yields a string of one rune, not a string of digits (did you mean fmt.Sprint(x)?)
	//fmt.Println(string(i))
	fmt.Println(strconv.Itoa(i))
	fmt.Println(strconv.FormatInt(int64(i), 10))

	str := "1000"
	parseInt, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(parseInt)

}
