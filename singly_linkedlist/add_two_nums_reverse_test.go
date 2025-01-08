package singly_linkedlist

// https://leetcode.cn/problems/sum-lists-lcci/

// 递归思路
//
// 递推公式：求用链表表示的两个数的和，和也用链表表示 = 从个位开始，将对应数位上的数字相加，
// 对 10 取模的结果作为结果链表的节点，整除 10 的结果作为下一个数位的进位数，参与对应数位
// 上数字的求和运算。
// 终止条件：两个链表都遍历完成
// 边缘情况：一个链表长，一个链表短。
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	carry := 0 // 进位数
	dummy := &ListNode{}
	tail := dummy
	for p1 != nil || p2 != nil {
		d1, d2 := 0, 0
		if p1 != nil {
			d1 = p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			d2 = p2.Val
			p2 = p2.Next
		}
		r := d1 + d2 + carry
		carry = r / 10
		tail.Next = &ListNode{Val: r % 10}
		tail = tail.Next
	}
	// 处理走完链表后，carry 有值的情形
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}

// 思路同上
// 终止条件：任何一个链表走完
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	carry := 0
	dummy := &ListNode{}
	tail := dummy
	for p1 != nil && p2 != nil {
		r := p1.Val + p2.Val + carry
		carry = r / 10
		tail.Next = &ListNode{Val: r % 10}
		tail = tail.Next

		p1 = p1.Next
		p2 = p2.Next
	}
	// 处理其中一个链表还未走完的情形
	if p2 != nil {
		p1 = p2
	}
	for p1 != nil {
		r := p1.Val + carry
		carry = r / 10
		tail.Next = &ListNode{Val: r % 10}
		tail = tail.Next
		p1 = p1.Next
	}
	// 处理走完链表后 carry 大于 0 的情形
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
