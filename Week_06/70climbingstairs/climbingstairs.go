package main
import (
    "fmt" 
)

func main(){
    num := 4 
    ret := climbStairs(num)
    fmt.Printf("ans:%d\n",ret)
}
func climbStairs(n int) int {
    if n == 1{
        return 1
    }
    if n == 2{
        return 2
    }
    dp := [3]int{0}
    dp[0]=1
    dp[1]=2
    for i:= 2;i<=n;i++{
        dp[i%3] = dp[(i-1)%3] + dp[(i-2)%3]
    }
    return dp[(n-1)%3]
}
