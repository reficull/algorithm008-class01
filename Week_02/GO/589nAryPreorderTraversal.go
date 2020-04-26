package main
import (
    "fmt" 
)

type Node struct {
    Val int
    Children []*Node
}

func main(){
    node1 := Node{1,nil}
    node2 := Node{2,nil}
    node3 := Node{3,nil}
    node4 := Node{4,nil}
    node5 := Node{5,nil}
    node6 := Node{6,nil}

    node1.Children= []*Node{&node3,&node2,&node4}
    node3.Children= []*Node{&node5,&node6}

    ret := preorderTraversal(&node1)
    fmt.Println("ret:",ret)
}

func preorderTraversal(root *Node) []int {
    res := make([]int,0)
    doTraverse(root,&res)
    return res

}

func doTraverse(root *Node,res *[]int){
    if root == nil{
        return

    }
    *res = append(*res,root.Val)
    for i:=0;i<len(root.Children);i++{
        doTraverse(root.Children[i],res)
    }

}
