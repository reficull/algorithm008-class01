package main
import (
    "fmt" 
)

const MaxUint = ^uint(0) 
const MinUint = 0 
const MaxInt = int(MaxUint >> 1) 
const MinInt = -MaxInt - 1

func main(){
    word1 := "horse"
    word2 := "ros"
    ret := minDistance(word1,word2)
    fmt.Printf("ret:%d\n",ret)
}

func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    var dp = make([][]int, m+1)
    for i := 0; i <= m; i++ {
        dp[i] = make([]int, n+1)

    }
    // 空字符串编辑为空字符串的编辑代价为0
    dp[0][0] = 0
    // dp[i][0]表示编辑为空串的代价，即为将所有字符删除的代价
    for i := 1; i <= m; i++ {
        dp[i][0] = i 

    }
    // dp[0][i]表示将空串编辑为str2[:i]的代价，即插入字符的代价
    for i := 1; i <= n; i++ {
        dp[0][i] = i

    }
    // 下面是动态规划的主方法
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]              // 如果 word1[i-1] 与 word2[i-1]相等

            } else {
                dp[i][j] = dp[i-1][j-1] + 1          // 替换代价

            }
            dp[i][j] = Min(dp[i][j], dp[i][j-1] + 1) // 插入代价
            dp[i][j] = Min(dp[i][j], dp[i-1][j] + 1) // 删除代价

        }

    }
    printMatrix(dp)
    return dp[m][n]


}

func Min(a,b int) int{
    if a < b{
        return a
    }
    return b
}

func printMatrix(matrix [][]int){
    for r,col1 := range matrix{
        fmt.Printf("row %d:%v\n",r,col1)
    }
}
