package main

import(
    "fmt"
    "strings"
    "regexp"
)

func main(){
    words := "A man, a plan, a canal: Panama"
    ret:= isPalindrome(words)
    fmt.Println("ret:",ret)
}


func isPalindrome(s string) bool {
    if len(s) < 2{
        return true
    }
    sl := strings.ToLower(s) 
    reg, _:= regexp.Compile("[^a-zA-Z0-9]+")
    processedString := reg.ReplaceAllString(sl, "")
    fmt.Println("s:",processedString)
   
    return true
}
