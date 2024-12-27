package str

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

// https://leetcode.cn/problems/check-permutation-lcci/

// 先对字符串中的字符排序，再比较
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	return sortStr(s1) == sortStr(s2)
}

func sortStr(s string) string {
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}

// 假定使用的是 ASCII 字符集。用一个整型数组记录每个字符的出现状态。
// 对于第一个字符串，累加出现次数。对于第二个，累减出现次数。
func CheckPermutation2(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	charSet := make([]int, 26)
	for i, char := range s1 {
		charSet[char-'a']++
		charSet[s2[i]-'a']--
	}
	for _, n := range charSet {
		if n != 0 {
			return false
		}
	}
	return true
}

// 改进 CheckPermutation2
func CheckPermutation3(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	charSet := make([]int, 26)
	for _, char := range s1 {
		charSet[char-'a']++
	}
	for _, char := range s2 {
		i := char - 'a'
		charSet[i]--
		if charSet[i] < 0 { // 如果为负，就永不可能变为正了
			return false
		}
	}
	return true
}

func TestSortStr(t *testing.T) {
	assert.Equal(t, "abc", sortStr("cab"))
}
