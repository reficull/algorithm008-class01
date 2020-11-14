/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 package main

 import (
	 "fmt"
 )
 type ListNode struct {
	     Val int
	     Next *ListNode
	 }
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// loop  l1 and l2
// add 2 num and get result put into array[0]
// if > 10 than 
	ret := &ListNode{0,nil}
	add2andret(l1,l2,ret)
	return ret
}

func add2andret(l1 *ListNode,l2 *ListNode,l3 *ListNode) int{
	ret := 0
	if l1.Next == nil {
		l3.Val = l1.Val + l2.Val;
		
	}else{
		l3.Next = &ListNode{0,nil}
		l3.Val = l1.Val + l2.Val + add2andret(l1.Next,l2.Next,l3.Next)
	}
	fmt.Printf("L1:%v,l2:%v,l3:%v\n",l1.Val,l2.Val,l3.Val)
	if l3.Val > 9{
		l3.Val -= 10
		fmt.Printf(">10 fix l3:%v\n",l3.Val)
		ret = 1
	}
	return ret
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
    }
    return pre
}   

func printNodes(node *ListNode){
    for {
        fmt.Printf( "%d",node.Val )
        if node.Next != nil{
            node = node.Next
            fmt.Printf("->")
        }else{
            fmt.Println()
            break
        }
    }
}   

func main(){
	node1 := ListNode{2,nil}
    node2 := ListNode{4,nil}
	node3 := ListNode{3,nil}
	node1.Next = &node2
	node2.Next = &node3

	node4 := ListNode{5,nil}
	node5 := ListNode{6,nil}
	node6 := ListNode{4,nil}
	node4.Next = &node5
	node5.Next = &node6

	tmp := addTwoNumbers(&node1,&node4)
	ans := reverseList(tmp);
	printNodes(ans)

}
// @lc code=end

