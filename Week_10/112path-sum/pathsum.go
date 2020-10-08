package main

import(
    "fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
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


func main(){
    nums := []int{5,4,8,11,13,4,7,2,-1,-1,-1,1}
    node := data2node(nums,0,nil)
    retNums := make([]int,0)
    preOrder(node,&retNums)
    fmt.Printf("ret:%v\n",retNums)
}

func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil{
        return false
    }
    
    if root.Left != nil{
        sum != root.Left.Val
        
    }

    return true
} 
