package main
import (
    "fmt" 
)
/*
    1. 记录初始索引, 遇到( 则加入)到数组,遇到)则删除最后一个),如果删除到空,则记录这个)的位置,
    2. 重复以上过程直到结速
*/
func main(){
    //str := "(()())(())"
    str := "(()())(())(()(()))"
    ret := removeOuterParentheses(str)
    fmt.Printf("start:%s, ret:%s\n",str,ret)
}

func removeOuterParentheses(S string) string {
    stackState := make([]int,len(S)+1)                                                                                                           
    stackState[0]=0                                                                                                                                                                                                                      
    stackHigh := 0                                                                                                                                                                                                                          
    var newS []rune                                                                                                                    
    for i:=0;i<len(S);i++{                                                                                                                                                                                                               
        if S[i] == '('{                                                                                                                                                                                                                  
            stackHigh++                                                                                                                                                                                                                
        }else if S[i]==')'{                                                                                                                                                                                                              
            stackHigh--                                                                                                               
        }                                                                                                                                                                                                                                
        stackState[i+1] = stackHigh                                                                                                                                                                                                     
        // fmt.Println(" stack:",stack)              
        if stackHigh!=0 &&  !(stackHigh==1 && stackState[i]==0){                                                                                                                                                                       
            newS = append(newS,rune(S[i]))                 
        }                                                                                                                                                                                                                                
    }
    return string(newS)
}
