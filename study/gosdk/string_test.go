package gosdk

import (
	"fmt"
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
