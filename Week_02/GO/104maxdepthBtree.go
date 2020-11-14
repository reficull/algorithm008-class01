package main
import (
    "fmt" 
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
var max = 0
var depth = 0

func main(){
    node1 := TreeNode{3,nil,nil}
    node2 := TreeNode{9,nil,nil}
    node3 := TreeNode{20,nil,nil}
    node4 := TreeNode{15,nil,nil}
    node5 := TreeNode{7,nil,nil}

    node1.Left = &node2 
    node1.Right = &node3
    node3.Left = &node4
    node3.Right = &node5
    ret := maxDepth(&node1)
    fmt.Println("ans:",ret)
}

func maxDepth(root *TreeNode) int {
    if root == nil{
        return 0
    }
    queue := make([]*TreeNode,0)
    queue = append(queue,root)
    level_size := 0
    ans := 0
    for len(queue)>0{
        level_size = len(queue)
        for level_size > 0{
            node := queue[0]
            queue = queue[1:]
            level_size--

            if node.Left != nil{
                queue = append(queue,node.Left)
                fmt.Println("append:",node.Left.Val)
            }
            if node.Right != nil{
                queue = append(queue,node.Right)
                fmt.Println("append:",node.Right.Val)
            }
        }
        ans++
        fmt.Println("level plus:",ans)
    }
    return ans
}

func doTraverse(node *TreeNode) {
    if node != nil{
        depth++
        max = Max(max,depth)
        fmt.Println("max:",max)
        doTraverse(node.Left)
        doTraverse(node.Right)
        depth--
    }
}

func Max(a int,b int) int{
    if a>b{
        return a
    }else{
        return b
    }
}
