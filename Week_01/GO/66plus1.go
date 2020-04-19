package main
import (
    "fmt" 
    "math"
)

func plusOne(digits []int) []int {
    //array2dec(digits)
    return tailP1(&digits);

}

func tailP1(digits *[]int) []int {
    if (*digits)[len(*digits)-1] < 9{
        (*digits)[len(*digits)-1]++;
        return (*digits);
    }else{
        (*digits)[len(*digits)-1] = 0;
        c := len(*digits)-2;
        return carry1(digits ,c );
    }
}

func carry1(digits *[]int, i int) []int{
    if i >=0 {
        if (*digits)[i] < 9{
            (*digits)[i]++;
            return (*digits);
        }else{
            (*digits)[i] = 0;
            left1 := i-1;
            return carry1(digits,left1);
        }
    }else{
        // add one at beginning
        //var ret []int  = (*digits)[0:len(*digits)+1];
        var ret []int  = make([]int,len(*digits)+1,len(*digits)+1);
        fmt.Println("before copy:",ret);
        copy(ret[1:],(*digits)[:]);
        fmt.Println("after copy:",ret);
        ret[0]=1;
        fmt.Println("new slice len:",len(ret));
        fmt.Println("after set 1:",ret);
        return ret;
    }
}


func array2dec(digits []int) int{
    ret := float64(0)
    idec := 0
    for i:= len(digits)-1;i>=0;i--{
        tenPow := math.Pow(10,float64(idec))
        d := float64(digits[i]) * tenPow 
        ret += d
        fmt.Println("i:",i," digit:",d," pow:",tenPow," ret:",ret)
        idec++
    }
    fmt.Println("result:",ret)
    return int(ret)
}

func main(){
    //di := []int{1,2,3,4}
    di := []int{9,9}
    r := plusOne(di)
    fmt.Println(r)
}

