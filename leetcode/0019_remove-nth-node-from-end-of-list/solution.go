package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy

	for i := 0; i < n; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next

	return dummy.Next
}

func main() {
	listNode := buildNode([]int{1, 2, 3, 4, 5})
	printNode(listNode)
	removeNthFromEnd(listNode, 2)
	printNode(listNode)
}

func buildNode(nodes []int) *ListNode {
	list := &ListNode{}
	head := list

	list.Val = nodes[0]
	for i := 1; i < len(nodes); i++ {
		list.Next = &ListNode{}
		list = list.Next
		list.Val = nodes[i]
	}
	return head
}

func printNode(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}
