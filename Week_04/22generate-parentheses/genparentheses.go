package main
import (
    "fmt" 
)
func main(){
    n := 3
    str := generateParenthesis(n)
    fmt.Println("ret:",str)
    
}

func generateParenthesis(n int) []string {
    str := make([]string,0)
    brackets := make([]byte,n*2)
    helper(brackets,0,0,n,&str) 
    return str
}

func helper(brackets []byte,il int ,ir int,n int,str *[]string) {
    if il == n && ir == n {
        *str = append(*str,string(brackets))
        return 
    } 
    if il < ir {
        return
    }
    if il < n{
        brackets[il+ir] = '('
        helper(brackets,il+1,ir,n,str)
    }
    if ir < n{
        brackets[il+ir] = ')'
        helper(brackets,il,ir+1,n,str)
    }


}
