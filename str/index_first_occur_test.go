package str

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

// 暴力计算
func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	for i, c := range haystack {
		if byte(c) != needle[0] {
			continue
		}
		j := i + 1
		k := 1
		// 第一次写时遗漏了 j < len(haystack)
		for ; j < len(haystack) && j < i+len(needle) && k < len(needle); j++ {
			if haystack[j] != needle[k] {
				break
			}
			k++
		}
		if k == len(needle) { // 表明找到 needle 了
			return i
		}

		if len(haystack)-1-i < len(needle) {
			break
		}
	}

	return -1
}

func TestStrStr(t *testing.T) {
	assert.Equal(t, 0, strStr("sadbutsad", "sad"))
	assert.Equal(t, -1, strStr("bbaa", "aab"))
}
