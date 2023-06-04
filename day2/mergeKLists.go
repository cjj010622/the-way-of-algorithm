package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func mergeLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	dummy := &ListNode{-1, nil}
	p := dummy
	//优先级队列，最小堆
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	//将k个链接的头节点加入最小堆
	for _, head := range lists {
		if head != nil {
			heap.Push(&pq, head)
		}
	}

	for pq.Len() > 0 {
		//获取最小节点，接到结果链表中
		node := heap.Pop(&pq).(*ListNode)
		p.Next = node
		if node.Next != node {
			heap.Push(&pq, node.Next)
		}
		//p指针不断前进
		p = p.Next
	}

	return dummy.Next
}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*ListNode)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}
