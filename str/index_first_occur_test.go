package str

import (
	"github.com/stretchr/testify/assert"

	"bytes"
	"crypto/sha256"
	"math"
	"testing"
)

// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/

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

// 使用 RK 算法的第一个版本。使用自定义的 hash 算法计算 hash 值，比较 hash 值。
//
// 比较主串中的子串和模式串（needle）时比较 hash 值。
// 假设字符串使用的字符集是 26 个小写字母，采用如下 hash 算法：
// 这 26 个小写字母对应 26 进制的基本数字符号，分别表示 0-25。计算
// 子串和模式串的 hash 值，即 26 进制数的 10 进制表示。
//
// hash 值超出了数据类型范围怎么办？
func strStr2(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	for i, _ := range haystack {
		if i+len(needle)-1 >= len(haystack) {
			return -1
		}
		if base26Hash(haystack[i:i+len(needle)]) == base26Hash(needle) {
			return i
		}
	}
	return -1
}

func base26Hash(s string) uint64 {
	var sum uint64 = 0
	for i, c := range s {
		sum += uint64(c-'a') * uint64(math.Pow(float64(26), float64(len(s)-1-i)))
	}
	return sum
}

// 使用 RK 算法的第二个版本。使用两个自定义的 hash 算法计算两次 hash 值，比较两次。
//
// hash 值超出了数据类型范围怎么办？每个字符对应 26 进制的数字符号，然后将所有数位上的数字直接相加，
// 所得的和为 hash 值。这种方法哈希冲突概率应该不小。
//
// 如果发生 hash 冲突，怎么办？首先，如果子串和模式串 hash 值相同，那么如何识别是否是 hash 冲突？
// 貌似除了逐个字符比较，没有办法提前识别 hash 冲突。那么如何避免 hash 冲突时误判呢？
// 考虑使用另一种 hash 算法计算 hash 值，降低冲突的概率。例如，让 26 个小写字母表示前 26 个素数。
//
// 注意，对于 "baa" 和 "aab"，上面提及的这两种 hash 算法的计算结果都是相同的。
func strStr3(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	subSLen := len(needle)
	for i, _ := range haystack {
		subS := haystack[i : i+subSLen]
		if hash1(subS) == hash1(needle) && hash2(subS) == hash2(needle) {
			return i
		}
		if i+subSLen-1 >= len(haystack) {
			return -1
		}
	}
	return -1
}

// 使用 RK 算法的第三个版本。使用 go 内置的 hash 实现
func strStr4(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	needleLen := len(needle)
	hash := sha256.New()
	hash.Write([]byte(needle))
	needleHash := hash.Sum(nil)
	for i, _ := range haystack {
		if i+needleLen-1 >= len(haystack) {
			return -1
		}

		hash.Reset()
		hash.Write([]byte(haystack[i : i+needleLen]))
		if bytes.Equal(needleHash, hash.Sum(nil)) {
			return i
		}
	}
	return -1
}

// "baa" 和 "aab" 不一样，但是哈希值一样
func hash1(s string) int64 {
	var sum int64 = 0
	for _, c := range s {
		sum += int64(c - 'a')
	}
	return sum
}

var first26PrimeNums = []int64{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101,
}

// "baa" 和 "aab" 不一样，但是哈希值一样
func hash2(s string) int64 {
	sum := int64(0)
	for _, c := range s {
		sum += first26PrimeNums[c-'a']
	}
	return sum
}

func TestStrStr(t *testing.T) {
	assert.Equal(t, 0, strStr("sadbutsad", "sad"))
	assert.Equal(t, -1, strStr("bbaa", "aab"))
	assert.Equal(t, 0, strStr2("sadbutsad", "sad"))
	assert.Equal(t, -1, strStr2("bbaa", "aab"))
	assert.Equal(t, 0, strStr2("a", "a"))
}
