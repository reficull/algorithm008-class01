package main
import (
    "fmt" 
)

func main(){
    matrix := [][]int{
        {1,2,3},
        {4,5,6},
        {7,8,9},
    }

    ans := spiralOrder(matrix)
    fmt.Println("ans:",ans)
}
func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []int{}
    }
    res := []int{}
    u := 0
    d := len(matrix) - 1
    l := 0
    r := len(matrix[0]) - 1
    for{
        for i := l; i <= r; i++ {
            res = append(res,matrix[u][i])
        }
        u++
        if u > d{
            break
        }

        for i := u; i <= d; i++ {
            res = append(res,matrix[i][r])
        }
        r--
        if r < l{
            break
        }

        for i := r; i >= l ; i-- {
            res = append(res,matrix[d][i])
        }
        d--
        if d < u {
            break
        }

        for i := d; i>= u; i-- {
            res = append(res,matrix[i][l])
        }
        l++
        if l > r {
            break
        }
    }
    return res
}
