package main
import (
    "fmt" 
)

func main(){
    num := 5
    ret := getKthMagicNumber(num)
    fmt.Println("num:",5, " ans:",ret)
}

func getKthMagicNumber(k int) int {
    if k == 0{
        return 1
    }
    ary := make([]int,k)
    p3,p5,p7 := 0,0,0
    ary[0]=1
    for i:=1;i< k;i++{
        ary[i]  = min(ary[p3]*3,min(ary[p5]*5,ary[p7]*7))
        //fmt.Printf("test p3*3:%d,p5*5:%d,p7*7:%d\n",p3*3,p5*5,p7*7)
        if ary[i]== ary[p3] * 3  {
            //if ary[i]>1{ary = append(ary,num)}
            p3++ 
         //   fmt.Println("P3 + 1",p3," num:",ary[i])
        } 
        if ary[i]== ary[p5]*5 {
            //if num >1{ary = append(ary,num)}
            p5++
          //  fmt.Println("P5 + 1",p5," num:",ary[i])
        } 
        if ary[i]== ary[p7]*7 {
            //if num >1{ary = append(ary,num)}
            p7++
           // fmt.Println("P7 + 1",p7," num:",ary[i])
        } 


    }
   // fmt.Println("ary:",ary)
    return ary[k-1] 
}

func min(a int,b int) int{
    if a <= b{
        return a
    }else{
        return b
    }
}
