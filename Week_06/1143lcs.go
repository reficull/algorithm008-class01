package main
import (
    "fmt" 
)

func main(){
    str1 := "abcde"
    str2 := "ace"
    ret := longestCommonSubsequence(str1,str2) 
    fmt.Printf("result:%d\n",ret)
}

func longestCommonSubsequence(text1 string, text2 string) int {
    dp := make([][]int,0)
    l1 := len(text1)
    l2 := len(text2)
    for i :=0;i<l2+1;i++{
        row := make([]int,0)
        for j := 0;j < l1+1;j++{
            row = append(row,0)
        }
        dp = append(dp,row)
    }
    for i := 1;i<l2+1;i++{
        for j := 1;j<l1+1;j++{
            if text1[j-1] == text2[i-1]{
                dp[i][j]=1 + dp[i-1][j-1]
            }else{
                dp[i][j]=Max(dp[i-1][j],dp[i][j-1])
            }
        }
    }

    printMatrix(dp)

    return dp[l2][l1] 
}
func Max(a,b int) int{
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
