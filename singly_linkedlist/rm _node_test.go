package singly_linkedlist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 假设将 ListNode.Val 看作是结点的 key，根据 key 做删除
func removeNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	prev, curr := dummy, head
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
			curr.Next = nil
			break
		}
		prev = curr
		curr = curr.Next
	}
	return dummy.Next
}

func TestRemoveNode(t *testing.T) {
	// 操作空链表
	assert.Nil(t, removeNode(nil, 0))
	// 链表中只有一个结点，且是要删除的
	head := &ListNode{
		Val: 1,
	}
	assert.Nil(t, nil, removeNode(head, 1))
	// 删除头节点
	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	assert.Equal(t, &ListNode{Val: 2}, removeNode(head, 1))
	fmt.Printf("head: %+v", head)
	assert.Equal(t, &ListNode{Val: 1, Next: nil}, head)
	// 删除尾结点
	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	assert.Equal(t, &ListNode{Val: 1, Next: nil}, removeNode(head, 2))
}
