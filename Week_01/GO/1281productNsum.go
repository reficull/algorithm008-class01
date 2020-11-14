package main
import (
    "fmt"
    "strconv"
) 

func subtractProductAndSum(n int) int {
    product := 1
    sum := 0
    str := strconv.Itoa(n)
    for _,s := range str{
        c := string(s)
        if i,err := strconv.Atoi(c);err==nil{
            product *= i 
            sum += i
        }
    }
    return product - sum
}

func main(){
    i := 234
    ret := subtractProductAndSum(i)
    fmt.Println("result:",ret)
}
