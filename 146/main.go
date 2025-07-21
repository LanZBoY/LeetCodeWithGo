package main

import "fmt"

type Node struct {
	key  int
	data int
	prev *Node
	next *Node
}

func NewNode(key int, data int) *Node {
	return &Node{
		key:  key,
		data: data,
		prev: nil,
		next: nil,
	}
}

type LRUCache struct {
	size    int
	cap     int
	hashMap map[int]*Node
	head    *Node
	tail    *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{size: 0, cap: capacity, hashMap: make(map[int]*Node), head: nil, tail: nil}
}

func (cache *LRUCache) Get(key int) int {

	if node, ok := cache.hashMap[key]; ok {
		if node == cache.head {
			return node.data
		}
		prevNode := node.prev
		nextNode := node.next

		if prevNode != nil {
			prevNode.next = nextNode
		}

		if nextNode != nil {
			nextNode.prev = prevNode
		}

		node.prev = nil
		node.next = cache.head

		cache.head.prev = node

		cache.head = node

		if cache.tail == node {
			cache.tail = prevNode
		}
		return node.data
	}

	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	if cache.size == 0 {
		newNode := NewNode(key, value)
		cache.head = newNode
		cache.tail = newNode
		cache.hashMap[key] = newNode
		cache.size++
		return
	}

	if node, ok := cache.hashMap[key]; ok {
		if node == cache.head {
			node.data = value
			return
		}
		prevNode := node.prev
		nextNode := node.next

		if prevNode != nil {
			prevNode.next = nextNode
		}

		if nextNode != nil {
			nextNode.prev = prevNode
		}

		node.prev = nil
		node.next = cache.head

		cache.head.prev = node

		cache.head = node

		if cache.tail == node {
			cache.tail = prevNode
		}

		node.data = value
		return
	}

	newNode := NewNode(key, value)
	cache.hashMap[key] = newNode
	cache.head.prev = newNode
	newNode.next = cache.head
	cache.head = newNode
	cache.size++

	if cache.size > cache.cap {
		delete(cache.hashMap, cache.tail.key)
		prevNode := cache.tail.prev
		prevNode.next = nil
		cache.tail.prev = nil
		cache.tail = prevNode
		cache.size--
	}

}

func main() {
	lru := Constructor(2)

	lru.Put(2, 1)
	lru.Put(2, 2)
	fmt.Printf("%v\n", lru.Get(2))
	lru.Put(1, 1)
	lru.Put(4, 1)
	fmt.Printf("%v\n", lru.Get(2))

}
