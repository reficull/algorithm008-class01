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

    ret := levelOrder(&node1)
    fmt.Println("ret:",ret)
}

func levelOrder(root *Node) [][]int {
    var res [][]int
    var arr[]*Node
    if root!=nil{
        arr=append(arr,root)
        for len(arr)!=0{
            temp:=len(arr)
            var level []int
            for i:=0;i<temp;i++{
                level=append(level,arr[i].Val)
                for _,v:=range arr[i].Children{
                    arr=append(arr,v)
                }

            }
            res=append(res,level)
            arr=arr[temp:]
        }
    }
    return res
    /* use queue
    var ans [][]int
    if root == nil {
        return ans

    }
    var queue []*Node
    queue = append(queue, root)
    for len(queue) != 0 {
        size := len(queue)
        var level = make([]int, size)
        for i := 0; i < size; i++ {
            level[i] = queue[i].Val
            queue = append(queue, queue[i].Children...)

        }
        ans = append(ans, level)
        fmt.Println("size:",size)
        queue = queue[size:]

    }
    return ans

    */
} 

