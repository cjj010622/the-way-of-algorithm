package main

func findFromEnd(head *ListNode, k int) *ListNode {
	p1 := head

	for i := 0; i < k; i++ {
		p1 = p1.Next
	}
	p2 := head

	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p2
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{-1, head}

	// 删除倒数第 n 个，要先找倒数第 n + 1 个节点
	x := findFromEnd(dummy, n+1)

	x.Next = x.Next.Next

	return dummy.Next
}

func middleNode(head *ListNode) *ListNode {
	// 快慢指针初始化指向 head
	slow, fast := head, head
	// 快指针走到末尾时停止
	for fast != nil && fast.Next != nil {
		// 慢指针走一步，快指针走两步
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 慢指针指向中点
	return slow
}
