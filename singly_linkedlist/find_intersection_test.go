package singly_linkedlist

import "fmt"

//  1. 判定两个链表是否相交
//     a. 遍历两个链表，并计算其长度，记录最后一个结点
//     b. 判断最后一个结点的引用是否相同。相同则相交
//  2. 若相交，则找到相交结点
//     a. 再次遍历两个链表。让较长链表的遍历指针先多走 |lenA-lenB| 步
//     b. 让两个遍历指针同步走，走到结点引用相同时为止。此结点就是相交结点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	ptrA, ptrB := headA, headB
	lenA, lenB := 0, 0
	for {
		lenA++
		if ptrA.Next == nil {
			break
		} else {
			ptrA = ptrA.Next
		}
	}
	for {
		lenB++
		if ptrB.Next == nil {
			break
		} else {
			ptrB = ptrB.Next
		}
	}
	if ptrA != ptrB {
		fmt.Printf("headA, headB doesn't intersect, ptrA:%v, ptrB:%v", ptrA, ptrB)
		return nil
	}
	fmt.Printf("headA, headB intersect, ptrA:%v, ptrB:%v", ptrA, ptrB)

	// 2 找到相交结点

	ptrA, ptrB = headA, headB
	ptr := ptrA
	moreCount := lenA - lenB
	if lenB > lenA {
		ptr = ptrB
		moreCount = lenB - lenA
	}
	for i := 1; i <= moreCount; i++ {
		ptr = ptr.Next
	}
	if lenB > lenA {
		ptrB = ptr
	} else {
		ptrA = ptr
	}

	for ptrA != nil && ptrB != nil {
		if ptrA == ptrB {
			return ptrA
		}
		ptrA = ptrA.Next
		ptrB = ptrB.Next
	}

	return nil
}
