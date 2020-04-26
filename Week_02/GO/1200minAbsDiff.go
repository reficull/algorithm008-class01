package main
import (
    "fmt" 
    "sort"
)

func main(){
    ary := []int{40,11,26,27,-2}
    //ary := []int{3,8,-10,23,19,-4,-14,27}
    //ary := []int{1,3,6,10,15}
    //ary := []int{4,2,1,3}
    ans := minimumAbsDifference(ary)
    fmt.Println("ans:",ans)
}



func minimumAbsDifference(arr []int) [][]int {
    sort.Ints(arr)
    fmt.Println()
    min := 100000
    var ans = [][]int{}
    //get min gap
    for i:=0;i<len(arr)-1;i++{
        if min > (arr[i+1]-arr[i]){
            min = arr[i+1]-arr[i]
        }
    }

    for i:=0;i<len(arr)-1;i++{
        if min == arr[i+1]-arr[i]{
            ans = append(ans,[]int{arr[i],arr[i+1]})
            fmt.Println("append res:",ans) 
        }
    }
    return ans
}
