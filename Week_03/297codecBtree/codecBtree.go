package main
import (
    "fmt" 
    "encoding/json"
    "strconv"
    "strings"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func main(){
    str := "[1,2,3,null,null,4,5]"
    codec := Codec{}
    root := codec.deserialize(str)
    fmt.Println("deserializes root:",root)
    //inOrder(root)
    
    fmt.Println("Serializes:",codec.serialize(root))
}

type Codec struct {

    root *TreeNode;
}

func Constructor() Codec {
    codec := Codec{}
    return codec
}

func inOrder(node *TreeNode){
    if node != nil{
        fmt.Println(node.Val)
        inOrder(node.Left)
        inOrder(node.Right)
    }else{
        fmt.Println(" nil ")
    }
}
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    var strData []string
    var queue = []*TreeNode{root}
    for len(queue) > 0{
        if queue[0]!= nil{
            strData = append(strData,strconv.Itoa(queue[0].Val))
            queue = append(queue,queue[0].Left,queue[0].Right)
        }else{
            strData = append(strData,"#")
        }
        queue = queue[1:]
    } 
    fmt.Println("strData:",strData)
    return strings.Join(strData,",") 
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    var nums []int
    //nums := []int 
    err := json.Unmarshal([]byte(data), &nums)
    if err != nil{
        fmt.Printf("str to json err:%v\n",err)
    }
    fmt.Println("nums:",nums)
    this.root = this.data2node(nums,0,nil)
    return this.root
}

func (this *Codec) data2node(nums []int,i int,p *TreeNode) *TreeNode{
    if (i > len(nums)-1) || (nums[i] == 0)  {
        return nil
    }

    fmt.Println("make node:",nums[i])
    node := &TreeNode{nums[i],nil,nil}
    node.Left = this.data2node(nums,2*i+1,node)
    node.Right = this.data2node(nums,2*i+2,node)

    return node

}


/**
* Your Codec object will be instantiated and called as such:
* obj := Constructor();
* data := obj.serialize(root);
* ans := obj.deserialize(data);
*/
