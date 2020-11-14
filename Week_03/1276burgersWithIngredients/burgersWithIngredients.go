package main
import (
    "fmt" 
)
var jumboCnt = 0
var smallCnt = 0
func main(){
    toma := 2 
    cheese := 1 
    ans := numOfBurgers(toma,cheese)
    fmt.Println("start :toma",toma," cheese:",cheese," ans:",ans)
}
func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
    var ret []int
    if (tomatoSlices - 2 * cheeseSlices) % 2 == 0{
        j := (tomatoSlices - 2 * cheeseSlices) / 2
        s := cheeseSlices - j
        if j>=0 && s >=0 {
            ret = []int{j,s}
        }
    }
    return ret
}



   /* 
//s = (cheeseslices - tomatoslices/4) * 2
//j = cheeseslices - s 
// make small first
    s := (cheeseSlices - tomatoSlices/4) * 2 
    fmt.Println("small cnt:",s)
    j := cheeseSlices - s
    fmt.Println("jumbo cnt:",j, " cheeseSlices:",cheeseSlices)
    fmt.Printf("small first toma:%d,cheese:%d,jumboCnt:%d,smallCnt:%d\n",tomatoSlices,cheeseSlices,j,s)
    if j>=0 && s>=0{
        return []int{j,s}
    }else{
        // make jumbo first
        //j = cheeseSlices/2 + tomatoSlices/2
        // s = cheeseSlices - j
        j = cheeseSlices/2 + tomatoSlices/2
        s = cheeseSlices - j

        fmt.Printf("jumbo first toma:%d,cheese:%d,jumboCnt:%d,smallCnt:%d\n",tomatoSlices,cheeseSlices,j,s)
        if j>=0 && s>=0{
            return []int{j,s}
        } 

    }
    return []int{}
    */
