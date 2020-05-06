package main
import (
    "fmt" 
    "math"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
var valid = true

func main(){
    //bArr := []int{5,1,4,0,0,3,6}
    bArr := []int{2,3,1}
    //bArr := []int{2,1,3}
    //bArr := []int{}


    var root = &TreeNode{}
    
    root = array2btree(bArr,root,0)
    inOrder(root)
    ans := isValidBST(root)
    fmt.Println("ans:",ans)
}
func isValidBSTNumber(root *TreeNode, min, max int) bool {

    if root == nil {
        return true
    }

    if root.Val <= min  || root.Val >= max{
        return false
    }

    return isValidBSTNumber(root.Left, min, root.Val) && isValidBSTNumber(root.Right, root.Val, max)
}

func isValidBST(root *TreeNode) bool {
    return isValidBSTNumber(root, math.MinInt64, math.MaxInt64)
}

/* best 2
func isValidBST(root *TreeNode) bool {
    return dfs(root,-1<<63,1<<63-1)
}

func dfs(root *TreeNode,min,max int) bool{
    if root != nil {fmt.Printf("node:%v,min:%d,max:%d\n",root.Val,min,max) }
    
    return root == nil || min < root.Val && root.Val < max &&
        dfs(root.Left,min,root.Val) &&
        dfs(root.Right,root.Val,max)
}
*/

func checkValid(root *TreeNode){
    if root == nil{ return }
    if root.Left != nil && root.Left.Val > root.Val{ valid = false;return } 
    if root.Right != nil && root.Right.Val < root.Val{ valid = false;return } 
    checkValid(root.Left)
    checkValid(root.Right)
}

func inOrder(node *TreeNode){
    if node != nil{
        fmt.Println(node.Val)
        inOrder(node.Left)
        inOrder(node.Right)
    }else{
        //fmt.Println(" \n")
    }
}
func newNode(data int) *TreeNode{
    fmt.Printf("set data:%d\n",data)
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
