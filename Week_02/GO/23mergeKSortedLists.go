package main
import (
    "fmt" 
)

type ListNode struct {
    Val int
    Next *ListNode
}


func main(){
    node11 := ListNode{1,nil}
    node12 := ListNode{4,nil}
    node13 := ListNode{5,nil}
    node11.Next = &node12
    node12.Next = &node13
    
    node21 := ListNode{1,nil}
    node22 := ListNode{3,nil}
    node23 := ListNode{4,nil}
    node21.Next = &node22
    node22.Next = &node23
    
    node31 := ListNode{2,nil}
    node32 := ListNode{6,nil}
    node31.Next = &node32

    lists := []*ListNode{&node11,&node21,&node31}
    ans := mergeKLists(lists)
    printNodes(ans)
}

func printNodes(node *ListNode){

    for {
        fmt.Printf( "%d",node.Val )
        if node.Next != nil{
            node = node.Next
            fmt.Printf("->")
        }else{
            fmt.Println("\n")
            break
        }
    }
}

func reverseList(head *ListNode) *ListNode {
    var pre *ListNode 
    //var temp *ListNode         
    cur := head                        
    for nil != cur{                    
        temp := cur.Next               
        cur.Next = pre                 
        pre = cur                      
        cur = temp                     
    }                              
    return pre                         

}

func mergeKLists(lists []*ListNode) *ListNode {
    p1 := lists[0]
    p2 := lists[1]
    //p3 := lists[2]
    var newHead = ListNode{0,nil}
    for {

        if p1.Val <= p2.Val{
            newHead.Val = p1.Val
            newHead.Next = &ListNode{p2.Val,nil}
            newHead.Next.Next = &ListNode{p1.Val,nil}
            newHead = *newHead.Next.Next

        }else{
            newHead.Val = p2.Val
            newHead.Next = &ListNode{p1.Val,nil}
            newHead.Next.Next = &ListNode{p2.Val,nil}
            newHead = *newHead.Next.Next
        }
        printNodes(&newHead)
        p1 = p1.Next
        p2 = p2.Next

        if p1 == nil {
            break
        }
        if p2 == nil {
            break
        }
    }
    r := reverseList(&newHead)
    return r

}


