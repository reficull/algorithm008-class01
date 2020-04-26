package main
import (
    "fmt" 
)

func main(){
    //nums := []int{0,1,2}
    //nums := []int{0,2}
   // nums := []int{2,2}
    //nums := []int{2,1,2}
    //nums := []int{1,2,0}
    //nums := []int{2,0,2,1,1,0}
    cases := [][]int{
        {2,0,2,1,1,0},
        {1,2,0},
        {2,1,2},
        {2,2},
        {0,2},
        {0,1,2},
    }
    correct := [][]int{
        {0,0,1,1,2,2},
        {0,1,2},
        {1,2,2},
        {2,2},
        {0,2},
        {0,1,2},
    }
    for i,nums := range cases{
        fmt.Println("start:",nums)
        sortColors(nums)
        fmt.Println("res:",nums)
        fmt.Println("cor:",correct[i])
    } 
}

func sortColors(nums []int)  {
    l := len(nums)
    for p0,p1,p2 := 0,0, l-1 ;(p1<=p2);{
        if nums[p1]==0{
            nums[p0] , nums[p1] = nums[p1],nums[p0]
            p0++;p1++
            fmt.Println("move 0:",nums)
        }else if nums[p1]==2{
            nums[p1],nums[p2] = nums[p2],nums[p1]
            p2--;
            p0--;if p0<0 {p0 = 0}
            p1=p0
            fmt.Println("move 2:",nums, " p0,p1,p2",p0,p1,p2)
        }else{
            p1++
        }
        //if p1 >= p2{ return }
    }
}


/*
//2 pointers

    cntZero := 0
    cntTwo := 0
    l := len(nums)
    for i:=0;i<len(nums);i++{
        fmt.Println("check i:",i)
        if nums[i]==0 {
            nums[cntZero],nums[i] = nums[i],nums[cntZero]
            cntZero++
            fmt.Println("move 0:",nums)
        }

        if nums[i]==2 && (i < l-cntTwo){
            fmt.Println("found 2 i:",i)
            cntTwo++
            for{
                if cntTwo<l && nums[l-cntTwo] == 2 && nums[i]!=0{
                fmt.Println("found 2 at the endi:",i)
                    if cntTwo == l{return}
                    cntTwo++
                }else{
                    break
                }
            }
            fmt.Println("before move 2:",nums," i:",i)
            if i >= (l-cntTwo) { continue}
            nums[l-cntTwo],nums[i] = nums[i],nums[l-cntTwo]
            fmt.Println("move 2:",nums)
            if nums[i]==0{
                nums[cntZero],nums[i] = nums[i],nums[cntZero]
                fmt.Println("move 0:",nums)
                cntZero++
            }
        }
    }
    */
