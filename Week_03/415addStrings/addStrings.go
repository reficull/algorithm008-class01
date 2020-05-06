package main
import (
    "fmt" 
    //"strconv"
)

func main(){
    num1 := "23454567"
    num2 := "11128888888"
    ans := addStrings(num1,num2)
    fmt.Printf("num1:%s num2:%s ans:%s\n",num1,num2,ans)
}

func addStrings(num1 string, num2 string) string {
    var (
        j, i, k, t int
        tmp int
        s string

    )
    carry := 0
    j = len(num2) - 1
    i = len(num1) - 1
    if i >= j {
        k =  i + 1

    }else{
        k = j + 1

    }
    num := make([]byte,k+1)
    for ; j >= 0 && i >= 0;{
        tmp = int(num1[i]+num2[j]-96) + carry
        //fmt.Printf("%d\n",tmp)
        if tmp >= 10 {
            carry = 1

        } else {
            carry = 0;

        }
        num[k] = byte((tmp % 10) + 48)
        //fmt.Printf("%c\n", num[k])
        k = k - 1
        i = i -1
        j = j - 1

    }
    if i >= 0 {
        t = i
        s = num1

    }else {
        t = j
        s = num2

    }
    for ; t >= 0 ; t-- {
        tmp = int(s[t]-48) + carry
        //fmt.Println(tmp)
        if tmp >= 10 {
            carry = 1

        } else {
            carry = 0;

        }
        num[k] = byte((tmp % 10)+48)
        k = k - 1 
        //fmt.Printf("%c",num[k])

    }
    var str string
    if carry == 1 {
        num[k] = '1'
        str = string(num[k:])

    }else{
        str = string(num[k+1:])

    }

    return str
    /*
    nums := []string{num1,num2}
    big := Bigger(nums)
    small := 1 - big
    bigStr := Reverse(nums[big])
    smallStr := Reverse(nums[small])
    fmt.Println("bigger :",nums[big]," big i:",big," smaller:",nums[small])
    increase := 0 
    ans := make([]rune,0)
    for ci:=0;ci < len(bigStr);ci++{
        num1 := int(bigStr[ci]-'0')
        num2 := 0 
        if ci < len(smallStr){
            num2 = int(smallStr[ci]-'0')
        }
        numP := num1 + num2 + increase
        if numP > 9{ increase = 1}else{ increase = 0}
        c := byte( numP % 10)
        ans = append(ans, rune(c))
        fmt.Printf("nums bigger:%d,ans:%v\n",int(bigStr[ci]-'0'),ans)
    } 

    if increase == 1{ 
        ans = append(ans,rune('1'))
    }
    fmt.Printf("ans :%v",ans)
   // ans = ReverseR(ans)
    return string(ans) 
    */
}


func ReverseR(runes []rune) string{
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)

}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func Bigger(n []string) int{
    l1 := len(n[0])
    l2 := len(n[1])
    fmt.Printf("1 len:%d,2 :%d\n",l1,l2)
    if l1 > l2{
        return 0
    }else{
        return 1
    }
}
