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
    node12 := ListNode{2,nil}
    node13 := ListNode{3,nil}
    node14 := ListNode{4,nil}
    node15 := ListNode{5,nil}
    node11.Next = &node12
    node12.Next = &node13
    node13.Next = &node14
    node14.Next = &node15

    printNodes(&node11)
    ans := reverseList(&node11)
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
    var temp *ListNode
    cur := head
    for nil != cur{
        temp = cur.Next
        cur.Next = pre
        pre = cur
        cur = temp
        if pre != nil && cur != nil{
            fmt.Printf("cur:%v,pre:%v \n",cur.Val,pre.Val)
        }
        printNodes(cur)
    }
    return pre
}

    /*
    if head == nil || head.Next == nil{
        return head
    }
    var cur *ListNode = nil
    cur = reverseList(head.Next)
    cur.Next.Next = cur
    cur.Next = nil
    return cur
    */
