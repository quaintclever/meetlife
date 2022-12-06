package hot

// 1. 两数之和
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if j, ok := m[v]; ok {
			return []int{j, i}
		} else {
			m[target-v] = i
		}
	}
	return nil
}
