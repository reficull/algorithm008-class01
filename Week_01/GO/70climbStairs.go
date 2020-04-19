package main
import (
    "fmt" 
)

func main(){
    num := 5
    res := climbStairs(num)
    fmt.Printf("num:%d, has solu:%d\n",num,res)
}

func climbStairs(n int) int {
    if n <= 2 { return n }
    f1,f2,f3 := 1,2,3
    for i := 3;i<= n; i++{
        f3 = f2 + f1
        f1 = f2
        f2 = f3
    } 
    return f3
}
