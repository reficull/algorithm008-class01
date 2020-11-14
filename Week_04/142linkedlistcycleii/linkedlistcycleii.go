package main
import (
    "fmt" 
)
type ListNode struct {
     Val int
     Next *ListNode
}

func main(){
    
    node11 := ListNode{3,nil}
    node12 := ListNode{2,nil}
    node13 := ListNode{0,nil}
    node14 := ListNode{-4,nil}

    node11.Next = &node12
    node12.Next = &node13
    node13.Next = &node14
    node14.Next = &node12

    node := detectCycle(nil)
    fmt.Println("res:%v",node)
}

func detectCycle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    lMap := make(map[*ListNode]bool,0) 
    for head.Next != nil{
        if lMap[head] == true{
            return head
        }
        lMap[head]=true
        head = head.Next
    } 
    return nil
}
