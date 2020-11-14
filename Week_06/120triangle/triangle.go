package main
import (
    "fmt" 
)

func main(){
    triangle := [][]int{
        {2},
        {3,4},
        {6,5,7},
        {4,1,8,3},
    }

    printMatrix(triangle)
    ret := minimumTotal(triangle)
    printMatrix(triangle)
    fmt.Printf("ret:%d\n",ret)
}

func minimumTotal(triangle [][]int) int {
    row := len(triangle)
    for i:=row-2;i >=0;i--{
        for j:=0;j<len(triangle[i]);j++{
            triangle[i][j]=Min(triangle[i+1][j]+triangle[i][j],triangle[i+1][j+1]+triangle[i][j])
        }
    }
    return triangle[0][0]

}

func Min(a , b int) int{
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
