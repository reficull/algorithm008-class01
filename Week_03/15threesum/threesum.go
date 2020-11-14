package main
import (
    "fmt" 
    "sort"
//    "strings"
    "encoding/json"
    "io/ioutil"
)
func main(){
    
    //nums := []int{-1, 0, 1, 2, -1, -4}
    nums := loadFile("data.json")
    ret := threeSum(nums)
    fmt.Printf(" nums:%v,ans:%v\n",nums,ret)
}

func loadFile(file string) []int{
    data, err := ioutil.ReadFile("data.json")
    if err != nil {
        fmt.Println(err)
        return nil
    }
    //fmt.Print("data:  ",string(data))
    var nums []int
    err = json.Unmarshal(data, &nums)
    if err != nil {
        fmt.Println(err)
        return nil
    }
    return nums
}

func threeSum(nums []int) [][]int {
    /*
    sort.Ints(nums) 
    numsMap := make(map[string]bool)
    ret := make([][]int,0)
    pR := len(nums)-1
    for p1 := 0;p1 < pR;p1++{
        for pL,pR :=p1+1,len(nums)-1; pL < pR;{
            if (nums[p1] + nums[pL] + nums[pR]) == 0{
                //fmt.Printf("check=0 :%d,%d,%d\n",nums[p1],nums[pL],nums[pR])
                ans := []int{nums[p1],nums[pL],nums[pR]}
                anStr := strings.Trim(strings.Replace(fmt.Sprint(ans), " ", ",", -1), "[]")
                //fmt.Printf("str:%v\n",anStr)
                if !numsMap[anStr]{
                    //fmt.Printf("found :%v",ans)
                    ret = append(ret,ans)
                    numsMap[anStr] = true
                }
                pL++
            }
            if (nums[p1] + nums[pL] + nums[pR])<0{
                //fmt.Printf("check <0:%d,%d,%d\n",nums[p1],nums[pL],nums[pR])
                pL++
            } 
            if (nums[p1] + nums[pL] + nums[pR])>0{
                //fmt.Printf("check >0:%d,%d,%d\n",nums[p1],nums[pL],nums[pR])
                pR--
            } 

        } 
        //fmt.Println("end inner for:",p1,pL,pR)
    }
    return ret
    */
    if nums == nil || len(nums) < 3 {
        return nil

    }
    sort.Ints(nums)
    ans := [][]int{}
    length := len(nums)
    for k := 0 ; k < length - 2 ; k++ {

        if nums[k] > 0 {
            break

        }
        if k > 0 && nums[k] == nums[k-1] {
            continue

        }

        l := k + 1
        r := length - 1
        for l < r {
            if nums[l] + nums[r] < -nums[k] {
                l++
                continue

            }else if nums[l] + nums[r] > -nums[k] {
                r--
                continue

            }else if nums[l] + nums[r] == -nums[k] {
                ans = append(ans, []int{nums[k], nums[l], nums[r]})
                for l < length-1 && nums[l] == nums[l+1] {
                    l++

                }
                for r > k+1 && nums[r] == nums[r-1] {
                    r--

                }
                l++
                r--

            }


        }

    }
    return ans
}

