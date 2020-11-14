package main
import (
    "fmt"
)

func moveZeroes(nums []int)  {  
    //test all zeroes
    flag := 0;
    for _,num := range nums{
        if num != 0{
            flag = 1 
        }
    }
    if flag == 0{
        return
    }
    tmp := make([]int,len(nums))
    move1(nums,0,&tmp) 
    fmt.Println("get result:",tmp)
    copy(nums,tmp) 
}

func move1(nums []int,i int,res *[]int) {
    fmt.Println("i:",i)
    if i > len(nums)-1 {
        fmt.Println("recursive end,res:",nums)
        *res = nums
        return
    }
    tmp := make([]int,len(nums))
    copy(tmp,nums)
    if i== 0{
        //first digit is 0
        if nums[i] == 0{
            copy(tmp,append(nums[1:],0))
            fmt.Println("tmp:",tmp)
            //copy(nums,tmp)
            fmt.Println("first digit 0,result:",tmp)
            
            if tmp[0] == 0{
                fmt.Println("next digit 0,do it again:",tmp)

                move1(tmp,0,res)

            }else{
                move1(tmp,1,res)
            }
            
        }
    }else{
        if nums[i] == 0{
            if i==len(nums)-1 {
                return
            }
            copy(tmp,append(nums[0:i],nums[i+1:]...))
            copy(tmp,append(nums,0))
            fmt.Println("not first digit 0,result:",tmp)

        }else{
            fmt.Println("not zero,copy ary")
            copy(tmp,nums)
        }
        j := i+1
        fmt.Println("check next digit:",j)
        move1(tmp,j,res)
    }
    //fmt.Printf("returning:%v, i:%d",nums,i)

}


func main(){
    nums := []int{1,0}
    fmt.Println("begin:",nums)
    moveZeroes(nums)
    fmt.Println("result:",nums)
}
