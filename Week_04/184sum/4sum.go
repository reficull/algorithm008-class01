package main
import (
    "fmt" 
    "sort"
)

func main(){
    nums := []int{1, 0, -1, 0, -2, 2}
    target := 0
    ans := fourSum(nums,target)
    printMatrix(ans)
}


func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums) 
    fmt.Printf("sorted:%v\n",nums)
    l := len(nums)
    var ans  [][]int
    p1,p2,p3,p4 := 0,1,l-2,l-1
    //max value < target return
    if nums[p4]+nums[p3]+nums[p3-1] + nums[p3-2] < target{
        return ans
    }
    //min value > target return
    if nums[p1]+nums[p2]+nums[p2+1] + nums[p2+2] > target{
        return ans
    }
    for p1=0;p1 < l-3;p1++{
        if p1 > 0 && nums[p1]==nums[p1-1]{
            continue
        }
        for p2=p1+1;p2 < p3;p2++{
            if p2>p1+1 &&  nums[p2] == nums[p2-1]{
                p2++
                continue
            } 
            p3 = p2+1
            for p3 < p4{
                cur := nums[p1] + nums[p2] + nums[p3] + nums[p4]
                if cur == target{
                    ary := []int{nums[p1] , nums[p2] , nums[p3] , nums[p4]}
                    fmt.Printf("inner p1=%d,p2=%d,p3=%d,p4=%d\n",p1,p2,p3,p4)
                    ans = append(ans,ary)
                    p3++
                    for p3<p4 && nums[p3]==nums[p3-1]{
                        p3++
                    }
                    p4--
                    for p3<p4 && nums[p4]==nums[p4+1]{
                        p4--
                    }

                }else if cur < target{
                    p3++
                }else{
                    p4--
                }
            }
            
        }
    }
    return ans
}

func printMatrix(matrix [][]int){
    for r,col1 := range matrix{
        fmt.Printf("row %d:%v\n",r,col1)
    }
}

/*
            if nums[p1] + nums[p2] + nums[p3] + nums[p4] == target{
                ary := []int{nums[p1] , nums[p2] , nums[p3] , nums[p4]}
                fmt.Printf("inner p1=%d,p2=%d,p3=%d,p4=%d\n",p1,p2,p3,p4)
                ans = append(ans,ary)
            }
            if nums[p1] + nums[p2] + nums[p3] + nums[p4] < target{
                p2++
                continue
            }else{
                p3--
                continue
            }
        if nums[p1] + nums[p2] + nums[p3] + nums[p4] == target{
            ary := []int{nums[p1] , nums[p2] , nums[p3] , nums[p4]}
            fmt.Printf("outter p1=%d,p2=%d,p3=%d,p4=%d\n",p1,p2,p3,p4)
            ans = append(ans,ary)
        }
        if nums[p1] + nums[p2] + nums[p3] + nums[p4] < target{
            p1++
            continue
        }else{
            p4--
            continue
        }

*/
