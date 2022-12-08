package hot

import . "github.com/quaintclever/meetlife/algo/tool/leetcode"

// 1. 两数之和
// https://leetcode.cn/problems/two-sum/
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

/**
 *  2. 两数相加
 * https://leetcode.cn/problems/add-two-numbers/
 *
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	ans := root
	x := 0
	for l1 != nil || l2 != nil {
		if x == 0 {
			if l1 == nil {
				root.Next = l2
				break
			} else if l2 == nil {
				root.Next = l1
				break
			} else {
				if l1.Val+l2.Val > 9 {
					x = 1
				}
				root.Next = &ListNode{
					Val:  l1.Val + l2.Val - x*10,
					Next: nil,
				}
				l1 = l1.Next
				l2 = l2.Next
				root = root.Next
			}
		} else {
			if l1 == nil {
				if l2.Val+1 > 9 {
					root.Next = &ListNode{}
					l2 = l2.Next
					root = root.Next
				} else {
					l2.Val = l2.Val + 1
					root.Next = l2
					x = 0
					break
				}
			} else if l2 == nil {
				if l1.Val+1 > 9 {
					root.Next = &ListNode{}
					l1 = l1.Next
					root = root.Next
				} else {
					l1.Val = l1.Val + 1
					root.Next = l1
					x = 0
					break
				}
			} else {
				if l1.Val+l2.Val+x > 9 {
					x = 1
				} else {
					x = 0
				}
				root.Next = &ListNode{
					Val:  l1.Val + l2.Val + 1 - x*10,
					Next: nil,
				}
				l1 = l1.Next
				l2 = l2.Next
				root = root.Next
			}
		}
	}
	if x == 1 {
		root.Next = &ListNode{1, nil}
	}
	return ans.Next
}
