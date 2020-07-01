package main

import(
    "fmt"
)


func main(){
   // pa := "()"
    //pa := ""
    //pa := "]"
    //pa := "}{"
    //pa := "[])"
    pa := "()[]{}"
    //pa := "()[]{}{"
    ret := isValid(pa)
    fmt.Printf("ret:%v\n",ret)
}

func isValid(s string) bool {
    if len(s) ==1 {return false}
    if len(s) ==0 {return true}
    stack := make([]rune,0)
    for i:=0;i<len(s);i++{
        c := rune(s[i])
        if c == '('{
            stack = append(stack,rune(')'))
        }else if c == '['{
            stack = append(stack,rune(']'))
        }else if c == '{'{
            stack = append(stack,rune('}'))
        }else{
            fmt.Println("stack len:%d",len(stack))
            if len(stack)>0{
                char := stack[len(stack)-1]
                fmt.Printf("string:%v,c:%v,char:%v\n",string(stack),c,char)
                if c != char{
                    return false
                }
                fmt.Println("stack len:%d",len(stack))
                stack = stack[:len(stack)-1]

            }else{
                return false
            }

        }

    }

    if len(stack) != 0{
        return false
    }
    return true

}
