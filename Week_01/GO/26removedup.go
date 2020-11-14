package main
import (
    "fmt"
    //"math"
)

func main(){
    nums := []int{0,0,1,1,1,2,2,3,3,4}
    //nums := []int{1,1,2}
    fmt.Println("start :",nums) 
    cnt := removeDuplicates(nums)
    fmt.Println("count:",cnt," result:",nums) 
}

func removeDuplicates(nums []int) int {
    /* best solution:
    */
    ln := len(nums)
    n := 1
    for i := 1; i < ln; i++ {
        if nums[i] != nums[i-1] {
            fmt.Println(" nums[i] != nums[i-1],i:",i, nums[i],nums[i-1] ," set nums[n] = num[i],nums[n]:",nums[n])
            nums[n] = nums[i]
            fmt.Println("after : nums[n]:",nums[n], " nums:",nums)
            n++
        }
    }
    return n

    /*
    cnt := 0
    if len(nums) == 1 {
        cnt = 1
        return cnt
    }
    for s,f := 0,1;f<len(nums); {
        if nums[s] < nums[f]{
            s++
            nums[s] = nums[f]
            cnt++
        }else if nums[s] == nums[f]{
            f++
        }
    }
    cnt++
    fmt.Println(" res:",nums, " count:",cnt)
    return cnt
   */ 

}
