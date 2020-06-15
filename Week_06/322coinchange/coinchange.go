package main
import (
    "fmt" 
    "sort"
    "math"
)

func main(){
    coins := []int{1, 2, 5}
    amount := 11
    //coins := []int{ 2}
    //amount := 3
    ret := coinChange(coins,amount)
    fmt.Println("ret:",ret)
}
func coinChange(coins []int, amount int) int {
    minCount := math.MaxInt32
    sort.Ints(coins)
    dfs(coins, len(coins)-1, amount, 0, &minCount)
    if minCount == math.MaxInt32 {
        return -1

    }
    return minCount

}
// 硬币组合, 硬币索引, 剩余金额,次数,最小次数
func dfs(coins []int, index int, amount int, count int, minCount *int) {
    fmt.Printf("dfs:index:%d,amount:%d,count:%d,minCount:%d\n",index,amount,count,minCount)
    if index < 0 {
        return

    }
    //根据索引取一个硬币
    coin := coins[index]
    //根据币面值计算取币次数作为循环次数
    for cnt := amount / coin; cnt >= 0; cnt-- {
        fmt.Printf("in for cnt:%d\n",cnt)
        //剩余金额= 金额- 次数*硬币面值
        remainAmount := amount - cnt*coin
        //当前次数 = 传入次数+计算出的剩余次数
        currentCount := count + cnt
        //如果剩余金属为0, 则有结果,退出
        if remainAmount == 0 {
            //最小次数为 比较传入的最小次数,当前次数
            *minCount = Min(*minCount, currentCount)
            break

        }

        // remainAmount不为0, 至少还需要使用一枚硬币, 结果不会更优
        if currentCount+1 >= *minCount {
            break

        }
        //下一层,索引数减1, 剩余金额, 次数为传入+使用的次数,最小次数
        dfs(coins, index-1, remainAmount, currentCount, minCount)

    }

}
/* dynamic program
func coinChange(coins []int, amount int) int {

    dp := make([]int,amount+1)
    // f(n) = min{f(n-k),for k in [1,2,5]} + 1 
    for i:=1;i<=amount;i++{
        dp[i]=-1
        for _,c := range coins{
            if i < c ||dp[i-c]==-1{
                continue
            }
            count := dp[i-c] + 1                
            if dp[i] == -1||dp[i] > count{
                dp[i] = count
            }

            fmt.Printf("i-c:%d dp[i-c]:%d \n",i-c,dp[i-c] )
        }
    }
    fmt.Println(dp)
    return dp[amount]
}
*/
func Min(a,b int) int{
    if a < b{
        return a
    }
    return b
}
