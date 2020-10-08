package main

import(
    "fmt"
)

func main(){
    num := 2.000
    n := -2 
    ret := myPow(num,n)
    fmt.Println("ret:",ret)
}


func myPow(x float64, n int) float64 {
    fmt.Printf("cal %d pow %d\n",x,n)
    if n == 0{
        return 1 
    }
    if n == 1{
        return x
    }
    var ret float64
    if n < 0{
        ret = 1 / myPow(x,-n)
    }
    if n == 2{
        return x * x
    }else if n> 2{
        if n &   1  == 0{
            tmp := myPow(x,n/2)    
            ret = tmp * tmp
        }else {
            tmp := myPow(x,(n-1)/2)
            ret = tmp * tmp * x
        }

    }

    return ret

}
