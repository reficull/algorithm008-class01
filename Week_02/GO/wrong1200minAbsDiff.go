package main
import (
    "fmt" 
    "sort"
    "container/heap"
)

func main(){
    //ary := []int{3,8,-10,23,19,-4,-14,27}
    //ary := []int{1,3,6,10,15}
    ary := []int{4,2,1,3}
    ans := minimumAbsDifference(ary)
    fmt.Println("ans:",ans)
}

func minimumAbsDifference(arr []int) [][]int {

    var ans =  [][]int{}
    l := len(arr)
    if l==0{
        return ans
    }
    //resMap := make(map[int]int,0)
    sort.Ints(arr)
    cnt := make([]int,arr[len(arr)-1])
    for i:=0;i<l;i++{ cnt[i]=0 }
    fmt.Println(arr)
    pq := make(PriorityQueue,l-1)
    for i:=0;i<len(arr)-1;i++{

        cnt[i] = arr[i+1]-arr[i]
        pq[i]=&Item{
            value : i,
            priority : cnt[i],
            index:i,
        }
        fmt.Println("check i:",i," i+1",i+1," bigger:",arr[i+1]," gap:",cnt[i], " pq[i]:",pq[i])
    }
    
    fmt.Println("cnt:",cnt)
    heap.Init(&pq)
    fmt.Println("after init:",&pq)
//    printQ(&pq)
    lastGap := 0
    for pq.Len() > 0{
        item := heap.Pop(&pq).(*Item)
        fmt.Println("pair:",item)
        curGap := arr[item.value+1] - arr[item.value]
        if lastGap == 0{
            lastGap = curGap
        }else{
            if lastGap < curGap{
                break
            }
        }
        fmt.Println(" last gap:",lastGap," curGap:",curGap)
        p:= []int{arr[item.value],arr[item.value+1]}
        ans = append(ans,p)
    }
    return ans
}

func printQ(pq *PriorityQueue){
}

type Item struct {
    value    int // The value of the item; arbitrary.
    priority int    // The priority of the item in the queue.
    // The index is needed by update and is maintained by the heap.Interface methods.
    index int // The index of the item in the heap.

}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq)  }

func (pq PriorityQueue) Less(i, j int) bool {
    // We want Pop to give us the highest, not lowest, priority so we use greater than here.
    fmt.Println("less i:",i, " j:",j)
    return pq[i].priority < pq[j].priority

}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j

}

func (pq *PriorityQueue) Push(x interface{}) {
    n := len(*pq)
    item := x.(*Item)
    item.index = n
    *pq = append(*pq, item)

}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    old[n-1] = nil  // avoid memory leak
    item.index = -1 // for safety
    *pq = old[0 : n-1]
    return item

}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
    item.value = value
    item.priority = priority
    heap.Fix(pq, item.index)

}
