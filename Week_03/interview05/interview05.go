package main
import (
    "fmt" 
)

func main(){
    s := "We are happy."
    ans := replaceSpace(s)
    fmt.Println("ans:",ans)
}

func replaceSpace(s string) string {
    ans := []rune{}
    for i:=0;i<len(s);i++{
        if s[i] != ' ' {
            ans = append(ans,int32(s[i]))
        }else{
            ans = append(ans,'%')
            ans = append(ans,'2')
            ans = append(ans,'0')
        }
    }
    return string(ans)

}
