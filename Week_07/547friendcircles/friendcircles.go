package main
import (
    "fmt" 
)
var count int
func main(){
    count = 0
    f := [][]int{
        //{1,1,0},
        //{1,1,0},
        //{0,0,1},
        {1,1,0},
        {1,1,1},
        {0,1,1},
    }
    c := findCircleNum(f)
    fmt.Println("ret:",c)
}

func findCircleNum(M [][]int) int {
    l := len(M)
    visited := make([]bool,l)
    for i:=0;i<l;i++{ visited[i]=false }
    for i:=0;i<l;i++{
        if visited[i] == false {
            dfs(M,visited,i)
            count++
            
        }
    }

    return count
}

func dfs(m [][]int,visited []bool,i int){
    for j:=0;j<len(m);j++{
        if m[i][j] == 1 && !visited[j]{
            visited[j] = true
            dfs(m,visited,j)
        }
    }
}


