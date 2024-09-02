package main

import (
	"leetcode/helper"
)

type Node = helper.LinkedListNode
type LinkedList = helper.LinkedList

// LRUCache LRU 缓存
type LRUCache struct {
	capacity   int
	mapping    map[int]*Node
	linkedList *LinkedList
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{}
	lruCache.capacity = capacity
	lruCache.mapping = make(map[int]*Node, capacity)
	lruCache.linkedList = helper.NewLinkedList()
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
		node := this.linkedList.RemoveLast()
		delete(this.mapping, node.Key)
	}
	node := &Node{Key: key, Val: value}
	this.linkedList.InsertFirst(node)
	this.mapping[key] = node
}

func (this *LRUCache) touch(node *Node) {
	this.linkedList.Remove(node)
	this.linkedList.InsertFirst(node)
}

func (this *LRUCache) dump() {
	lToR, rToL := this.linkedList.Dump()
	helper.Log(lToR)
	helper.Log(rToL)
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
