package goutils

import "fmt"

func CompareArr(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	m := map[string]int{}
	for _, v := range arr2 {
		m[v] += 1
		fmt.Printf("m[%s] = %d\n", v, m[v])
	}
	fmt.Println("=======================")
	for _, v := range arr1 {
		m[v] -= 1
		fmt.Printf("m[%s] = %d\n", v, m[v])
	}
	fmt.Println("=======================")
	for k, v := range m {
		fmt.Printf("m[%s] = %d\n", k, v)
		if v != 0 {
			return false
		}
	}
	return true
}
