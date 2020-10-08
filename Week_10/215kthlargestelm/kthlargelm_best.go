package main

import(
    "fmt"
)

func main(){
    nums := []int{3,2,1,5,6,4}
    key := 2
    fmt.Println("start:",nums," key:",key)
    ret := findKthLargest(nums,key)
    fmt.Println("ans:",ret)
}

type MinPQ struct {
    capacity int
    n        int
    // a[1] 到 a[capacity] 存储数据
    a []int

}



func findKthLargest(nums []int, k int) int {
    mq := MinPQ{}
    mq.init(k)
    for _, num := range nums {

        mq.Push(num)


    }
    fmt.Printf("%+v", mq.a)
    return mq.pop()

}

func (q *MinPQ) init(K int) {
    q.capacity = K
    q.a = make([]int, K+1)

}

func (q *MinPQ) Push(v int) {
    if q.n == q.capacity && v < q.a[1] {
        return

    }
    if q.n == q.capacity {
        q.pop()

    }
    q.n++
    q.a[q.n] = v
    //fmt.Printf("push %d, len %d, n %d\n", v, len(q.a), q.n)
    //exch(q.a, 1, q.n)
    q.swim(q.n)
    //fmt.Printf("after push: %v\n", q.a)

}

func (q *MinPQ) pop() int {
    v := q.a[1]
    exch(q.a,1,q.n)
    q.n--
    q.sink(1)
    //fmt.Printf("--POP: %v\n", v)
    //fmt.Printf("after pop: %v\n", q.a)
    return v

}

func greater(a, b int) bool {
    return a > b

}

func (q *MinPQ) swim(i int) {
    for {
        parent := i / 2
        if parent > 0 && greater(q.a[parent], q.a[i]) {
            exch(q.a,parent,i)

        } else {
            break

        }
        i = i / 2

    }

}

func (q *MinPQ) sink(i int) {
    for i*2 <= q.n {
        left := i * 2
        right := left + 1
        little := left
        if less(q.a[right], q.a[left]) {
            little = right

        }
        if less(q.a[little], q.a[i]) {
            exch(q.a, little, i)

        } else {
            break

        }
        i = little

    }

}


func less(a, b int) bool {
    return a < b

}


func exch(a []int, i, j int) {
    tmp := a[i]
    a[i] = a[j]
    a[j] = tmp

}
