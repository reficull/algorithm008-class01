package main
import (
    "fmt" 
)
var output [][]int
var num int
var k int
var option []int
func main(){
    num = 4
    k = 2
    output = combine(num,k)

    printMatrix(output) 
}


func combine(n int, k int) [][]int  {
    if k == 0 {
        return [][]int{}

    }
    // path用来存储路径
    path := make([]int, 0)
    // ans用来存储最终结果集合
    ans := make([][]int, 0)
    dfs(n, &path, &ans, 1, k)
    return ans

}
func dfs(n int, path *[]int, ans *[][]int, index int, k int) {
    // 收集的长度够了 添加到结果集中
    if len(*path) == k {
        tmp := make([]int, len(*path))
        // 这里需要将path中的数据拷贝到新的slice中
        copy(tmp, *path)
        *ans = append(*ans, tmp)
        return

    }
    // 索引已经走到最后了 也需要return
    if index == n+1 {
        return

    }
    // 添加剪枝,当前长度+剩余可用的数不够k个不用在向下递归了
    if len(*path)+(n-index)+1 < k {
        return

    }
    // 不选当前index的数
    dfs(n, path, ans, index+1, k)

    // 选当前index的数
    *path = append(*path, index)
    dfs(n, path, ans, index+1, k)
    // 撤销本次选择
    *path = (*path)[:len(*path)-1]

}


func backtrack(first int,idx int){
    fmt.Printf("process : %d\n",first)
    for i:=first;i<= num+1 - (k-1+idx);i++{
        option[idx] = i
        if idx == k-1{
            output = append(output,option)
            fmt.Printf("append: %v\n",option)
        }else{
            backtrack(i+1,idx+1)

        }
    }
} 

func printMatrix(matrix [][]int){
    for r,col1 := range matrix{
        fmt.Printf("row %d:%v\n",r,col1)
    }
}
