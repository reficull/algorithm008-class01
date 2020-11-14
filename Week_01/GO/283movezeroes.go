package main
import (
    "fmt" 
    //"strings"
    //"strconv"
)
// https://leetcode-cn.com/problems/move-zeroes
    
func moveZeroes(nums []int)  {
/*
    ret := make([]int,len(nums));
    var newAry []int;
    var zeroAry []int;
    for _,c := range nums{
        if c != 0{
            newAry = append(newAry,c)
        }else{
            zeroAry = append(zeroAry,c)
        }
    }
    ret = append(newAry,zeroAry...)
    copy(nums,ret)
    */
    // best solution
    l := len(nums)
    count := 0
    for i := 0;i < l;i++{
        if nums[i] == 0{
            count++
            fmt.Println("found 0,i:",i," count:",count)
        }else{
            fmt.Println("not 0,i:",i," count:",count, " swap nums[i - count]:",nums[i-count]," and nums[i]:",nums[i])
            nums[i - count] , nums[i] = nums[i] , nums[i - count]
            fmt.Println("after:",nums)
        }
    }
}

func main(){
    nums := []int{0,1,0,3,12,0}
    fmt.Println("start:",nums)
    moveZeroes(nums)
    fmt.Println("result:",nums," len:",len(nums))
}
