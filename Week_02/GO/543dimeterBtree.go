package main
import (
    "fmt" 
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var res = 0

func main(){

    node1 := TreeNode{1,nil,nil}
    node2 := TreeNode{2,nil,nil}
    node3 := TreeNode{3,nil,nil}
    node4 := TreeNode{4,nil,nil}
    node5 := TreeNode{5,nil,nil}

    node1.Left = &node2
    node1.Right = &node3
    node2.Left = &node4
    node2.Right = &node5

    ans := diameterOfBinaryTree(&node1)
    fmt.Println("ret:",ans)
}

func diameterOfBinaryTree(root *TreeNode) int {
    var max int = 0
    var dfs = func (root *TreeNode) int {
        if root == nil {
            return 0

        }
        if root.Left == nil && root.Right == nil {
            return 1

        }
        left := dfs(root.Left)
        right := dfs(root.Right)
        //更新路径长度
        max = Max(max, left + right)
        return Max(left, right) + 1

    }
    dfs(root)
    return max
}

func dfs(root *TreeNode) int {
    if root == nil{
        return 0
    }
    if root.Left == nil && root.Right == nil {
        return 1
    }
    l:= dfs(root.Left) 
    r:= dfs(root.Right) 
    path := Max(l,r)
    res = Max(l+r,res)
    path++
    fmt.Printf("node :%d,l=%d,r=%d,res:%d,path:%d\n",root.Val,l,r,res,path)
    return path
}


func Max(a int,b int) int{
    if a>b{
        return a
    }else{
        return b
    }
}
