package main
import (
    "fmt" 
)

func main(){
    n := 11
    ret := nthUglyNumber(n)
    fmt.Println(" n:",n," chou Num:",ret)
}

// 从i=2开始遍历自然数,直到找到丑数个数为n
func nthUglyNumber(n int) int {
    if n == 1 {
        return 1

    }
    i, j, k := 0, 0, 0
    result := []int{1}
    res := 0
    for n > 1 {
        a := result[i] * 2
        b := result[j] * 3
        c := result[k] * 5
        res = min(min(a, b), c)
        if res == a {
            i++

        }
        if res == b {
            j++

        }
        if res == c {
            k++

        }
        result = append(result, res)
        n--

    }
    return result[len(result)-1]

}
