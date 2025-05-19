package main

import (
	"wentee/leetcode/LinkList"
)

type ListNode = LinkList.ListNode

func reverseList(head *ListNode) *ListNode {

	var prev *ListNode = nil

	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}

	return prev
}

func main() {
	showList := LinkList.ShowListList
	newList := LinkList.NewLinkList

	data := newList(1, 2, 3, 4, 5)
	showList(data)
	showList(reverseList(data))

	data = newList(1, 2)
	showList(data)
	showList(reverseList(data))

	data = newList()
	showList(data)
	showList(reverseList(data))
}
