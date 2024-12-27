package two_pointer

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

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
