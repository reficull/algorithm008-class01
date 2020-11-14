package main
import (
    "fmt" 
)

func main(){
    str := string("1123")
    guess := string("0111")
    //str := string("1807")
    //guess := string("7810")
    fmt.Printf("start secret: %v,guess:%v\n",str,guess)
    ret := getHint(str,guess)
    fmt.Println("secret:",str," guess:",guess," hint:",ret)
}


func getHint(secret string, guess string) string {
//best NO.1
    m := make([]int, 128)
    l := len(secret)
    for i := 0; i < l; i++ {
        m[secret[i]]++

    }
    fmt.Println("count array:",m)
    aN, bN := 0, 0
    for i := 0; i < l; i++ {
        n := m[guess[i]]
        fmt.Println("n:",n)
        if secret[i] == guess[i] {
            if n==0{
                bN--
                fmt.Printf("guess right,n==0,bN--:%d\n",bN)
            }else{
                m[guess[i]]--
                fmt.Printf("guess right,n=%d,%s count --:%d\n",n,string(guess[i]),m[guess[i]])
            }
            aN++
        } else {
            if n > 0 {
                bN++
                m[guess[i]]--

                fmt.Printf("guess wrong,n=%d,%s count --:%d , bN:%d\n",n,string(guess[i]),m[guess[i]],bN)
            }

        }

    }
    return  fmt.Sprintf("%dA%dB", aN, bN)
//    return strconv.Itoa(aN) + "A" + strconv.Itoa(bN) + "B"

//best NO.2
/*
var mmS = [10]int{}
var mmG = [10]int{}
var sBytes = []byte(secret)
var gBytes = []byte(guess)
var bulls,cows = 0,0
for i:=0;i<len(secret);i++{
    if sBytes[i] == gBytes[i]{
        bulls++

    }else{
        mmS[sBytes[i]-'0']++
        mmG[gBytes[i]-'0']++

    }

}
for i:=0;i<10;i++{
    cows += min(mmG[i],mmS[i])

}
var ret = fmt.Sprintf("%dA%dB",bulls,cows)
return ret

*/


    /*
    bulls,cows := 0,0
    numbers := [10]int{0}
    for i:=0;i<len(secret);i++{
        if secret[i] == guess[i] {
            bulls++
        }else{
            numbers[secret[i] - '0']++
            if numbers[secret[i] - '0'] < 0 {cows++}
            numbers[guess[i] - '0']--
            if numbers[guess[i] - '0'] < 0 {cows++}

            fmt.Printf("numbers:%v,i:%d\n",numbers,i)
        }
    }
    fmt.Println("bulls:",bulls," cows:",cows)
    return fmt.Sprintf("%dA%dB",bulls,cows)
   //================= 
    restMap := make(map[rune]int,0)
    guessStr := make([]rune,len(guess)) 
    A ,B:= 0,0

    for i,v := range secret{
        if secret[i] == guess[i]{
            A++
        }else{
            if _,ok := restMap[v];ok{
                restMap[v]++
            }else{
                restMap[v]=1
            }
            guessStr = append(guessStr , rune(guess[i]))
            fmt.Println("guss str add:",string(v))
        }
    } 
    
    fmt.Println("guess wrong:",guessStr, " A:",A, " rest map:",restMap)
    for _,v := range guessStr{
        if j,ok := restMap[v];ok{
            B++
            if j <= 1{
                fmt.Println("loop rest str left 1:",v, "num:",j)
                delete(restMap,v)
            }else{
                fmt.Println("loop rest str more then 1:",v, "num:",j)
                restMap[v]--
            }
        }
    }
    ret := fmt.Sprintf("%dA%dB",A,B)
    return ret
    */
}

func min(a,b int)int{
    if a < b{
                return a
                    
    }
        return b

}
