package singly_linkedlist

// https://leetcode.cn/problems/remove-duplicate-node-lcci/?envType=problem-list-v2&envId=two-pointers

// 使用散列表保存已经看见的结点的键
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	seen := make(map[int]int, 64)
	dummy := &ListNode{
		Val:  1,
		Next: head,
	}
	prev, curr := dummy, head
	for curr != nil {
		if _, present := seen[curr.Val]; present {
			prev.Next = curr.Next
			curr.Next = nil
			// 让 curr 指向下一个结点，但 prev 指针不能移动，因为下一个结点可能也应该被删除
			curr = prev.Next
		} else {
			seen[curr.Val] = 0
			// 正常移动指针
			prev = curr
			curr = curr.Next
		}
	}
	return dummy.Next
}

// 不使用缓冲区
func removeDuplicateNodes2(head *ListNode) *ListNode {
	curr := head
	// 让每一个结点跟它之后的结点比较一遍，删除后出现的重复结点
	for curr != nil {
		// 注意，runner 不能指向 curr.next，因为如果下一个结点应该被删除，那么
		// 需要 prev 指针。这里，runner 扮演的就是 prev 角色
		runner := curr
		for runner.Next != nil { // 循环的目的是将 curr 跟后续每个结点比较，因此这里检查的是 runner.Next
			if runner.Next.Val == curr.Val {
				// 删除 runner.Next
				runner.Next = runner.Next.Next
				// 注意，这里不能移动 runner 指针，因为 runner.Next.Next 可能也应该被删除
			} else {
				// todo 移动指针
				runner = runner.Next
			}
		}

		curr = curr.Next
	}
	return head
}
