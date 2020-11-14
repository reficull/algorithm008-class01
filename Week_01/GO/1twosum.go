package main
import (
    "fmt" 
)
// https://leetcode-cn.com/problems/two-sum/
func main(){
    nums := []int{2,7,11,15}
    target := 9 
    fmt.Println("start:",nums," target:",target)
    res := twoSum(nums,target)
    fmt.Println("res:",res)
}

func twoSum(nums []int, target int) []int {
    /*
    first := 0 
    for i := 0;i < len(nums);i++{
        first = nums[i]
        for j := i+1;j<len(nums);j++{
            if nums[j] == target - first{
                return []int{i,j}
            }
        }
    }
    return []int{}
    */
    res := make(map[int]int, 0)
    fmt.Println("map start:",res)
    for i, num := range nums {
        res[num] = i
        fmt.Println("map in for:",res)
        if j, ok := res[target-num]; ok {
            fmt.Println("test: ",target-num, " j:",j)
            return []int{i, j}
        }
        fmt.Println("set map:",num, " =",i)

    }
    return nil
}
