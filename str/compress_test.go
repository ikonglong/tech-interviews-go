package str

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func compressString(s string) string {
	if len(s) == 0 {
		return s
	}

	bytes := make([]byte, 0, len(s))
	// 循环开始前，先把第一个字符放入压缩结果中，且计数初始化为 1，因为第一个字符已经出现了一次。
	// 为什么这么干？因为脑子蠢，总觉得第二次发现才应该开始累加，吐血。这个愚蠢的认识导致整个实现
	// 很不整洁，例如循环结束后还要再打个补丁。
	bytes = append(bytes, s[0])
	charCount := 1
	for i, c := range s {
		if i == 0 {
			continue
		}

		if bytes[len(bytes)-1] == byte(c) {
			charCount++
		} else { // 遍历到一个不同的字符时才将上一个字符的计数追加到结果中
			// bytes = append(bytes, byte(charCount))
			bytes = append(bytes, strconv.Itoa(charCount)...)
			bytes = append(bytes, byte(c))
			charCount = 1
		}
	}

	// 对于最后一个字符，由于在循环中无法触发将其写入到压缩结果中，因此循环结束后再处理。
	// 但是这样的实现总感觉到处打补丁，不整洁。
	bytes = append(bytes, strconv.Itoa(charCount)...)
	if len(bytes) < len(s) {
		return string(bytes[:len(bytes)])
	} else {
		return s
	}
}

// compressString 的改进版本
func compressString2(s string) string {
	if len(s) == 0 {
		return s
	}

	compressed := make([]byte, 0, len(s))
	charCount := 0
	for i, c := range s {
		// 只要在看到不同字符时先将 charCount 重置为 0，就可以
		// 认为 charCount 是对当前字符的计数
		charCount++

		// 看看下一个字符是否相同。若不同，就将当前字符和计数写入结果，
		// 再将 charCount 重置为 0，为下一个字符的处理做准备。注意，
		// 还要考虑当前已走到末尾的情形。
		if i+1 >= len(s) || s[i+1] != byte(c) {
			compressed = append(compressed, byte(c))
			compressed = append(compressed, strconv.Itoa(charCount)...)
			charCount = 0
		}
	}
	if len(compressed) < len(s) {
		return string(compressed[:len(compressed)])
	}
	return s
}

func TestStrBuilder(t *testing.T) {
	builder := strings.Builder{}
	builder.WriteByte(65)
	fmt.Println(builder.String())
}
