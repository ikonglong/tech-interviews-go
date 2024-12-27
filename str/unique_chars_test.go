package str

import (
	"fmt"
	"testing"
)

// 使用一个 bool 数组，长度等于字符串使用的字符集大小。第 i 个标记表示
// 是否含有字符集第 i 个字符。第一次出现时，将对应标记置为 true。第二次
// 出现时，直接返回 false
func isUnique(str string) bool {
	// 假定字符串只包含 ASCII 字符集中的字符
	if len(str) > 128 {
		return false
	}

	charPresent := make([]bool, 128)
	for _, char := range str {
		if charPresent[char] {
			return false
		}
		charPresent[char] = true
	}
	return true
}

// 使用位向量（bit vector），将空间占用减少为原来的 1/8
func isUnique2(str string) bool {
	// 假定字符串只包含 ASCII 字符集中的字符
	if len(str) > 128 {
		return false
	}

	var checker int = 0
	for _, ch := range str {
		val := ch - 'a'
		if checker&(1<<val) != 0 {
			return false
		}
		checker |= 1 << val
	}
	return true
}

func TestBitOp(t *testing.T) {
	for i := 0; i < 127; i++ {
		fmt.Printf("%d ", 1<<i)
	}
}
