package two_pointer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.cn/problems/sorted-merge-lcci/?envType=problem-list-v2&envId=two-pointers

// 正向双指针
func merge(A []int, m int, B []int, n int) {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for p1 < m && p2 < n {
		if A[p1] <= B[p2] {
			sorted = append(sorted, A[p1])
			p1++
		} else {
			sorted = append(sorted, B[p2])
			p2++
		}
	}
	if p1 < m {
		for ; p1 < len(A); p1++ {
			sorted = append(sorted, A[p1])
		}
	}
	if p2 < n {
		for ; p2 < len(B); p2++ {
			sorted = append(sorted, B[p2])
		}
	}
	// 虽然 A 是值对象，看起来不可修改，但这只是对于标准库的使用者来说。
	// 对于标准库来说，下面这个内置函数就修改了 A 的底层数组。
	copy(A, sorted)
}

// 逆向双指针
func merge2(A []int, m int, B []int, n int) {
	p := m + n - 1
	pa := m - 1
	pb := n - 1
	for pa >= 0 && pb >= 0 {
		if A[pa] > B[pb] {
			A[p] = A[pa]
			pa--
		} else {
			A[p] = B[pb]
			pb--
		}
		p--
	}
	if pb <= -1 {
		return
	}
	if pa <= -1 {
		for p >= 0 && pb >= 0 {
			A[p] = B[pb]
			p--
			pb--
		}
	}
}

func TestMergeWithTwoPtr(t *testing.T) {
	doTestMerge(t, merge)
	doTestMerge(t, merge2)
}

func doTestMerge(t *testing.T, merge func([]int, int, []int, int)) {
	var A1 = make([]int, 0)
	merge(A1, 0, []int{}, 0)
	assert.Equal(t, []int{}, A1)

	A2 := []int{1, 2, 3, 0, 0, 0}
	merge(A2, 3, []int{2, 5, 6}, 3)
	fmt.Printf("%v\n", A2)
	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, A2)

	A3 := []int{1, 2, 3, 0, 0}
	merge(A3, 3, []int{5, 6}, 2)
	assert.Equal(t, []int{1, 2, 3, 5, 6}, A3)

	A4 := []int{4, 5, 6, 0, 0}
	merge(A4, 3, []int{1, 2}, 2)
	assert.Equal(t, []int{1, 2, 4, 5, 6}, A4)
}
