package main
import (
    "fmt" 
    "strconv"
)

func main(){
    n := 15
    str := fizzBuzz(n)
    fmt.Printf("start:%d\n fizzbuzz:%v ",n,str)
}
func fizzBuzz(n int) []string {
    ret := make([]string,n)
    for i:=1;i<=n;i++{
        
        if i % 15 == 0{
            ret[i-1] = "FizzBuzz"
        }else if i % 3 == 0{
            ret[i-1] = "Fizz"
        }else if i % 5 == 0{
            ret[i-1] = "Buzz"
        }else{
            ret[i-1] = strconv.Itoa(i) 
        }
    }
    return ret
}
