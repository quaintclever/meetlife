package algo1st

import (
	"reflect"
	"strings"
	"testing"
)

func TestToLowerCase(t *testing.T) {
	testStr := "Hello World!"

	got := toLowerCase(testStr)
	want := strings.ToLower(testStr)

	if !reflect.DeepEqual(got, want) {
		t.Error("error")
	}
}
