package main

import(
    "fmt"
)
func main(){
    num := 0 
    ret := isPowerOfTwo(num)
    fmt.Println(num, " is 2^x :",ret)
}


func isPowerOfTwo(n int) bool {
    if n == 0 {return false}
    if  n & (-n) != n{
        return false
    }
    return true

}
