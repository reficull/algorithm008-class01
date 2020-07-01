package main

import(
    "fmt"
    "math"
)

func main(){
    nums := []int{2,3,1,2,4,3}
    target := 7
    ret := minSubArrayLen(target,nums)
    fmt.Println("ret:",ret)
}

func minSubArrayLen(s int, nums []int) int {
    l := len(nums)
    if l<2{
        return 0
    }
    sum,start,end := 0,0,0 
    ans := math.MaxInt32 
    for end < l{
        sum += nums[end] 
        for sum >= s{
            sum -= nums[start]
            start++
            if sum == s{ans = Min(ans,end-start+1)} 
            //fmt.Printf("start:%d,end:%d,sum:%d,ans:%d\n",start,end,sum,ans  )
        }
        end++
    }
    return ans
}

func Min(a,b int) int{
    if a < b{
        return a
    }
    return b
}
