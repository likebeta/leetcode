package main

import (
	"encoding/json"
	"leetcode/helper"
	"log"
	"math/rand"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	cache := make(map[int]*Node)
	var deepClone func(*Node) *Node
	deepClone = func(node *Node) *Node {
		if node == nil {
			return node
		}
		if node, ok := cache[node.Val]; ok {
			return node
		}
		cloneNode := &Node{Val: node.Val}
		cache[node.Val] = cloneNode
		for i := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, deepClone(node.Neighbors[i]))
		}
		return cloneNode
	}
	return deepClone(node)
}

func makeGraph(in string) *Node {
	var adjList [][]int
	if err := json.Unmarshal([]byte(in), &adjList); err != nil {
		log.Panic("parse json error", err)
	}
	var node *Node
	if N := len(adjList); N > 0 {
		arr := make([]*Node, N)
		for i := range adjList {
			arr[i] = &Node{Val: i + 1}
		}

		for i := range adjList {
			for j := range adjList[i] {
				arr[i].Neighbors = append(arr[i].Neighbors, arr[adjList[i][j]-1])
			}
		}
		node = arr[rand.Intn(N)]
	}
	return node
}

func dumpGraph(node *Node) string {
	adjList := make([][]int, 0)

	if node != nil {
		cache := make(map[int][]int)
		queue := []*Node{node}
		for len(queue) > 0 {
			node, queue = queue[0], queue[1:]
			if _, ok := cache[node.Val]; ok {
				continue
			}
			cache[node.Val] = make([]int, len(node.Neighbors))
			for i := range node.Neighbors {
				queue = append(queue, node.Neighbors[i])
				cache[node.Val][i] = node.Neighbors[i].Val
			}
		}
		adjList = make([][]int, len(cache))
		for k, v := range cache {
			adjList[k-1] = v
		}
	}

	data, err := json.Marshal(&adjList)
	if err != nil {
		log.Fatal("dump json error", err)
	}
	return string(data)
}

func testOne(in string) {
	node := makeGraph(in)
	helper.Assert(in == dumpGraph(cloneGraph(node)))
}

func main() {
	testOne("[[2,4],[1,3],[2,4],[1,3]]")
	testOne("[[]]")
	testOne("[]")
	testOne("[[2],[1]]")
}
