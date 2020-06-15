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
    nums := []int{3,5,1,6,2,0,8,-1,-1,7,4}
    root := array2btree(nums,nil,0)
    nums1 := btree2array(root)
    fmt.Printf("after bfs:%v\n",nums1)

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
      
    return nil
}

func newNode(data int) *TreeNode{
    //fmt.Printf("new node set data:%d\n",data)
    node := TreeNode{data,nil,nil}
    return &node
}
func array2btree(arr []int,node *TreeNode,i int) *TreeNode{
    if i < len(arr){
        tmp := newNode(arr[i]) 
        node = tmp

        node.Left = array2btree(arr,node.Left,2*i+1) 
        node.Right = array2btree(arr,node.Right,2*i+2) 
    }
    return node
}
/*
    如果node 有左右孩子，则加入队列
    如果node 左孩子是空，队列加个值为-1，无子树的节点
    如果node 右孩子是空，队列加个值为-1, 无子树的节点
    如果node 值为-1，结果中加-1,不加入队列
*/
func btree2array(root *TreeNode) []int{
    res := []int{}
    var queue  =[]*TreeNode{root}
    for len(queue) > 0{
        node := queue[0]
        queue = queue[1:]
        fmt.Printf("process %d\n",node.Val) 
        res = append(res,node.Val)
        if node.Left != nil {
            fmt.Printf("node %d left is %d append to queue \n",node.Val,node.Left.Val)
            queue = append(queue,node.Left)
        }
        if node.Val != -1 && node.Left == nil{
            n := &TreeNode{-1,nil,nil}
            queue = append(queue,n)
        }
        if node.Right!= nil{
            fmt.Printf("node %d right is %d append to queue\n",node.Val,node.Right.Val)
            queue = append(queue,node.Right)
        }
        if node.Val != -1 && node.Right == nil{
            n := &TreeNode{-1,nil,nil}
            queue = append(queue,n)
        }
        fmt.Printf("queue len:%d",len(queue))
    }
    return res

}
func getI(){


}
