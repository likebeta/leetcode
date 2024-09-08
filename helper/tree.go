package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func (root *TreeNode) Clone() *TreeNode {
	return CloneTree(root)
}

func (root *TreeNode) Print(width int) {
	PrintTree(root, width)
}

func (root *TreeNode) ToArray() []*int {
	if root == nil {
		return nil
	}
	nodes, last := []*TreeNode{root}, 0
	for i := 0; i < len(nodes); i++ {
		if nodes[i] != nil {
			last = i
			nodes = append(nodes, nodes[i].Left, nodes[i].Right)
		}
	}
	arr := make([]*int, last+1)
	for i := range arr {
		if nodes[i] != nil {
			arr[i] = &nodes[i].Val
		}
	}
	return arr
}

func (root *TreeNode) Dump() string {
	if root == nil {
		return "[]"
	}
	arr := root.ToArray()
	data, err := json.Marshal(arr)
	if err != nil {
		log.Panic("dump failed:", err)
	}
	return string(data)
}

func (root *TreeNode) DumpNode() string {
	if root == nil {
		return "null"
	} else {
		return strconv.Itoa(root.Val)
	}
}

func ParseTree(bfs string) *TreeNode {
	// bfs: [10,5,-3,3,2,null,11,3,-2,null,1]
	var arr []*int
	if err := json.Unmarshal([]byte(bfs), &arr); err != nil {
		log.Panic("parse failed:", err)
	}

	var root *TreeNode
	cnt := len(arr)
	if cnt > 0 {
		root = &TreeNode{Val: *arr[0]}
		arr = arr[1:]
		list := []*TreeNode{root}
		for len(list) > 0 && len(arr) > 0 {
			curr := list[0]
			list = list[1:]
			if len(arr) > 0 {
				if arr[0] != nil {
					curr.Left = &TreeNode{Val: *arr[0]}
					list = append(list, curr.Left)
				}
				arr = arr[1:]
			}

			if len(arr) > 0 {
				if arr[0] != nil {
					curr.Right = &TreeNode{Val: *arr[0]}
					list = append(list, curr.Right)
				}
				arr = arr[1:]
			}
		}
	}
	return root
}

var NewTree = ParseTree

func PrintTree(root *TreeNode, width int) {
	fmt.Println("-------------------Tree---------------------")
	lines := formatInOrder(nil, root, 0, "H", width)
	fmt.Println(strings.Join(lines, "\n"))
	fmt.Println("--------------------End---------------------")
}

func formatInOrder(lines []string, node *TreeNode, height int, to string, length int) []string {
	if node != nil {
		lines = formatInOrder(lines, node.Right, height+1, "v", length)
		val := fmt.Sprintf("%s%v%s", to, node.Val, to)
		lenM := len(val)
		lenL := (length - lenM) / 2
		lenR := length - lenM - lenL
		val = strings.Repeat(" ", lenL) + val + strings.Repeat(" ", lenR)
		lines = append(lines, strings.Repeat(" ", height*length)+val)
		lines = formatInOrder(lines, node.Left, height+1, "^", length)
	}
	return lines
}

func CloneTree(head *TreeNode) *TreeNode {
	if head == nil {
		return head
	}
	node := &TreeNode{Val: head.Val}
	node.Left = CloneTree(head.Left)
	node.Right = CloneTree(head.Right)
	return node
}

func InOrder(node *TreeNode) []int {
	arr := make([]int, 0)
	var inOrder func(*TreeNode)
	inOrder = func(node *TreeNode) {
		if node != nil {
			inOrder(node.Left)
			arr = append(arr, node.Val)
			inOrder(node.Right)
		}
	}

	inOrder(node)
	return arr
}
