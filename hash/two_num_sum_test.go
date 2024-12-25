package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))
	for i, n := range nums {
		numMap[n] = i
	}
	for i, n := range nums {
		if j, present := numMap[target-n]; present && j > i {
			return []int{i, j}
		}
		numMap[n] = i
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, present := numMap[target-n]; present {
			return []int{j, i}
		}
		numMap[n] = i
	}
	return []int{}
}

func TestTwoSum(t *testing.T) {
	assert.Equal(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{1, 2}, twoSum([]int{3, 2, 4}, 6))
	assert.Equal(t, []int{0, 1}, twoSum([]int{3, 3}, 6))

	assert.Equal(t, []int{0, 1}, twoSum2([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{1, 2}, twoSum2([]int{3, 2, 4}, 6))
	assert.Equal(t, []int{0, 1}, twoSum2([]int{3, 3}, 6))
}
