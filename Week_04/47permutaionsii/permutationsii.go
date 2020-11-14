package main
import (
    "fmt" 
)

func main(){
    nums := []int{1,1,2}
    ret := permuteUnique(nums)
    printMatrix(ret)    
}

func permuteUnique(nums []int) [][]int {
    ret := make([][]int)
    for i:=0;i<len(nums);i++{

    }
     
}

func makePermute(nums []int,i) []int{

}

func printMatrix(matrix [][]int){
    for r,col1 := range matrix{
        fmt.Printf("row %d:%v\n",r,col1)
    }
