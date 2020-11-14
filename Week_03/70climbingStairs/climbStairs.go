package main
import (
    "fmt" 
)

func main(){
    num := 6 
    res := climbStairs(num)
    fmt.Printf("num:%d, has solu:%d\n",num,res)
}

func climbStairs(n int) int {
    var dp [3]int
    dp[0]=1
    dp[1]=1
    for i:=2;i<=n;i++{
        dp[i%3]=dp[(i-1)%3]+dp[(i-2)%3]
        fmt.Printf("i:%d,dp[i//3]:%d = dp[(i-1)//3]:%d+dp[(i-2)//3]:%d\n",i,dp[i%3],dp[(i-1)%3],dp[(i-2)%3])
        fmt.Println("dp:",dp)
    }
    return dp[n%3]
}
