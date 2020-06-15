package main
import (
    "fmt" 
)

func main(){
    nums := []int{2,3,1,1,4}
    ret := jump(nums)
    fmt.Printf("ans:%d",ret)
}

func jump(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0

    }
    cnt, pre, reached := 0, 0, nums[0] //pre记录上一个能够抵达的最大位置， reached记录能够抵达的最大位置
    for i := 0; i <= reached && i < n; i++ {
        if i > pre { // 当前位置大于上一个能够抵达的最大位置时，更新
            pre = reached
            cnt++
        }
        reached = Max(nums[i]+i, reached) // 能够到达最远的位置

    }
    return cnt

}

func Max(a ,b int) int{
    if a > b{
        return a
    }
    return b
}

