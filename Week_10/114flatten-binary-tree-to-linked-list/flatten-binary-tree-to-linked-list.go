package main

import(
    "fmt"
)
/**
best
func flatten(root *TreeNode) {
    if root == nil {
        return

    }
    flatten(root.Left)
    if root.Left != nil {
        l := root.Left
        for l.Right != nil {
            l = l.Right

        }
        l.Right = root.Right
        root.Right = root.Left
        root.Left = nil

    }
    flatten(root.Right)

}

best2
func flatten(root *TreeNode) {
    curr := root
    for curr != nil {
        if curr.Left == nil {
            curr = curr.Right
            continue

        }
        leftMostRight := curr.Left
        for leftMostRight.Right != nil && leftMostRight.Right != curr {
            leftMostRight = leftMostRight.Right

        }
        if leftMostRight.Right == curr {
            leftMostRight.Right, curr.Right, curr.Left = curr.Right, curr.Left, nil
            curr = leftMostRight.Right

        } else {
            leftMostRight.Right = curr
            curr = curr.Left

        }

    }

}

**/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func main(){
    nums := []int{1,2,5,3,4,-1,6}            
    node := data2node(nums,0,nil)
    //inOrder(node)
    flatten(node)
}

func flatten(root *TreeNode)  {
    numsP := make([]int,0)
    preOrder(root,&numsP)
    newNode := TreeNode{}
    root = &newNode
    for _,v := range numsP{
        if v != -1{
            newNode.Val = v
            newNode.Right = &TreeNode{}
        }
    }
    fmt.Printf("np:%v\n",numsP)
}

func preOrder(node *TreeNode,np *[]int){
    if node != nil{
        fmt.Println("pre order:",node.Val)
        *np = append(*np,node.Val)
        preOrder(node.Left,np)
        preOrder(node.Right,np)
    }else{
        fmt.Println(" nil ")
    }
}

func inOrder(node *TreeNode){
    if node != nil{
        fmt.Println("inOrder:",node.Val)
    }
    if node != nil{
        inOrder(node.Left)
        fmt.Println(node.Val)
        inOrder(node.Right)
    }else{
        fmt.Println(" nil ")
    }
}
func data2node(nums []int,i int,p *TreeNode) *TreeNode{
    if (i > len(nums)-1) || (nums[i] == 0)  {
        return nil
    }
    if p != nil{
        fmt.Println("make node:",nums[i]," parent:",p.Val)
    }
    node := &TreeNode{nums[i],nil,nil}
    node.Left = data2node(nums,2*i+1,node)
    node.Right = data2node(nums,2*i+2,node)

    return node

}

