package main

import(
    "fmt"
    "strings"
)
var visited [][]int
var ans [][][]int
func main(){
    n :=4 
    ret := solveNQueens(n)
    fmt.Println(ret)
}

func makeAns(nums [][]int) []string{
    r := len(nums)
    ret := make([]string,0)
    for i:=0;i<r;i++{
        row := ""
       // if nums[i][1]==0{
        //    row = row+"Q"+strings.Repeat(".",r-2)
        //}else{
            row = row+strings.Repeat(".",nums[i][1])+"Q"+strings.Repeat(".",r-nums[i][1]-1)
       // }
        ret = append(ret,row)
        //fmt.Println(row)
    } 
    return ret
}  

func solveNQueens(n int) [][]string {
    row,cols,pie,la := make(map[int]bool,0),make(map[int]bool,0),make(map[int]bool,0),make(map[int]bool,0)
    //visited = make([][]int,n)
    ans = make([][][]int,0)
    queens := make([][]int,0)
    for i:=0;i<n;i++{
     //   visited[i] = make([]int,n)
        //ans[i] = make([]int,2)
        for j:=0;j<n;j++{
      //      visited[i][j]=0
            pie[i+j]=false
            la[i-j]=false
        }
        row[i]=false
        cols[i]=false
    }
    dfs(queens,n,0,pie,la)
    r := make([][]string,0)
    for i:=0;i<len( ans);i++{
        r1 := makeAns(ans[i])
        r = append(r,r1)
    }

    return r
}

func dfs(queens [][]int,n,i int,pie ,la map[int]bool) bool{
    //fmt.Printf("i:%d\n",i)
    if i == n{
        ans1 := make([][]int,0)
        ans1 = append(ans1,queens...)
        ans = append(ans,ans1)
        fmt.Printf("get ans i:%d, queens:%v\n",i,queens)
        return  true
    }
    for j:=0;j<n;j++{
        if !numInAry(j,queens)&& !pie[i+j]&&!la[i-j]{
        //    fmt.Printf("in test:%d,%d\n",i,j)
            pie[j+i]=true
            la[i-j]=true
            queens = append(queens,[]int{i,j})
            dfs(queens,n,i+1,pie,la)
            pie[j+i]=false
            la[i-j]=false
            queens = queens[:len(queens)-1]

        } 
    }
    //if row[r] ||  
    return false

}

func Max(a ,b int) int{
    if a >b {
        return a
    }
    return b
}

func numInAry(j int,nums [][]int) bool{
    if len(nums) == 0 {return false}
    for _,n := range nums{
        if n[1] == j{
            return true
        }
    }
    return false
}


func printMatrix(matrix [][]int){
    for r,col1 := range matrix{
        fmt.Printf("row %d:%v\n",r,col1)
    }
}
