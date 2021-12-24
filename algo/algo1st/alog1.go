package algo1st

import (
	"fmt"
	"meetlife/algo/tool"
)

// 1518. 换酒问题
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

// 807. 保持城市天际线
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

// 709. 转换成小写字母
func toLowerCase(str string) string {
	ans := []byte(str)
	for i := 0; i < len(str); i++ {
		if ans[i] >= 65 && ans[i] <= 90 {
			ans[i] += 32
		}
	}
	return string(ans)
}