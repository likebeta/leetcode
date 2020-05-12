package helper

import (
    "encoding/json"
    "fmt"
    "strings"
)

type TreeNode struct {
    Left  *TreeNode
    Right *TreeNode
    Val   int
}

// [10,5,-3,3,2,null,11,3,-2,null,1]
func NewTree(s string) *TreeNode {
    var arr []*int
    if err := json.Unmarshal([]byte(s), &arr); err != nil {
        fmt.Println(err)
        return nil
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
