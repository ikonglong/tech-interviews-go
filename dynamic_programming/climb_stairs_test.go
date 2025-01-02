package dynamic_programming

import "testing"

func climbStairs(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// 有记忆的，避免重复计算
func climbStairs2(n int) int {
	return doClimb(n, make([]int, n+1))
}

func doClimb(n int, mem []int) int {
	if n < 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if mem[n] <= 0 {
		mem[n] = doClimb(n-1, mem) + doClimb(n-2, mem)
	}
	return mem[n]
}

func TestClimbStairs(t *testing.T) {
	climbStairs(45)
}
