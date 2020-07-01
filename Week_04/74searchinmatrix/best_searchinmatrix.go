package main
import (
    "fmt" 
    "math"
)

func main(){
    matrix := [][]int{
        {-9,-9,-9,-8,-8,-7,-6,-4,-4,-3},
        {0,1,2,2,4,5,5,5,7,9},
        {12,12,14,14,15,17,19,19,19,21},
        {22,23,23,25,25,26,26,28,28,29},
        {31,31,31,33,34,36,37,38,38,39},
        {40,42,43,44,46,46,46,47,49,50},
        {52,54,55,57,59,60,60,62,64,66},
        {68,68,70,71,71,72,74,76,78,80},
        {82,83,85,85,85,87,88,88,89,89},
        {90,90,90,91,93,94,94,95,95,97},
        {98,98,99,99,101,103,105,106,108,109},
        {112,112,112,113,113,113,114,116,118,118},
        {119,121,122,124,125,125,125,126,127,128},
        {131,133,134,134,134,135,135,137,137,139},
        {141,143,145,147,148,150,150,150,150,152},
        {153,153,154,155,157,157,157,159,161,162},
        {164,166,167,167,167,169,170,170,171,173},
        {176,176,178,179,181,182,183,183,184,186},

        //{1},
        //{3},
        //{1,   3,  5,  7},
        //{10, 11, 16, 20},
        //{23, 30, 34, 50},   
        //{51, 55, 60, 65},   
    }

    ret := searchMatrix(matrix,135)
    fmt.Println("ret:",ret)
}

func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0{
        return false

    }
    l,r,mid := 0,len(matrix) * len(matrix[0]),0
    for l <= r {
        mid = l + (r - l) >> 1
        if mid == 0 {
            l = 1
            continue

        }
        tmp := math.Ceil(float64(mid)/float64(len(matrix[0]))) - 1
        row := int(tmp)
        col := 0
        if col = (mid % len(matrix[0])) - 1; col == -1 {
            col = len(matrix[0]) - 1

        }
        midV := matrix[row][col]
        if target == midV {
            return true

        } else if target > midV {
            l = mid + 1

        } else {
            r = mid - 1

        }


    }
    return false
}