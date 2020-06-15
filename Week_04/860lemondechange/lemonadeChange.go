package main
import (
    "fmt" 
)

func main(){
    bils := []int{5,5,5,5,10,5,10,10,10,20}
    ret := lemonadeChange(bils)
    fmt.Println("ret:",ret)
}

func lemonadeChange(bills []int) bool {
    if bills == nil {
        return false
    }
    var getBils  []int
    for _,m := range bills{
        if m == 5{
            getBils = append(getBils,m)
            continue
        }
        if m > 5{
            getBils,sucesChange := change5(getBils,m)
            if sucesChange{
                getBils = append(getBils,m)
            }else{
                return false
            }
        }
    }
    return true 
}

func change5(bills []int,newBil int) ([]int, bool){
    target := newBil - 5
    fmt.Printf("change target:%d,new bil:%d\n",target,newBil)
    for i:=len(bills)-1;i>0;i--{
        target = target - bills[i] 
        bills = bills[0:i]
        if target == 0{
            fmt.Printf("true after change new bills:%v\n",bills)
            return bills, true
        }
        if target < 0{
            fmt.Printf("false after change new bills:%v\n",bills)
            return bills,false
        } 
    }
    return bills,false
}
