

func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0 || k <= 0 || k > len(nums){
        return nil

    }
    var maxNums []int
    maxNow:=max(nums[0:k])
    maxNums=append(maxNums,maxNow)
    for i:=1;i<=len(nums)-k;i++{
        if nums[i-1]==maxNow{
            maxNow=max(nums[i:i+k])
            maxNums=append(maxNums,maxNow)

        }else{
            if nums[i+k-1]<=maxNow{
                maxNums=append(maxNums,maxNow)

            }else{
                maxNow=nums[i+k-1]
                maxNums=append(maxNums,maxNow)

            }

        }

    }
    return maxNums

}

func max(nums []int) int {
    left:=0
    right:=len(nums)-1
    for left<right{
        if nums[left]>nums[right]{
            right--

        }else{
            left++

        }

    }
    return nums[left]

}
