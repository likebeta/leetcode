package main

import (
	"leetcode/helper"
)

type Node = helper.LinkedListNode
type LinkedList = helper.LinkedList

// LFUCache LFU 缓存
type LFUCache struct {
	capacity    int
	minFreq     int
	keyToNode   map[int]*Node
	keyToFreq   map[int]int
	freqToNodes map[int]*LinkedList
}

func Constructor(capacity int) LFUCache {
	lfuCache := LFUCache{}
	lfuCache.capacity = capacity
	lfuCache.keyToNode = make(map[int]*Node, capacity)
	lfuCache.keyToFreq = make(map[int]int, capacity)
	lfuCache.freqToNodes = make(map[int]*LinkedList)
	return lfuCache
}

func (this *LFUCache) Get(key int) int {
	if node, ok := this.keyToNode[key]; ok {
		this.touch(node)
		return node.Val
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if node, ok := this.keyToNode[key]; ok {
		node.Val = value
		this.touch(node)
		return
	}

	if len(this.keyToNode) == this.capacity {
		nodes := this.freqToNodes[this.minFreq]
		node := nodes.RemoveLast()
		if nodes.Empty() {
			delete(this.freqToNodes, this.minFreq)
		}
		delete(this.keyToNode, node.Key)
		delete(this.keyToFreq, node.Key)
	}

	node := &Node{Key: key, Val: value}
	this.keyToNode[key] = node
	this.keyToFreq[key] = 1
	nodes, ok := this.freqToNodes[1]
	if !ok {
		nodes = helper.NewLinkedList()
		this.freqToNodes[1] = nodes
	}
	nodes.InsertFirst(node)
	this.minFreq = 1
}

func (this *LFUCache) touch(node *Node) {
	freq := this.keyToFreq[node.Key]
	delete(this.keyToFreq, node.Key)

	nodes := this.freqToNodes[freq]
	nodes.Remove(node)
	if nodes.Empty() {
		delete(this.freqToNodes, freq)
		if this.minFreq == freq {
			this.minFreq++
		}
	}

	freq++
	nodes, ok := this.freqToNodes[freq]
	if !ok {
		nodes = helper.NewLinkedList()
		this.freqToNodes[freq] = nodes
	}
	nodes.InsertFirst(node)
	this.keyToFreq[node.Key] = freq
}

func main() {
	cache := Constructor(2) // 缓存容量
	cache.Put(1, 1)
	cache.Put(2, 2)
	helper.Assert(cache.Get(1) == 1)  // 返回 1
	cache.Put(3, 3)                   // 该操作会使得 2 作废
	helper.Assert(cache.Get(2) == -1) // 返回 -1 (未找到)
	helper.Assert(cache.Get(3) == 3)  // 返回 3

	cache.Put(4, 4)                   // 该操作会使得 1 作废
	helper.Assert(cache.Get(1) == -1) // 返回 -1 (未找到)
	helper.Assert(cache.Get(3) == 3)  // 返回 3
	helper.Assert(cache.Get(4) == 4)  // 返回 4
}
