package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 填充每个节点的下一个右侧节点指针
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
		for node := leftmost; node != nil; node = node.Next {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}
	return root
}

func main() {

}
