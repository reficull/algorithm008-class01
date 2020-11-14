
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
    nums := []int{1,2,5,3,4,-1,6}            
    node := data2node(nums,0,nil)
    flatten(node)
    nodeList := make([]int,0)    
    preOrder(node,&nodeList)
    fmt.Printf("ret:%v\n",nodeList)
}

func flatten(root *TreeNode)  {   
    if root == nil{
        return
    } 

    flatten(root.Left)
    if root.Left != nil{
        l := root.Left
        for l.Right != nil{
            l = l.Right
        } 
        l.Right , root.Right,root.Left =  root.Right,root.Left,nil
    }
    
    flatten(root.Right)
    
}
