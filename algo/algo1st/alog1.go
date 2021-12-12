package algo1st

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
