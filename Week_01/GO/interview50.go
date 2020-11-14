package main
import (
    "fmt" 
)

func main(){
    str := string("")
    //str := string("leetcode")
    //str := string("abaccdeff")
    ret := firstUniqChar(str)
    fmt.Println("start:",str, " ret:",string(ret))
}

func firstUniqChar(s string) byte {
    ss := []rune(s)
    smap := make(map[rune]int)
    ret := byte(32) 
    for _,c := range ss{
        if _,ok := smap[c];ok{
            fmt.Println("check char found:",string(c))
            smap[c]++
        }else{
            fmt.Println("check char not found:",string(c))
            smap[c]=1
        }
    }
    for k,v := range ss{
        fmt.Println("check map:",string(k))
        if smap[v] == 1{
            ret = byte(ss[k])
            break
        }
    }
    fmt.Println("smap:",smap, " found first 1:",string(ret))
    return ret
}
