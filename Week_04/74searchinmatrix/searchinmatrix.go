package main
import (
    "fmt" 
    "strings"
)

func main(){
    matrix := [][]int{
        {1,3,5,7},{10,11,16,20},{23,30,34,50},

        //{1},
        //{3},
        //{1,   3,  5,  7},
        //{10, 11, 16, 20},
        //{23, 30, 34, 50},   
        //{51, 55, 60, 65},   
    }
    
    ret := searchMatrix(matrix,3)
    fmt.Println("ret:",ret)
}

func searchMatrix(matrix [][]int, target int) bool {
    cnt := len(matrix) 
    if cnt == 0{
        return false
    }
    c := len(matrix[0]) 
    if c == 0{
        return false
    } 
    l := 0
    r := cnt
    for l<= r && r <= cnt{
        mid := (l + r) >> 1
        fmt.Printf("mid row:%d\n,first num:%d,last:%d",mid,matrix[mid][0],matrix[mid][c-1])//,printLine(matrix[mid]))

        if matrix[mid][0] <= target && matrix[mid][c-1] >= target || len(matrix)==1{
            fmt.Printf("start bin search row:%d,row:%v\n",mid,matrix[mid])
            return binSearch(matrix[mid],target)
        }else if matrix[mid][c-1] < target && (mid+1  <= r)  {
            l = mid +  1
            if l>cnt-1{
                return false
            }
            fmt.Println("on the right, left:(mid + 1) < (r - 1) left:",l)
        }else if matrix[mid][0] > target && mid-1 >= l{
            r = mid - 1
            fmt.Println("on the left, right:",r)
        }else{
            return false
        }
        fmt.Printf("mid:%d,left:%d,right:%d\n",mid,l,r)
    }
    return false 
}

func binSearch(nums []int,target int) bool{  
    fmt.Println("in bin search")
    cnt := len(nums)
    l := 0
    r := cnt-1
    for l <= r{
        mid := (l + r) / 2
        fmt.Printf("in bin search mid col:%d",mid)
        if nums[l]==target || nums[r]==target || nums[mid] == target {
        //    fmt.Println("found!")
            return true
        }
        
        if nums[mid] < target{
         //   fmt.Printf("mid:%d,target:%d,on the right\n",nums[mid],target)
            l= mid+1
        }
        if nums[mid] > target{
          //  fmt.Printf("mid:%d,target:%d,on the left\n",nums[mid],target)
            r = mid-1
        }
        
    }
    return false
}

func printLine(l []int) string{
    var str []string
    for _,num := range l{
        str = append(str,fmt.Sprintf("%s",num))
    }
    return strings.Join(str,",")
}
