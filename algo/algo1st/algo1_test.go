package algo1st

import (
	"reflect"
	"strings"
	"testing"
)

func Test(t *testing.T) {

}

func TestMaxIncreaseKeepingSkyline(t *testing.T) {
	test := [][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}
	got := maxIncreaseKeepingSkyline(test)
	want := 35
	if !reflect.DeepEqual(got, want) {
		t.Error("error")
	}
}

func TestToLowerCase(t *testing.T) {
	testStr := "Hello World!"

	got := toLowerCase(testStr)
	want := strings.ToLower(testStr)

	if !reflect.DeepEqual(got, want) {
		t.Error("error")
	}
}
