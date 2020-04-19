package main
import (
    "fmt" 
)


func main(){
    nums := []int{1,8,6,2,5,4,8,3,7}
    ret := maxArea(nums)
    fmt.Println("start:",nums, " ret:",ret)
}

func maxArea(height []int) int {
    var i,j,res=0,len(height)-1,0
    for i<j{
        res=max(res,(j-i)*min(height[i],height[j]))
        if height[i]>height[j]{
            j--
        }else{
            i++
        }
    } 
    return res
}

func max(a,b int)int{
    if a>b{
                return a
                    
    }
        return b

}
func min(a,b int)int{
    if a>b{
                return b
                    
    }
        return a

}

