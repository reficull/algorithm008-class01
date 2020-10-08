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

func findKthLargest(nums []int, k int) int {

    ps := _findKthSmallest(nums,0,len(nums)-1,len(nums)-k)
    return ps 
}

func _findKthSmallest(arr []int,l ,r ,k int)int{
    if l>= r{
        return arr[l]

    }
    p := _partition(arr,l,r)
    if k > p {
        fmt.Println("k:",k , " > p:",p)
        return _findKthSmallest(arr,p+1,r,k)

    }else if k == p{
        fmt.Println("k:",k ," == p:",p,",return arr[p]:",arr[p])
        return arr[p]

    }else{
        fmt.Println("k:",k , " < p:",p)
        return _findKthSmallest(arr,l,p-1,k)

    }

}

func _partition(arr []int,l, r int)int{
    fmt.Println("partition r:",r)
    v := arr[l]
    j := l
    //i := j+1
    for i :=j+1;i<=r;i++{
        if arr[i]<v{
            fmt.Printf("arr[%d]:%d < %d,switch arr[%d]:%d,arr[%d]:%d\n",i,arr[i],v,j+1,arr[j+1],i,arr[i])
            arr[j+1],arr[i] = arr[i],arr[j+1]
            fmt.Printf("after switch %v\n",arr)
            j++
        }

    }
    fmt.Printf("after loop:j:%d,%v\n",j,arr)
    arr[l],arr[j]= arr[j],arr[l]
    fmt.Printf("after switch l,j:%v,return p:%d\n",arr,j)
    return j

}

