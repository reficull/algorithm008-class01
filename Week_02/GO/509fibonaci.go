package main
import (
    "fmt" 
)
func main(){
    n := 6
    ret := fib(n)
    fmt.Println("n:",n, " fib:",ret)
}

func fib(N int) int {
    if N == 0{
        return 0
    }
    if N==1{
        return 1
    }
    if N < 3{
        return N-1
    }

    n1, n2 := 1, 1  // n1为n-1，n2为n-2
    for i := 3; i < N; i++ {
        n1, n2 = n1+n2, n1

    }
    return n1 + n2

    /*best
    var goldenRatio float64 = float64((1 + math.Sqrt(5)) / 2);
    return int(math.Round(math.Pow(goldenRatio, float64(N)) / math.Sqrt(5)));
    */
}

