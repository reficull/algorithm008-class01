package main
import (
    "fmt" 
)

const MaxUint = ^uint(0) 
const MinUint = 0 
const MaxInt = int(MaxUint >> 1) 
const MinInt = -MaxInt - 1

func main(){
    //nums := []int{ -4,-3}
    nums := []int{ -1,-2,-9,-6}
    //nums := []int{ 3,-1,4}
    //nums := []int{ 2,3,-2,4}
    //nums := []int{ 2,-5,-2,-4,3}
   
   //nums := []int{ 0,2}

    fmt.Println("ret:",nums)

    ret := maxProduct(nums)
    fmt.Println("ret:",ret)
}


func maxProduct(nums []int) int {
    l := len(nums)
    dpMax := make([]int,l)
    copy(dpMax ,nums)
    dpMin := make([]int,l)
    copy(dpMin ,nums)
    for i:=1;i<len(nums);i++{

        dpMax[i] = Max(dpMax[i-1] * nums[i], Max(dpMin[i-1] * nums[i],nums[i])) 
        dpMin[i] = Min(dpMin[i-1] * nums[i], Min(dpMax[i-1] * nums[i],nums[i])) 

        fmt.Printf("dpMin[i-1](%d) * nums[i](%d):%d dpMax[i-1] * nums[i]:%d dpMin[i]:%d\n",dpMin[i-1],nums[i], dpMin[i-1] * nums[i],dpMax[i-1] * nums[i],dpMin[i])
    }
    fmt.Println(dpMax)
    fmt.Println(dpMin)
    max1 := maxInAry(dpMax)
    return max1
}

func maxInAry(nums []int) int{
    m:= MinInt  
    for i:=0;i<len(nums);i++{
        m =Max(m, nums[i])
    }
    return m 
}

func Min(a, b int) int{
    if a < b{
        return a
    }
    return b
}

func Max(a, b int) int{
    if a > b{
        return a
    }
    return b
}
