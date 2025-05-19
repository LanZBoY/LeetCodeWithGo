package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

const NULL_NODE = -1001

func NewRandomList(datas [][]int) *Node {
	N := len(datas)

	indexPoint := make([]int, N)
	nodeSlice := make([]*Node, N)

	headData := datas[0]

	head := &Node{Val: headData[0], Next: nil, Random: nil}
	indexPoint[0] = headData[1]
	nodeSlice[0] = head

	var curr *Node = head
	for i := 1; i < N; i++ {
		nodeData := datas[i]
		curr.Next = &Node{Val: nodeData[0], Next: nil, Random: nil}
		indexPoint[i] = nodeData[1]
		nodeSlice[i] = curr.Next
		curr = curr.Next
	}

	for i := range N {
		if indexPoint[i] == NULL_NODE {
			continue
		}
		nodeSlice[i].Random = nodeSlice[indexPoint[i]]
	}

	return head
}

func showRandomList(head *Node) {

	cur := head
	fmt.Printf("[")
	for cur != nil {
		fmt.Printf("{")
		fmt.Printf("%v", cur.Val)

		if cur.Random != nil {
			fmt.Printf("-> %v", cur.Random.Val)
		}
		fmt.Printf("}")

		cur = cur.Next
	}
	fmt.Printf("]\n")
}

func showRandomListP(head *Node) {

	cur := head
	fmt.Printf("[")
	for cur != nil {
		fmt.Printf("{")
		fmt.Printf("%p", cur)

		if cur.Random != nil {
			fmt.Printf("-> %p", cur.Random)
		}
		fmt.Printf("}")

		cur = cur.Next
	}
	fmt.Printf("]\n")
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	cur := head

	for cur != nil {
		tmp := cur.Next
		newNode := &Node{Val: cur.Val, Next: cur.Next}
		cur.Next = newNode
		cur = tmp
	}

	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	cur = head
	newHead := head.Next
	for cur != nil {
		copy := cur.Next
		cur.Next = copy.Next
		if copy.Next != nil {
			copy.Next = copy.Next.Next
		}
		cur = cur.Next
	}
	return newHead
}

func main() {
	showRandomListP(copyRandomList(NewRandomList([][]int{{7, NULL_NODE}, {13, 0}, {11, 4}, {10, 2}, {1, 0}})))
	showRandomListP(copyRandomList(NewRandomList([][]int{{1, 1}, {2, 1}})))
	showRandomListP(copyRandomList(NewRandomList([][]int{{3, NULL_NODE}, {3, 0}, {3, NULL_NODE}})))
}
