package main

import "fmt"

func main() {
	var l1 *ListNode
	l2 := &ListNode{
		Val: 0,
	}
	fmt.Println(mergeTwoLists(l1, l2))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	temp := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			temp.Next = list1
			list1 = list1.Next
		} else {
			temp.Next = list2
			list2 = list2.Next
		}
		temp = temp.Next
	}

	if list1 != nil {
		temp.Next = list1
	} else {
		temp.Next = list2
	}

	return head.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Print(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val, " ")
		l = l.Next
	}
	fmt.Println()
}

func MakeListNode(x []int) *ListNode {
	list := &ListNode{}
	head := list

	list.Val = x[0]
	for i:=1; i < len(x); ++ {
		list.Next = &ListNode{}
		list = list.Next
		list.Val = x[i]
	}
	return head
}