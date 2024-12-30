package singly_linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 哈希表
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	seen := make(map[*ListNode]byte)
	p := head
	for p != nil {
		if _, present := seen[p]; present {
			return p
		} else {
			seen[p] = 0
			p = p.Next
		}
	}
	return nil
}

// 快慢指针
func detectCycle2(head *ListNode) *ListNode {
	if head != nil && head.Next == nil {
		return nil
	}

	// 1 判定是否有环
	fast, slow := head, head
	for fast != nil && slow != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}

	// 2 找到入环点
	ptr := head
	for {
		if ptr == slow {
			return ptr
		}
		ptr = ptr.Next
		slow = slow.Next
	}
	return nil
}

func TestDetectCycle2(t *testing.T) {
	head := &ListNode{
		Val: 1,
	}
	head.Next = head
	assert.Same(t, head, detectCycle2(head))

	head2 := &ListNode{
		Val: 1,
	}
	assert.Nil(t, detectCycle2(head2))
}
