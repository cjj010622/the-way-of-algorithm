package main

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

func detectCycle(head *ListNode) *ListNode {
	var fast, slow *ListNode
	fast, slow = head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}

func getIntersectionNode(headA *ListNode, headB *ListNode) *ListNode {
	// p1 指向 A 链表头结点，p2 指向 B 链表头结点
	p1, p2 := headA, headB
	for p1 != p2 {
		// p1 走一步，如果走到 A 链表末尾，转到 B 链表
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		// p2 走一步，如果走到 B 链表末尾，转到 A 链表
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}
