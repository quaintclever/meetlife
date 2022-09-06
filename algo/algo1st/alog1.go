package algo1st

import (
	"container/list"
	"fmt"
	"math"
	"meetlife/algo/tool"
	"sort"
)

// 8-71. 简化路径
func simplifyPath(path string) string {
	stk := list.New()
	l := len(path)
	for i := 0; i < l; i++ {
		str := ""
		mul := true
		for i++; i < l; i++ {
			if mul = mul && path[i] == '/'; mul {
				continue
			}
			if path[i] == '/' {
				break
			}
			str += string(path[i])
		}
		if i--; str == ".." {
			// 如果是 .. 表示回到上级目录
			if front := stk.Back(); front != nil {
				stk.Remove(front)
			}
		} else if str != "." && str != "" {
			stk.PushBack(str)
		}
	}
	ans := ""
	for stk.Len() > 0 {
		ele := stk.Front()
		ans += "/" + ele.Value.(string)
		stk.Remove(ele)
	}
	if ans == "" {
		ans = "/"
	}
	return ans
}

// 7-1576. 替换所有的问号
func modifyString(s string) string {
	ans := []byte(s)
	l := len(ans)
	pre := byte('a')
	for i := 0; i < l-1; i++ {
		if ans[i] == '?' {
			if pre > 'm' {
				ans[i] = pre - byte(i%2) - 1
				if ans[i] == ans[i+1] {
					ans[i] -= 1
				}
			} else {
				ans[i] = pre + byte(i%2) + 1
				if ans[i] == ans[i+1] {
					ans[i] += 1
				}
			}
		}
		pre = ans[i]
	}
	if ans[l-1] == '?' {
		if pre != 'z' {
			ans[l-1] = pre + 1
		} else {
			ans[l-1] = pre - 1
		}
	}
	return string(ans)
}

var f [2 * 55 * 55][55][55]int
var g [][]int
var n int

// 6-913. 猫和老鼠
func catMouseGame(graph [][]int) int {
	g = graph
	n = len(graph)
	for k := 0; k < n*n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				f[k][i][j] = -1
			}
		}
	}
	return dfs(0, 1, 2)
}

// return 0: draw / 1: mouse / 2: cat
func dfs(k, a, b int) int {
	ans := f[k][a][b]
	if a == 0 {
		ans = 1
	} else if b == a {
		ans = 2
	} else if k >= n*n {
		ans = 0
	} else if ans == -1 {
		if k%2 == 0 { // mouse
			win, draw := false, false
			for _, v := range g[a] {
				switch dfs(k+1, v, b) {
				case 1:
					win = true
					break
				case 0:
					draw = true
				}
			}
			if win {
				ans = 1
			} else if draw {
				ans = 0
			} else {
				ans = 2
			}
		} else { // cat
			win, draw := false, false
			for _, v := range g[b] {
				if v == 0 {
					continue
				}
				switch dfs(k+1, a, v) {
				case 2:
					win = true
					break
				case 0:
					draw = true
				}
			}
			if win {
				ans = 2
			} else if draw {
				ans = 0
			} else {
				ans = 1
			}
		}
	}
	f[k][a][b] = ans
	return ans
}

// 5-507. 完美数
func checkPerfectNumber(num int) bool {
	ans := 1
	limit := int(math.Sqrt(float64(num)))
	for i := 2; i <= limit; i++ {
		if num%i == 0 {
			ans += i + num/i
		}
	}
	return num != 1 && ans == num
}

// 4-825. 适龄的朋友
func numFriendRequests(ages []int) int {
	ans := 0
	l := len(ages)
	sort.Ints(ages)
	cache := make(map[int]int, 10)
	for y := 0; y < l; y++ {
		if i := cache[ages[y]]; i != 0 {
			ans += i
			continue
		}

		temp := 0
		for x := y + 1; x < l; x++ {
			// 把后面的不过的case 也剪掉
			if ages[y] > ages[x]/2+7 {
				temp++
			} else {
				break
			}
		}
		cache[ages[y]] = temp
		ans += temp
	}
	return ans
}

// 3-1518. 换酒问题
func numWaterBottles(numBottles int, numExchange int) int {
	ans := 0 // 喝了多少瓶
	b := 0   // 空瓶
	for numBottles > 0 {
		ans += numBottles               // 开喝
		temp := numBottles + b          // 本次空瓶数
		numBottles = temp / numExchange // 兑换酒
		b = temp % numExchange          // 换完 还剩多少瓶
	}
	return ans
}

// 2-807. 保持城市天际线
func maxIncreaseKeepingSkyline(grid [][]int) int {
	l := len(grid)
	row := make([]int, l, l)
	col := make([]int, l, l)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if grid[i][j] > row[i] {
				row[i] = grid[i][j]
			}
			if grid[i][j] > col[j] {
				col[j] = grid[i][j]
			}
		}
	}
	ans := 0
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if grid[i][j] < row[i] && grid[i][j] < col[j] {
				ans += tool.Min(row[i], col[j]) - grid[i][j]
			}
		}
	}
	fmt.Printf("row:=%v \n col:=%v", row, col)
	return ans
}

// 1-709. 转换成小写字母
func toLowerCase(str string) string {
	ans := []byte(str)
	for i := 0; i < len(str); i++ {
		if ans[i] >= 65 && ans[i] <= 90 {
			ans[i] += 32
		}
	}
	return string(ans)
}
