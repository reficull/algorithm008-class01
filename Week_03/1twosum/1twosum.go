package main
import (
    "fmt" 
)
func main(){
    nums := []int{2, 7, 11, 15}
    target := 26
    ret := twoSum(nums,target)
    fmt.Printf(" nums:%v,target:%d,ans:%v\n",nums,target,ret)
}


func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int,0) 
    for i:=0;i<len(nums);i++{
        q := target - nums[i]
        if _,ok := numMap[q]; ok{
            
            return []int{numMap[q],i}
        }
        numMap[nums[i]] = i
    }
    return nil
}
