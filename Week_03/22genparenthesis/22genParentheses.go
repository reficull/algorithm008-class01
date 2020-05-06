package main
import (
    "fmt" 
)

var result string

func main(){
    n := 4 
    str := generateParentheses(n)
    fmt.Println("n:",n," ret:",str)
}

func generateParentheses(n int) string{
    helper("",n,n)
    return result
}

func helper(str string,left int,right int) {
    if left == 0 && right == 0{
        result += str + "\n"
    }
    if left != 0 {
        helper(str + "(",left-1,right)
    }
    if right != 0 && right > left {
        helper(str + ")",left,right-1)
    }
}

