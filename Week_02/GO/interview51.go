package main
import (
    "fmt" 
    "encoding/json"
    "io/ioutil"
)

func main(){
    //nums := []int{7,5,6,4}
    data, err := ioutil.ReadFile("data.json")
    if err != nil {
        fmt.Println(err)
        return

    }
    //fmt.Print("data:  ",string(data))
    var nums []int
    err = json.Unmarshal(data, &nums)
    if err != nil {
        fmt.Println(err)
        return

    }
    //fmt.Printf("slice: %q\n",slice)
    ans := reversePairs(nums)
    fmt.Println("start :",nums," ans:",ans)
}


func reversePairs(nums []int) int {
    ans := 0 
    for i := 0;i < len(nums)-1;i++{
        for j:=i+1;j<len(nums);j++{
            if nums[i]>nums[j]{
                ans++
            }
        }
    } 

    return ans
}
