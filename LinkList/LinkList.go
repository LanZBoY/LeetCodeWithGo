package LinkList

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Construct LinkList via slice or array
func NewLinkList(datas ...int) *ListNode {

	if len(datas) == 0 {
		return nil
	}

	head := &ListNode{Val: datas[0], Next: nil}

	current := head
	for i := 1; i < len(datas); i++ {

		current.Next = &ListNode{Val: datas[i], Next: nil}
		current = current.Next
	}

	return head
}

func ShowListList(head *ListNode) {

	fmt.Printf("[ ")

	current := head

	for current != nil {

		if current.Next != nil {
			fmt.Printf("%v, ", current.Val)
		} else {
			fmt.Printf("%v", current.Val)
		}

		current = current.Next
	}

	fmt.Printf(" ]\n")
}
