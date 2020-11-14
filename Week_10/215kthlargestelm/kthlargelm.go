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

    ps := quickSort(nums,0,len(nums)-1,k)
    return ps 
}

func quickSort(nums []int,l,r,k int) int{
    if l > r{return -1}
    
    div := partSort(nums,l,r)
    if div == k-1 { 
        return div 
    } else if div < k-1{ 
        return quickSort(nums,div+1,r,k)  
    } else { 
        return quickSort(nums,l,div-1,k) 
    }

    quickSort(nums,l,div-1,k)
    quickSort(nums,div+1,r,k)
    return -1
}


func partSort(nums []int,l, r int) int{
    key := nums[r]
    last := l
    for l < r{
        for l < r && nums[l]>=key{
            fmt.Printf("nums[%d] :%d > key:%d,l:%d ++:%d\n",l,nums[l],key,l,l+1)
            l++
        }

        for l < r && nums[r]<=key{
            r--
        }
        fmt.Printf("found left nums[%d] :%d < key:%d,right nums[%d]:%d>key:%d \n",l,nums[l],key,r,nums[r],key)
        fmt.Printf("swap %d,%d,%v\n",l,r,nums)
        nums[l],nums[r] = nums[r],nums[l]
        fmt.Printf("swap aftger %v\n",nums)
    }
    fmt.Printf("swap begin : %d,last: %d,%v\n",l,r,nums)

    nums[l],nums[last] = nums[last],nums[l]
    fmt.Printf("swap after: %v\n",nums)
    return l
}
