package main

import (
	"fmt"
	"leetcode/helper"
	"strings"
)

type Node = helper.DoubleListNode

// LRUCache LRU 缓存
type LRUCache struct {
	capacity int
	mapping  map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	var head, tail Node
	head.Next = &tail
	tail.Prev = &head

	lruCache := LRUCache{}
	lruCache.capacity = capacity
	lruCache.mapping = make(map[int]*Node, capacity)
	lruCache.head = &head
	lruCache.tail = &tail
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.mapping[key]; ok {
		this.touch(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.mapping[key]; ok {
		node.Val = value
		this.touch(node)
		return
	}

	if len(this.mapping) == this.capacity {
		delete(this.mapping, this.tail.Prev.Key)
		this.removeNode(this.tail.Prev)
	}
	node := &Node{Key: key, Val: value}
	this.insertFirst(node)
	this.mapping[key] = node
}

func (this *LRUCache) touch(node *Node) {
	this.removeNode(node)
	this.insertFirst(node)
}

func (this *LRUCache) insertFirst(node *Node) {
	node.Prev, node.Next = this.head, this.head.Next
	this.head.Next, this.head.Next.Prev = node, node
}

func (this *LRUCache) removeNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) dump() {
	{
		var lines []string
		for node := this.head; node != nil; node = node.Next {
			lines = append(lines, fmt.Sprintf("[%d,%d]", node.Key, node.Val))
		}
		helper.Log(strings.Join(lines, "=>"))
	}

	{
		var lines []string
		for node := this.tail; node != nil; node = node.Prev {
			lines = append(lines, fmt.Sprintf("[%d,%d]", node.Key, node.Val))
		}
		helper.Log(strings.Join(lines, "<="))
	}
}

func main() {
	cache := Constructor(2) // 缓存容量
	cache.Put(1, 1)
	cache.Put(2, 2)
	helper.Assert(cache.Get(1) == 1)  // 返回 1
	cache.Put(3, 3)                   // 该操作会使得 2 作废
	helper.Assert(cache.Get(2) == -1) // 返回 -1 (未找到)
	cache.Put(4, 4)                   // 该操作会使得 1 作废
	helper.Assert(cache.Get(1) == -1) // 返回 -1 (未找到)
	helper.Assert(cache.Get(3) == 3)  // 返回 3
	helper.Assert(cache.Get(4) == 4)  // 返回 4
}
