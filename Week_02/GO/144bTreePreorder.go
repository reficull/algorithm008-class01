package main
import (
    "fmt" 
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func main(){
    node1 := TreeNode{1,nil,nil}
    node2 := TreeNode{2,nil,nil}
    node3 := TreeNode{3,nil,nil}

    node1.Right = &node2
    node2.Left = &node3

    ret := preorderTraversal(&node1)
    fmt.Println("ret:",ret)
}

func preorderTraversal(root *TreeNode) []int {
    res := make([]int,0)
    doTraverse(root,&res)
    return res

}

func doTraverse(root *TreeNode,res *[]int){
    if root == nil{
        return

    }
    *res = append(*res,root.Val)
    if root.Left != nil{
        doTraverse(root.Left,res)

    }
    if root.Right != nil{
        doTraverse(root.Right,res)

    }


}
