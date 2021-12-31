package algo1st

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	got := 1
	want := 1
	if !reflect.DeepEqual(got, want) {
		t.Error("error")
	}
}

func TestCheckPerfectNumber(t *testing.T) {
	cs := []int{28, 7, 6}
	want := []bool{true, false, true}
	for i, c := range cs {
		got := checkPerfectNumber(c)
		if !reflect.DeepEqual(got, want[i]) {
			t.Errorf("error, req:%v, want:%v, got:%v", c, want[i], got)
		}
	}
}

func TestNumFriendRequests(t *testing.T) {
	cs := [][]int{{16, 16}, {16, 17, 18}, {20, 30, 100, 110, 120}, {16, 16, 16, 16, 16, 16, 16, 16}}
	want := []int{2, 2, 3, 56}
	for i, c := range cs {
		got := numFriendRequests(c)
		if !reflect.DeepEqual(got, want[i]) {
			t.Errorf("error, req:%v, want:%v, got:%v", c, want[i], got)
		}
	}
}

func TestNumWaterBottles(t *testing.T) {
	got := numWaterBottles(15, 4)
	want := 19
	fmt.Println(got)
	if !reflect.DeepEqual(got, want) {
		t.Error("error")
	}
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
