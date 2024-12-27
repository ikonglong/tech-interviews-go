package str

func replaceSpaces(s string, len int) string {
	numSpaces := 0
	for i := 0; i < len; i++ {
		if s[i] == ' ' {
			numSpaces++
		}
	}
	lenAfterReplaced := len - numSpaces + numSpaces*3 // 一个空格替换为三个字符 %20
	p1 := len - 1
	p2 := lenAfterReplaced - 1
	sBytes := []byte(s)
	for p1 >= 0 && p2 >= 0 {
		if s[p1] == ' ' {
			// 先将 p2 向头部移动两个位置，加上移动前它所在的位置就有三个位置了，够写入 %20 了
			p2 = p2 - 2
			sBytes[p2] = '%'
			sBytes[p2+1] = '2'
			sBytes[p2+2] = '0'
		} else {
			sBytes[p2] = s[p1]
		}
		p2-- // 将 p2 移动至下一个可写入的位置
		p1-- // 将 p1 移动至下一个可读取的位置
	}
	return string(sBytes[:lenAfterReplaced])
}
