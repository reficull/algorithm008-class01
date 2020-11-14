package main

import(
    "fmt"
    "container/list"
    "strings"
)


func main(){
   // s := "2[abc]3[cd]ef"
    //s := "3[a]2[bc]"
    s := "3[a2[c]]"
    //s := "100[leetcode]"
    fmt.Println("start:",s)
    ret := decodeString(s)
    fmt.Println("ret:",ret)
}

func decodeString(s string) string {
    type RepStr struct{
        Rep int
        Str string
    }
    stack := list.New()
    tmpS := ""
    var tmpN int = 0 
    for i:=0;i<len(s);i++{
        nOrs := s[i]
        if nOrs >= '0' && nOrs <= '9'{
            fmt.Printf("tmpN num:%d\n",tmpN)
            tmpN = tmpN * 10 + int(nOrs - '0')
            fmt.Printf("get num:%d\n",tmpN)
            //tmpN,_ = strconv.Atoi(string(s[i]))
        }else if nOrs == '['{
            stack.PushBack(RepStr{tmpN,tmpS})
    //        fmt.Printf("push num:%d,str:%s\n",tmpN,tmpS)
            tmpN,tmpS  = 0,""
        }else if nOrs == ']'{
            temp := stack.Back().Value.(RepStr)
            stack.Remove(stack.Back())
            tmpS = temp.Str + strings.Repeat(tmpS,temp.Rep)
     //       fmt.Printf("tmpS:%s,repeat time:%d\n")
        }else{
            tmpS += string(nOrs)
        } 
    }
    
    //fmt.Printf("stack:%v,len:%d,tmpS:%s\n",stack,stack.Len(),tmpS)
    return tmpS 
}
