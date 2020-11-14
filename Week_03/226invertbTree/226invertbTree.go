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
    /*
    node11 := TreeNode{4,nil}
    node12 := TreeNode{2,nil}
    node13 := TreeNode{7,nil}
    node14 := TreeNode{1,nil}
    node15 := TreeNode{3,nil}
    node16 := TreeNode{6,nil}
    node17 := TreeNode{9,nil}
    node11.Left = &node12
    node11.Right = &node13
    node12.Left = &node14
    node12.Right = &node15
    node13.Left = &node6
    node13.Right = &node7
    */
    ary := []int{4,2,7,1,3,6,9}
    var root = &TreeNode{}
    
    root = array2btree(ary,root,0,7)
    fmt.Println("print tree:")
    inOrder(root)
    invertTree(root)
    fmt.Println("print tree after invert:")
    inOrder(root)
}
func inOrder(node *TreeNode){
    if node != nil{
        fmt.Printf(" %v",node.Val)
        fmt.Println("/ \\")
        inOrder(node.Left)
        inOrder(node.Right )
    }else{
        fmt.Println("")
    }
}

func newNode(data int) *TreeNode{
    fmt.Printf("set data:%d\n",data)
    node := TreeNode{data,nil,nil}
    return &node
}
func array2btree(arr []int,node *TreeNode,i int,n int) *TreeNode{
    if i < n{
        tmp := newNode(arr[i]) 
        node = tmp

        node.Left = array2btree(arr,node.Left,2*i+1,n) 
        node.Right = array2btree(arr,node.Right,2*i+2,n) 
    }
    return node
}

func invertTree(root *TreeNode) *TreeNode {
    if root == nil{
        return nil

    }
    if root.Left == nil && root.Right == nil{
        return root

    }
    root.Left ,root.Right = root.Right,root.Left
    root.Left =  invertTree(root.Left)
    root.Right = invertTree(root.Right)


    return root

}
