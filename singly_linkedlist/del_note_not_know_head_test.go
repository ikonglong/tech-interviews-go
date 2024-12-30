package singly_linkedlist

// https://leetcode.cn/problems/delete-middle-node-lcci/

func deleteNode(node *ListNode) {
	if node == nil || node.Next == nil { // 如果是尾结点，可能需要将其标记为假结点或已删除
		return
	}
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
