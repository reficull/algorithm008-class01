package main
import (
    "fmt" 
)

const MaxUint = ^uint(0) 
const MinUint = 0 
const MaxInt = int(MaxUint >> 1) 
const MinInt = -MaxInt - 1

func main(){
    nums:=[]int{-2,1,-3,4,-1,2,1,-5,4}
    //nums:=[]int{-2,-1}
    fmt.Println(nums)
    ret := maxSubArray(nums)
    fmt.Println("ans:",ret)
}

func maxSubArray(nums []int) int {
    if len(nums)==1{
        return nums[0]
    }
    // max_sum(i) = Max(max_sum(i-1)+i,max_sum(i-1))
    dp := make([]int,len(nums)) 
    copy(dp ,nums)
    for i:=1;i<len(nums);i++{
        if dp[i-1]>0{
            dp[i] = Max(0,dp[i-1]) + nums[i] 
        }else{
     //       dp[i] = nums[i]
        }
    }
    fmt.Println(dp)
    return maxInAry(dp)

}

func maxInAry(nums []int) int{
    m:= MinInt  
    for i:=0;i<len(nums);i++{
        m =Max(m, nums[i])
    }
    return m 
}

func Max(a, b int) int{
    if a > b{
        return a
    }
    return b
}
