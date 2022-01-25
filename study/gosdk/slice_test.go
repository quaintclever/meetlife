package gosdk

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortSlice(t *testing.T) {
	sl := []int{9, 2, 3, 5, 4, 7}
	// asc 升序
	sort.SliceStable(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})
	fmt.Printf("%v \n", sl)
	// desc 降序
	sort.SliceStable(sl, func(i, j int) bool {
		return sl[i] > sl[j]
	})
	fmt.Printf("%v", sl)

}
