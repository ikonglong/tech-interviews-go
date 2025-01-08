package singly_linkedlist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 思路，先遍历两个链表，将节点分别放入两个数组中。明确基于数组下标的数位对应关系。
func addTwoNumbersForward(l1 *ListNode, l2 *ListNode) *ListNode {
	a1 := make([]int, 0, 16)
	a2 := make([]int, 0, 16)
	p := l1
	for p != nil {
		a1 = append(a1, p.Val)
		p = p.Next
	}
	p = l2
	for p != nil {
		a2 = append(a2, p.Val)
		p = p.Next
	}

	lenDiff := len(a1) - len(a2)
	if len(a2) > len(a1) {
		lenDiff = len(a2) - len(a1)
		a1, a2 = a2, a1
	}

	fmt.Printf("a1:%v, a2:%v, lenDiff:%v\n", a1, a2, lenDiff)

	carry := 0
	var head *ListNode
	for i := len(a1) - 1; i >= 0; i-- {
		d2 := 0
		if i-lenDiff >= 0 {
			d2 = a2[i-lenDiff]
		}
		r := a1[i] + d2 + carry
		carry = r / 10
		n := &ListNode{Val: r % 10, Next: head}
		head = n
		fmt.Printf("head:%v\n", head.listString())
	}
	// 处理走完链表后 carry 有值的情形
	if carry > 0 {
		n := &ListNode{Val: carry, Next: head}
		head = n
	}
	fmt.Printf("final head:%v\n", func() string {
		if head == nil {
			return "<nil>"
		}
		return head.listString()
	}())
	return head
}

func TestAddTwoNumsForward(t *testing.T) {
	assert.Nil(t, addTwoNumbersForward(nil, nil))
	assert.Equal(t, "[1]", addTwoNumbersForward(&ListNode{Val: 1}, nil).listString())
	assert.Equal(t, "[1,0]", addTwoNumbersForward(&ListNode{Val: 5}, &ListNode{Val: 5}).listString())
	assert.Equal(t, "[9,1,2]", addTwoNumbersForward(newList([]int{6, 1, 7}), newList([]int{2, 9, 5})).listString())
	assert.Equal(t, "[1,0,0,0,1,2]",
		addTwoNumbersForward(newList([]int{9, 9, 6, 1, 7}), newList([]int{3, 9, 5})).listString())
}
