package hash

// 解法：
// 1. 将数组的所有元素放入哈希表中，这样可以在接近常数的时间内完成查找操作，同时重复元素
// 2. 遍历哈希表，寻找序列的开头。
// 	如果当前整数的前一个整数不在表中，那么当前整数是序列的开头。查看下一个数是否在表中，
// 		若在，序列长度加 1，直到下一个数不在。跟最大长度比较，若更大就替换。
// 	如果在，就不是开头，跳过。

func longestConsecutive(nums []int) int {
	m := make(map[int]byte, len(nums))
	for _, n := range nums {
		m[n] = 0
	}

	maxLen := 0
	for n, _ := range m {
		if _, present := m[n-1]; present {
			continue
		}

		count := 1
		x := n + 1
		for {
			if _, present := m[x]; present {
				count++
				x++
			} else {
				break
			}
		}

		if count > maxLen {
			maxLen = count
		}
	}
	return maxLen
}
