package main

import(
    "fmt"
)
func main(){

    row := [9]int{0,0,0,0,0,0,0,0,0}
    
    bit := 1 << '4' 
    bit1 :=  1<< 4  
    row[0] |= bit
    row[0] |= bit1
    byte1 := byte('9')
    byte2 := byte(4)
    fmt.Printf("byte1:%#b,byte2:%#b\n",byte1,byte2)
    //ret := row[0]&bit
    flag := (row[0] & bit1  == bit1)
    fmt.Printf("ret:%v\n",flag)
   row[0] &= ^bit
    flag = (row[0] & bit == bit)
    flag1 := (row[0] & bit1 == bit1)
    fmt.Printf("ret:%v,flag:%v,flag1:%v\n",row[0],flag,flag1)
}

