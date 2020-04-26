package main
import (
    "fmt" 
)

func main(){
    nums1 := []int{4,9,5}
    nums2 := []int{9,4,9,8,4}
    nums := Intersect(nums1,nums2)
    fmt.Println("nums1:",nums1," nums2:",nums2," intersect:",nums)
}

func Intersect(nums1 []int, nums2 []int) []int {
    recordMap := make(map[int]int,0)
    numsRet := make([]int,0)
    for i:=0;i<len(nums1);i++{
        recordMap[nums1[i]]++
    }

    for i:=0;i<len(nums2);i++{
        if  recordMap[nums2[i]] > 0{
            numsRet = append(numsRet,nums2[i])
            recordMap[nums2[i]]--
        }
    }

    return numsRet;
    /* best
    m1 := make(map[int]int)
    res := make([]int,0)
    for _, v := range nums1 {
        m1[v] += 1

    }

    for _, v := range nums2 {
        if m1[v] > 0 {
            res = append(res, v)
            m1[v]--

        }

    }
    return res
    */
}

