package main
import (
    "fmt" 
)


func main(){
   // nums := []int{2,1,1,2}
    //nums := []int{0}
    nums := []int{1,2,3,1,3,2}
    fmt.Println(nums)
    ret := rob(nums)
    fmt.Printf("ret:%d\n",ret)
}

func rob(nums []int) int {
    l := len(nums)
    if l == 0{
        return 0 
    }
    if l == 1{
        return  nums[0] 
    }
    dp := make([]int,l)
    dp[0]=nums[0]
    dp[1]=Max(nums[0],nums[1])
 //   maxVal :=dp[1] 
    for i:=2;i<l;i++{
        dp[i] =  Max( dp[i-1],dp[i-2]+nums[i])
  //      maxVal = Max(maxVal,dp[i])
    }
    fmt.Println(dp)
    return Max(dp[l-1],dp[l-2])
}

func Max(a, b int) int{
    if a > b{
        return a
    }
    return b
}

func printMatrix(matrix [][]int){
    for r,col1 := range matrix{
        fmt.Printf("row %d:%v\n",r,col1)
    }
}
    /*
    f(1) = nums[1]
    f(2) = nums[2]
    f(3) = Max(nums[1] + f(1),f(2))

    f(n) = Max(nums[n]+f(n-2),f(n-1))
    dp[i][j]= Max(nums[i] + dp[i-2][j],dp[i-1][j-1])
    
    steal:
       dp[i][1]=  nums[i] + Max(dp[i-1][0],dp[i-1][1])
    not steal:
       dp[i][0] = Max(dp[i-1][0],dp[i-1][1])

    */
