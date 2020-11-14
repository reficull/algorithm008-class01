package main
import (
    "fmt" 
    "sort"
)

const MaxUint = ^uint(0) 
const MinUint = 0 
const MaxInt = int(MaxUint >> 1) 
const MinInt = -MaxInt - 1

func main(){
//    nums1 := []int{2,1,-2,5}
//    nums2 := []int{3,0,-6}
    nums1 := []int{-1,-1}
    nums2 := []int{1,1}
    fmt.Println("nums1:%v",nums1)
    fmt.Println("nums2:%v",nums2)
    ret := maxDotProduct(nums1,nums2)
    fmt.Printf("ret:%d\n",ret)
}
/*
    
*/
func maxDotProduct(nums1 []int, nums2 []int) int{
    l1 := len(nums1)
    l2 := len(nums2)
    dp := make([][]int,2)
    dp[0] = make([]int,l2+1)
    dp[1] = make([]int,l2+1)
    for i:=0;i<l2;i++{
        dp[1][i] =  MinInt
    }
    //i&1               当前行
    //(i+1)&1  preID    另外一行
    for i := 0; i < l1; i++ {
        preID := (i + 1) & 1
        dp[i&1][0] =  MinInt 

        for j := 1; j <= l2; j++ {

            preValue := dp[preID][j-1] //另一行的前一位
            if preValue < 0 {
                preValue = 0

            }

            // if i == 0 {
            //  dp[i&1][j] = Max(dp[i&1][j-1], dp[preID][j-1]+nums1[i]*nums2[j-1])

        //} else {
        tmp := Max(dp[preID][j], dp[i&1][j-1]) //另一行的当前值  与当前行的前一位
            dp[i&1][j] = Max(tmp, preValue+nums1[i]*nums2[j-1])
            fmt.Printf("i:%d,(i+1)&1:%d,i&1:%d,j:%d,%d*%d+%d max %d =dp[i&1][j]:%d\n",i,preID,i&1,j,nums1[i],nums2[j-1],preValue,tmp,dp[i&1][j])
            printMatrix(dp)
        }
    }

    return dp[(l1-1)&1][l2]
}

func printMatrix(matrix [][]int){
    for r,col1 := range matrix{
        fmt.Printf("row %d:%v\n",r,col1)
    }
}
/*
    l1 := len(nums1)
    l2 := len(nums2)
    dp := make([][]int,0)
    for i:=0;i<l1+1;i++{
        row := make([]int,0)
        for j:=0;j<l2+1;j++{
            row = append(row,MinInt)
        }
        dp=append(dp,row)
    }

    for i:=1;i<l1;i++{
        for j:=1;j<l2;j++{
            cur := nums1[i]*nums2[j]
            tmp := Max(0,dp[i][j]) + cur 
            fmt.Printf("cur:%d,num1:%d,num2:%d,tmp:%d\n",cur,nums1[i],nums2[j],tmp)
            dp[i+1][j+1]=Max(tmp, Max(dp[i+1][j],dp[i][j+1]))
            fmt.Printf("[i+1][j+1]:%d\n",dp[i+1][j+1])
        }
    }
    printMatrix(dp)
    return dp[l1][l2]

*/

func maxDotProduct1(nums1 []int, nums2 []int) int {
    l1 := len(nums1)
    l2 := len(nums2)
    maxLen := Max(l1,l2)
    fmt.Println("max len:%v",maxLen)
    dotMap := make(map[int]int,0)
    dot := make( []int,maxLen)
    fmt.Println("dot:%v",dot)
    for i:=0;i<maxLen-1;i++{
        dot[i] = nums1[i] * nums2[i]
        dotMap[dot[i]] = i
    }
    sort.Ints(dot)
    if dot[maxLen-1]  < 0{
        return  dotMap[dot[maxLen-1]]
    }
    var retAry []int
    ret := 0 
    for i:=maxLen-1;dot[i]>0;i--{
        retAry = append([]int{dot[i]},retAry...)
        ret = ret + dot[i] 
    }
    fmt.Println("ret ary:%v",retAry)
    fmt.Println("dot map:%v",dotMap)
    return ret
}
func Max(a ,b int) int{
    if a > b{
        return a
    }
    return b
}

func Min(a ,b int) int{
    if a < b{
        return a
    }
    return b
}
