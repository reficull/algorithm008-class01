package main
import "fmt" 

func moveZeroes(nums1 []int)  {
    nums := nums1
    fmt.Println(nums)
    lastNotZero := len(nums)-1
    for i:=0;i<len(nums);i++{
        if (nums)[i] == 0 && i <= lastNotZero{
            pop2end(&nums,i,lastNotZero)
            lastNotZero--;
        }
    }
}

func pop2end(nums *[]int,i int,l int){
    fmt.Println("pop to end:",i,l)
    for j:=i;j+1<=l;j++{
        fmt.Println("swap",(*nums)[j],(*nums)[j+1])
        tmp:=(*nums)[j]
        (*nums)[j] = (*nums)[j+1]
        (*nums)[j+1] = tmp
        fmt.Println("result ",(*nums)[j],(*nums)[j+1])
    }
    fmt.Println("result:",nums)
}


func main(){
    nums := []int{0,0,32}
    moveZeroes(nums)
    fmt.Println(nums)
}


