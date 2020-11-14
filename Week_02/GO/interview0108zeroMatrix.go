package main
import (
    "fmt" 
)

func main(){
    /*
    matrix := [][]int{
        {1,1,1},
        {1,0,1},
        {1,1,1},
    }
    */
    matrix := [][]int{
        {0,1,2,0},
        {3,4,5,2},
        {1,3,1,5},
    }

    printMatrix(matrix)
    setZeroes(matrix)
    printMatrix(matrix)
}


func setZeroes(matrix [][]int)  {
    record := make([][]int,0)

    for i:=0;i<len(matrix);i++{
        for j:=0;j<len(matrix[i]);j++{
            if matrix[i][j]==0{
                record = append(record,[]int{i,j})
            }
        }
    }
    fmt.Println("record:",record)
    for _,cr := range record{
        zeroCol(matrix,cr[1])
        zeroRow(matrix,cr[0])
    }

}

func zeroCol(matrix [][]int,c int){
    for i:=0; i<len(matrix);i++{
        if c > len(matrix[i])-1{
            return
        }
        matrix[i][c] = 0
    }
}

func zeroRow(matrix [][]int,r int){
    if r > len(matrix)-1 { return }
    for i:=0;i<len(matrix[r]);i++{
        matrix[r][i] = 0
    }
}
func printMatrix(matrix [][]int){
    for r,col1 := range matrix{
        fmt.Printf("row %d:%v\n",r,col1)
    }
}

