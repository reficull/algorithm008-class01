package main

import(
    "fmt"
)
func main(){
    //num := uint32(00000000000000000000000000001011)
   // num := uint32(00000000000000000000000010000000)
    num := uint64(11111111111111111111111111111101)
    ret := hammingWeight(num)

    fmt.Printf("ret:%d\n",ret)

}

func hammingWeight(num uint32) int {
    cnt := 0
    var bit uint32
    for num>0 {
        bit = num & 1

        if bit == 1{
            cnt++
        }
        num = num >> 1
    }
    return cnt
}
