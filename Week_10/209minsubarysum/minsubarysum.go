package main

import(
    "fmt"
)

func main(){
    nums := []int{2,3,1,2,4,3}
    target := 7
    ret := minSubArrayLen(target,nums)
    fmt.Println("nums:",nums)
    fmt.Println("ret:",ret)
}

func minSubArrayLen(s int, nums []int) int {
    l := len(nums)
    if l<2{
        return 0
    }
    minLen := l
    sum,start,end := 0,0,0 
    for  end < l{
        sum += nums[end] 
        for sum >= s{
            sum -= nums[start]
            start++
            fmt.Printf("start:%d,end:%d,sum:%d,ans:%d\n",start,end,sum,minLen)
            minLen = Min(minLen,end-start+1)
        }
        end++
    }
    if minLen == len(nums){
        minLen = 0
    }
    return minLen 
}

func Min(a,b int) int{
    if a < b{
        return a
    }
    return b
}
