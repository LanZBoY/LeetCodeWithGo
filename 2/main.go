package main

import (
	"wentee/leetcode/LinkList"
)

type ListNode = LinkList.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c := 0
	sum := l1.Val + l2.Val + c
	c = sum / 10
	r := sum % 10

	head := &ListNode{Val: r}

	cur := head

	l1 = l1.Next
	l2 = l2.Next

	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val + c
		c = sum / 10
		r = sum % 10
		cur.Next = &ListNode{Val: r}
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum = l1.Val + c
		c = sum / 10
		r = sum % 10
		cur.Next = &ListNode{Val: r}
		cur = cur.Next
		l1 = l1.Next
	}

	for l2 != nil {
		sum = l2.Val + c
		c = sum / 10
		r = sum % 10
		cur.Next = &ListNode{Val: r}
		cur = cur.Next
		l2 = l2.Next
	}

	if c != 0 {
		cur.Next = &ListNode{Val: c}
	}

	return head
}

func main() {
	LinkList.ShowListList(addTwoNumbers(LinkList.NewLinkList(2, 4, 9), LinkList.NewLinkList(5, 6, 4, 9)))
	LinkList.ShowListList(addTwoNumbers(LinkList.NewLinkList(2, 4, 3), LinkList.NewLinkList(5, 6, 4)))
	LinkList.ShowListList(addTwoNumbers(LinkList.NewLinkList(0), LinkList.NewLinkList(0)))
	LinkList.ShowListList(addTwoNumbers(LinkList.NewLinkList(9, 9, 9, 9, 9, 9, 9), LinkList.NewLinkList(9, 9, 9, 9)))
}
