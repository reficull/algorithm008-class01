package main
import (
    "fmt" 
    "math"
)

func sortedSquares(A []int) []int {
    /*
    for i,n := range A{
        A[i] = n * n
    }
    sort.Ints(A)
    return A
    */
    //negtive num pointer 负数指针
    var ans = make( []int,len(A))
    left := 0
    right := len(A)-1
    end := len(A)-1
    for ;left<=right; {
        //lSquare := A[left] * A[left]
        //rSquare := A[right] * A[right]
        if math.Abs(float64(A[left])) > math.Abs(float64(A[right])){
            fmt.Printf("abs[%d] < abs[%d]",A[left],A[right])
            ans[end] = A[left] * A[left]
            left++
        }else{
            fmt.Printf("abs[%d] > abs[%d]",A[left],A[right])
            ans[end] = A[right] * A[right]
            right--
        }
        fmt.Println("ans:",ans)
        end--
    }
    return ans 
}
func main(){
    a := []int{-7,-3,2,3,11}

    fmt.Println("start:",a)
    a = sortedSquares(a)
    fmt.Println("res:",a)
}


