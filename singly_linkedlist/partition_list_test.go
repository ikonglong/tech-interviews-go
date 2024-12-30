package singly_linkedlist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 创建两个链表，小于 x 的一个，大于 x 的一个。
// 遍历原链表，拆分成两个链表
// 合并
func partition(head *ListNode, x int) *ListNode {
	var beforeStart, beforeEnd, afterStart, afterEnd *ListNode
	curr := head
	for curr != nil {
		// 下一步会将当前结点从原链表分离，因此先保存 next
		next := curr.Next
		// 将当前结点从原链表分离
		curr.Next = nil

		if curr.Val < x {
			if beforeStart == nil {
				beforeStart = curr
				beforeEnd = curr
			} else {
				beforeEnd.Next = curr
				beforeEnd = curr
			}
		} else {
			if afterStart == nil {
				afterStart = curr
				afterEnd = curr
			} else {
				afterEnd.Next = curr
				afterEnd = curr
			}
		}

		curr = next
	}
	if beforeStart == nil {
		return afterStart
	}
	beforeEnd.Next = afterStart
	return beforeStart
}

func partition2(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	// 创建一个新链表
	// 为什么需要初始化为指向原链表的 head？请看使用它们的地方
	newHead := head
	newTail := head

	node := head
	for node != nil {
		// 加上这行代码会发生死循环，原因看下面的注释。
		// fmt.Printf("head:%v, tail:%v, new list: %s\n", newHead.nodeString(), newTail.nodeString(), head.listString())

		// 下一步会将当前结点从原链表分离，因此先保存 next
		next := node.Next
		// 是否需要明确分离？不需要，因为在头部插入时，node.Next 会被设置。在尾部插入时，node.Next 在下一次执行尾部插入操作时会被设置。
		// node.Next = nil // version1

		if node.Val < x {
			// 在新链表头部插入
			node.Next = newHead
			newHead = node // 指向新的头节点
			// 处理第一个结点时，执行完上面的代码后，newHead 指向了 node，node.Next 指向了它自身，形成了单个结点的环。
			// 如果打印 newHead，就会死循环。
			//
			// 如果输入的是 ([1], 0)，那么必须在循环结束后，将这个环在末尾处断开。这就是为什么循环后有代码行：`newTail.Next = nil`
		} else {
			// 在新链表尾部插入
			newTail.Next = node
			newTail = node
		}

		node = next
	}
	newTail.Next = nil // version2。为什么需要这行，看上面的注释

	return newHead
}

func TestPartition(t *testing.T) {
	// assert.Equal(t, "[2,2,1,5,3,4]", partition2(newList([]int{1, 4, 3, 2, 5, 2}), 3).listString())
	l := partition2(newList([]int{1, 4, 3, 2, 5, 2}), 3).listString()
	fmt.Println(l)
	assert.Regexp(t, `^\[\d+(,\d+)*\]$`, l)
	// 对于这个 case，version1 会死循环，version2 ok
	assert.Equal(t, "[1]", partition2(newList([]int{1}), 0).listString())
	assert.Equal(t, "[]", partition2(newList(nil), 0).listString())
}
