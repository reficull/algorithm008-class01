
package main
import (
    "fmt" 
    "math"
     //"gonum.org/v1/gonum/mat"
)
/*
给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。

两个相邻元素间的距离为 1 。
https://leetcode-cn.com/problems/01-matrix/
*/
func updateMatrix(matrix [][]int) [][]int {
    //fmt.Println(matrix)
    for row,col1 := range matrix{
        fmt.Printf("check row:%d, col:%v\n",row,col1)
        //check around
        //check up
        for c := range col1{
            minV := checkVertical(matrix,row,c)
            minH := checkHorizon(matrix,row,c)
            var min = float64(0)
            if minV==-1{
                min = float64(minH)
            }else if minH == -1{
                min = float64(minV)
            }else{
                min = math.Min(float64(minV),float64(minH))
            }
            if min >=2 {
                fmt.Printf("got row:%d,col:%d to zero distance >=2,minV:%d,minH:%d\n",row,c,minV,minH)
                matrix[row][c] += (int(min) -1)
            }
        }
    }
    return matrix
}
func checkHorizon(matrix [][]int,r int,c int) int{
    minRZ := float64(-1);
    minLeft := float64(-1);
    minRight:= float64(-1);
    if matrix[r][c] != 0 {
        if c > 0{
            startLeft := c - 1;
            //fmt.Println("start check horizon row left:",startLeft," cur val:",matrix[r][c])
            for ci := startLeft;ci>=0;ci--{
                if matrix[r][ci] == 0{
                    minLeft = float64(c - ci)
                    //fmt.Println("check left: col:",c," row:",r," val:",matrix[r][c]," r:",r,", compare to col:",ci," val:" ,matrix[r][ci],",minLeft:",minLeft)
                    break;
                } 
            }
        }
        fmt.Println(" col num:",len(matrix[r]))
        if c < len(matrix[r])-1 {
            startRight := c + 1;
            //fmt.Println("start check horizon row right:",startRight," cur val:",matrix[r][c]," first right:",matrix[r][startRight])
            for ci := startRight;ci<len(matrix[r])-1 ;ci++{
                if matrix[r][ci] == 0{
                    minRight = float64(ci - c) 
                    //fmt.Println("check right: col:",c," row:",r," val:",matrix[r][c]," r:",r,", compare to col:",ci," val:" ,matrix[r][ci],",minRight:",minRight)
                    break;
                }
            }
        }
        if minLeft == -1{
            minRZ = minRight;
        }else if minRight == -1{
            minRZ = minLeft;
        }else{
            minRZ = math.Min(minLeft,minRight)
        }
        fmt.Println("minLeft:",minLeft,", minRight:",minRight," min:",minRZ)
    }
    return int(minRZ);
}

func checkVertical(matrix [][]int,r int,c int) int{
    minRZ := float64(-1);
    minUp := float64(-1);
    minDown := float64(-1);
    //最上面的行返回
    //if r == 0 { return int(minRZ) }
    if matrix[r][c] != 0 {
        startUp := r - 1;
        //fmt.Println("start check vertical row up:",startUp," cur val:",matrix[r][c])
        for ri := startUp;ri >=0; ri--{
            if matrix[ri][c] == 0 {
                minUp = float64(r - ri);
                //fmt.Println("check up: row:",r," col:",c," val:",matrix[r][c]," r:",r,", compare to row:",ri," val:" ,matrix[ri][c],",minUp:",minUp)
                break;
            }
        }
        startDown := r + 1;
        //fmt.Println("start check vertical row down:",startDown," cur val:",matrix[r][c])
        for ri := startDown;ri <len(matrix); ri++{
            if matrix[ri][c] == 0 {
                minDown = float64(ri - r);
                //fmt.Println("check down: row:",r," col:",c," val:",matrix[r][c]," r:",r,", compare to row:",ri," val:" ,matrix[ri][c],",minDown:",minDown)
                break;
            }
        }
        if minUp == -1 {
            minRZ = minDown
        }else if minDown == -1{
            minRZ = minUp 
        }else{
            minRZ = math.Min(minUp,minDown);

        }
        fmt.Println("minup:",minUp,", minDown:",minDown," min:",minRZ)
    }
    return int(minRZ);
}

func printMatrix(matrix [][]int){
    for r,col1 := range matrix{
        fmt.Printf("row %d:%v\n",r,col1)
    }
}

func main(){
    matrix := [][]int{
        {0,1,0,1,1},
        {1,1,0,0,1},
        {0,0,0,1,0},
        {1,0,1,1,1},
        {1,0,0,1,1},

    }
    /*
    matrix := [][]int{
        {0,0,0},
        {0,1,0},
        {1,1,1},
    }
    */
    var matrixOrigin  = append([][]int(nil),matrix...)

    printMatrix(matrix)
    ret := updateMatrix(matrix)
    printMatrix(matrixOrigin)    
    fmt.Println("=================ret")
    printMatrix(ret)

}
