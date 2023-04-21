package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := buildLinkList([]int{1, 2, 3, 4, 5, 6, 7, 8})
	printLinkList(head)
	newlist := swapPairs(head)
	printLinkList(newlist)
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	temp := dummy
	for temp.Next != nil && temp.Next.Next != nil {
		first := temp.Next
		second := temp.Next.Next

		temp.Next = second
		first.Next = second.Next
		second.Next = first

		temp = first
	}

	return dummy.Next
}

func buildLinkList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	p := head
	for i := 1; i < len(arr); i++ {
		p.Next = &ListNode{
			Val: arr[i],
		}
		p = p.Next
	}
	return head
}

func printLinkList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}
