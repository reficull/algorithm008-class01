package main

import(
    "fmt"
)

func main(){
    cache := Constructor(2)  
    
    cache.Put(2,1)
    cache.Put(1,1)
    printLL(cache.Head)
    cache.Put(2,3)
    fmt.Printf("put %d\n",3)
    printLL(cache.Head)
    cache.Put(4,1)
    fmt.Printf("put %d\n",4)
    printLL(cache.Head)
    ret := cache.Get(1)
    fmt.Printf("get %d:%d size:%d\n",1,ret,len(cache.m))
    ret = cache.Get(2)
    printLL(cache.Head)
    fmt.Printf("get %d:%d\n",2,ret)
}


type Node struct{
    Key int
    Val int
    Next *Node
    Prev *Node
}

type LRUCache struct {
    Cap int
    Head *Node  
    Tail *Node
    m   map[int]*Node
    size int
}

func (this *LRUCache) Move2Front(node *Node){
    if node == this.Head.Next  {
        return
    }
    this.Remove(node) 
    this.Add2Front(node)

}

func (this *LRUCache) Remove(node *Node){
    node.Prev.Next = node.Next
    node.Next.Prev = node.Prev
    this.size--
}
func (this *LRUCache) RemoveLast() *Node{
    tmp := this.Tail.Prev
    this.Tail.Prev= this.Tail.Prev.Prev
    this.Tail.Prev.Next = this.Tail
    this.size--
    return tmp
}


func Constructor(capacity int) LRUCache {
    this := LRUCache{
        capacity,
        &Node{0,0,nil,nil},
        &Node{0,0,nil,nil},
        make(map[int]*Node,capacity),
        0,
    } 
    this.Head.Next = this.Tail
    this.Tail.Prev = this.Head
    return this
}


func (this *LRUCache) Get(key int) int {
    if this.Head.Next== nil{
        return -1 
    }
    fmt.Println("get :",key)
    if node,ok := this.m[key];!ok {
        return -1
    }else{
        this.Move2Front(node)
        return node.Val
    }
    
}

func (this *LRUCache) Put(key int, value int)  {
    //check size
    if _,ok := this.m[key];!ok {
        newNode := &Node{key,value,nil,nil}
        this.m[key] = newNode
        this.Add2Front(newNode)

        fmt.Printf("after put m size:%d\n",len(this.m))
        if this.Size()> this.Cap{
            //remove last one
            last := this.RemoveLast()
            delete(this.m,last.Key)
        }
    }else{
        node := this.m[key]
        node.Val = value
        this.Move2Front(node)
    }

    fmt.Printf("m:%v\n",this.m)
}

func (this *LRUCache) Add2Front(node *Node){
    node.Next = this.Head.Next
    node.Prev = this.Head

    this.Head.Next.Prev = node
    this.Head.Next = node
    this.size++
}

func (this *LRUCache) Size() int {
    return this.size
}

func printLL(node *Node){
    if node == nil{
        return 
    }
    for node != nil{
        fmt.Printf("[k:%d->v:%d]->",node.Key,node.Val)
        node = node.Next
    }
    fmt.Println()
}
