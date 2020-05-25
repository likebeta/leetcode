package main

import (
	"leetcode/helper"
)

// LRU缓存机制
type LRUCache struct {
	capacity int
	kvs      map[int]*Node
	head     *Node
	tail     *Node
}

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{}
	lru.capacity = capacity
	lru.kvs = make(map[int]*Node, capacity)
	return lru
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.kvs[key]; ok {
		this.touch(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) touch(node *Node) {
	if node.Prev != nil {
		if node == this.tail {
			this.tail = node.Prev
		}
		node.Prev.Next = node.Next
		if node.Next != nil {
			node.Next.Prev = node.Prev
		}
		node.Prev = nil
		node.Next = this.head
		this.head.Prev = node
		this.head = node
	}
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.kvs[key]; ok {
		v.Val = value
		this.touch(v)
	} else {
		if len(this.kvs) == this.capacity {
			delete(this.kvs, this.tail.Key)
			this.kvs[key] = this.tail
			this.tail.Key = key
			this.tail.Val = value
			this.touch(this.tail)
		} else {
			node := &Node{Key: key, Val: value}
			this.kvs[key] = node
			if this.head != nil {
				node.Next = this.head
				this.head.Prev = node
			}
			this.head = node
			if this.tail == nil {
				this.tail = node
			}
		}
	}
}

func main() {
	cache := Constructor(2) // 缓存容量
	cache.Put(1, 1)
	cache.Put(2, 2)
	helper.Assert(cache.Get(1) == 1)  // 返回  1
	cache.Put(3, 3)                   // 该操作会使得密钥 2 作废
	helper.Assert(cache.Get(2) == -1) // 返回 -1 (未找到)
	cache.Put(4, 4)                   // 该操作会使得密钥 1 作废
	helper.Assert(cache.Get(1) == -1) // 返回 -1 (未找到)
	helper.Assert(cache.Get(3) == 3)  // 返回  3
	helper.Assert(cache.Get(4) == 4)  // 返回  4
}
