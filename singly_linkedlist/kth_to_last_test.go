package singly_linkedlist

// 递归解法
func kthToLast(head *ListNode, k int) int {
	_, n := findKthToLast(head, k)
	if n != nil {
		return n.Val
	}
	return 0
}

func findKthToLast(head *ListNode, k int) (count int, node *ListNode) {
	if head == nil {
		return 0, nil
	}
	count, node = findKthToLast(head.Next, k)
	count++ // 即时已经找到了也可以加一，因为返回计数器值的目的就是为了找到目标结点，没有其他用途
	if count < k {
		return count, nil
	} else if count == k {
		return count, head
	} else {
		return count, node
	}
}

func kthToLast2(head *ListNode, k int) int {
	p1, p2 := head, head
	count := 1
	for ; count <= k-1 && p1 != nil; count++ { // count <= k 就多走了一步
		p1 = p1.Next
	}
	for p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2.Val
}
