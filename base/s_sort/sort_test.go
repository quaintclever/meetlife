package ssort

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestSort(t *testing.T) {

	list := []string{"TE", "RE", "ME", "AE", "WE", "ZE", "AE", "ZE"}
	listSelf := []string{"RE", "TE", "AE"}

	sort.Slice(listSelf, func(i, j int) bool {
		return strings.Compare(listSelf[i], listSelf[j]) < 0
	})

	sm := map[string]string{}
	for _, v := range list {
		sm[v] = "*" + v
	}

	// 聚合成list
	var ans []string
	for _, v := range listSelf {
		// 自己的排到前面
		ans = append(ans, v)
		delete(sm, v)
	}

	var key []string
	for k := range sm {
		key = append(key, k)
	}
	sort.Strings(key)
	for _, k := range key {
		ans = append(ans, sm[k])
	}
	fmt.Printf("%v\n", ans)
}
