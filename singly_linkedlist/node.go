package singly_linkedlist

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) nodeString() string {
	nextStr := "nil"
	if n.Next != nil {
		nextStr = fmt.Sprintf("%d", n.Next.Val)
	}
	return fmt.Sprintf("{Val:%d,Next:%s}", n.Val, nextStr)
}

func (n *ListNode) listString() string {
	if n == nil {
		return "[]"
	}

	sb := strings.Builder{}
	sb.WriteRune('[')
	curr := n
	for curr != nil {
		sb.WriteString(strconv.Itoa(curr.Val))
		if curr.Next != nil {
			sb.WriteRune(',')
		}
		curr = curr.Next
	}
	sb.WriteRune(']')
	return sb.String()
}

func newList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	var head, tail *ListNode
	for i, v := range vals {
		n := &ListNode{
			Val: v,
		}
		if i == 0 {
			head = n
			tail = n
		} else {
			tail.Next = n
			tail = n
		}
	}
	return head
}
