package main
import "fmt" 

func rotate(nums []int, k int)   {
     cut :=  len(nums) - k % len(nums);
     copy(nums,append(nums[cut:],nums[0:cut] ...));       
}

func main(){
    nums := []int{1,2,3,4,5,6,7}
    rotate(nums,3)
    fmt.Printf("result:%v\n",nums)
}
