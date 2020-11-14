package main
import (
    "fmt" 
)


func main(){
    //str := string("Test1ng-Leet=code-Q!")
    str := string("z<*zj")
    //str := string("ab-cd")
   // str := string("a-bC-dEf-ghIj")
    fmt.Println("start:",str)
    str = reverseOnlyLetters(str)
    fmt.Println("result:",str)
}

func reverseOnlyLetters(S string) string {
    /*
    ss := []byte(S)
    lFlag := false 
    rFlag := false 
    for l,r := 0,len(ss)-1;l < r && r > 0;{
        for{
            if isLetter(string(ss[r])){
                rFlag = true
                break;
            }else{
                if r>0 {r--}else{break}
            } 
        }
        for {
            if isLetter(string(ss[l])){
                lFlag = true
                break
            }else{
                if l<r {l++}else{break}
            }
        }
        if l>=r {break}
        if lFlag && rFlag {
            fmt.Println("do swap:",string(ss[l]),string(ss[r]))
            ss[l] = ss[r];
           // tmp := ss[l];ss[r] = tmp

            fmt.Println("do swap after:",string(ss[l]),string(ss[r]), " str:",string(ss))

            if r>0 {r--}
            if l<r {l++}
        }
    }
    return string(ss)
    */

    //best solution
    ss := []rune(S)
    n := len(S)
    i,j := 0,n-1
    for i < j{
        if !isLetter(ss[i]){
            i++
            continue
        }
        if !isLetter(ss[j]){
            j--
            continue
        }
        ss[i],ss[j] = ss[j],ss[i]
        i++
        j--
    }
    return string(ss)

}

func isLetter(s rune) bool{
    r := string(s)
    if(r < string('a') || r > string('z')) && (r < string('A') || r > string('Z')) {
        fmt.Println(s , " is not letter")
        return false
    }else{
        fmt.Println(s , " is letter")
        return true
    }
}
