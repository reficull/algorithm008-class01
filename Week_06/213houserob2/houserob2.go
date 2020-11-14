package main
import (
    "fmt" 
)

const MaxUint = ^uint(0) 
const MinUint = 0 
const MaxInt = int(MaxUint >> 1) 
const MinInt = -MaxInt - 1

func main(){
    nums := []int{1}
    
    //nums := []int{}
    //nums := []int{2,3,2}
    //nums := []int{1,2,1,1}
    //nums := []int{1,2,3,1}
    //nums := []int{2,1,1,2}
    //nums := []int{0}
    //nums := []int{1,2,3,1,3,2}
    fmt.Println(nums)
    ret := rob(nums)
    fmt.Printf("ret:%d\n",ret)
}

func rob(nums []int) int {
    res := 0
    l := len(nums)
    if l == 0{
        res = 0
    }else if l==1 {
        return nums[0] 
    }else{
        ret1 := doRob(nums[:len(nums)-1])
        ret2 := doRob(nums[1:len(nums)])
        res = Max(ret1,ret2)

    }
    return res
}

func doRob(nums []int) int{
    l := len(nums)
    if l == 0{
        return 0

    }
    if l == 1{
        return  nums[0]

    }
    dp := make([]int,l)
    dp[0]=nums[0]
    dp[1]=Max(nums[0],nums[1])
    for i:=2;i<l;i++{
        dp[i] =  Max( dp[i-1],dp[i-2]+nums[i] )

    }
    fmt.Println(dp)
    maxA := maxInAry(dp[:len(dp)-1])
    maxB := maxInAry(dp[1:])
    fmt.Printf("max a:%d,maxb:%d\n",maxA,maxB)

    return Max(maxA,maxB)
}

func Max(a, b int) int{
    if a > b{
        return a
    }
    return b
}

func printMatrix(matrix [][]int){
    for r,col1 := range matrix{
        fmt.Printf("row %d:%v\n",r,col1)
    }
}

func maxInAry(nums []int) int{
    m:= MinInt  
    for i:=0;i<len(nums);i++{
        m =Max(m, nums[i])
    }
    return m 
}
