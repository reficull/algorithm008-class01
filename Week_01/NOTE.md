# 周一算法心得

## 几种数据结构时间复杂度

### Array
prepend   O(1)
appepend  O(1)
lookup    O(1)
insert    O(n)
delete    O(n)

### LinkedList
prepend  O(1)
append   O(1)
lookup   O(n)
insert   O(1)
delete   O(1)

## 使用双指针或map降低复杂度O(N^2)->O(N)

[283 移动0到数组末尾](https://leetcode-cn.com/problems/move-zeroes)
最佳算法GO:
双指针, 一个是非0数指针i,一个是0组最左边那个0的指针,是用 i - count来得到,有点晦涩, count是0的数量,因为不断把发现的非0值和最左边的0交换位置, 所以最左边的0算法是 i - count
```
l := len(nums)   
count := 0   
for i := 0;i < l;i++{
	if nums[i] == 0{ 
		count++  
	}else{   
		nums[i - count] , nums[i] = nums[i] , nums[i - count]
	}
}   
```
[26 删除重复项](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/submissions/)
双指针, n为慢指针, 记录左边不重复项,i快指针寻找下一个非重复项
```
 ln := len(nums)                                                                    n := 1                                                                           for i := 1; i < ln; i++ {                                                            		if nums[i] != nums[i-1] {                                                        			nums[n] = nums[i]                                                             			n++                                                                          }                                                                            }                                                                              return n               
```
[1 两数之和](https://leetcode-cn.com/problems/two-sum/)
最佳题解, 使用map, 将已经遍历过的数放到map中, 在测试下一个值的时候, 使用 target - 当前测试值, 如果结果存在于map中则找到解
此法初看有点晦涩

```
res := make(map[int]int, 0)
for i, num := range nums {
		res[num] = i
		if j, ok := res[target-num]; ok {
				fmt.Println("test: ",target-num, " j:",j)
				return []int{i, j}
		}                                                                            }
		
return nil
```
[917 ](https://leetcode-cn.com/problems/squares-of-a-sorted-array/)
给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。

```
ss := []rune(S)
n := len(S)                                                                       i,j := 0,n-1                                                                        for i < j{                                                                            if !isLetter(ss[i]){
			i++
			continue                                                                        }                                                                              if !isLetter(ss[j]){                                                                  	j--
			continue                                                                        }                                                                              ss[i],ss[j] = ss[j],ss[i]                                                              i++                                                                              j--                                                                            }                                                                                return string(ss)               
```
[977有序数的平方](https://leetcode-cn.com/problems/squares-of-a-sorted-array/)
给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。
最麻烦的是有负数的数组, 此时用双指针,放在最左边和最右边开始往中间归并直到相遇, 因为负数的平方是最左边的最大,所以这样可以同时比较两个最大值, 将最大的放到结果数组右边, 使用第三个指针来放结果
```
   if len(A) == 0 {
        return nil
    }
    results := make([]int, len(A))
    end := len(results) - 1
    for left, right := 0, len(A) - 1; left <= right; {
        leftVal := A[left] * A[left]
        rightVal := A[right] * A[right]
        if leftVal > rightVal {
            results[end] = leftVal
            left++
        } else {
            results[end] = rightVal
            right--
        }
        end--
    }

    return results    
```
[15 三数之和](https://leetcode-cn.com/problems/3sum/)
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。
使用三指针,其中1个从最左到最右边枚举,另两个在最左和最右往中间归并, 数组必须先排序这样才不会漏掉
```
   sort.Sort(nums)
    var saveArray [][]int
    var addnum int
    //index为中间值
    var index,start,end int
    for index=1;index<len(nums)-1;index++{
        start,end = 0,len(nums)-1
        //去重的关键步骤
        if index>1&&nums[index]==nums[index-1]{
            start=index-1
        }
        for start<index&&index<end{
            addnum = nums[start] + nums[index] + nums[end]
            if start>0 && nums[start]==nums[start-1]{
                start++
                continue
            }
            if end<len(nums)-1 && nums[end]==nums[end+1]{
                end--
                continue
            }
            if addnum==0{
                saveArray=append(saveArray,[]int{nums[start],nums[index],nums[end]})
                start++
                end--
            }else if addnum>0{
                end--
            }else{
                start++
            }
        }
    }
    return saveArray

```

[206 返转链表](https://leetcode-cn.com/problems/reverse-linked-list/solution/go-by-bingohuang-2/)
思路:使用新链表, 遍历老链表,不断把新链表的头指向老链表的下一个节点直到完成
```
func reverseList1(head *ListNode) *ListNode {
	var newHead, next *ListNode

	for head != nil {
		next = head.Next
		head.Next = newHead
		newHead = head
		head = next
	}

	return newHead
}
```

[11 盛最多水的容器](https://leetcode-cn.com/problems/container-with-most-water/)

给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

思路: 使用双指针, 从最左和最右向中间归并, 容量为两指针距离乘较短的值, 比较如果左边较小则左指针右移, 右边较小则右边指针左移


```
func maxArea(height []int) int {
    var i,j,res=0,len(height)-1,0
    for i<j{
        res=max(res,(j-i)*min(height[i],height[j]))
        if height[i]>height[j]{
            j--
        }else{
            i++
        }
    } 
    return res

}

```

[141 环形链表](https://leetcode-cn.com/problems/linked-list-cycle/)
给定一个链表，判断链表中是否有环。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
思路: 使用hashMap, 记录遍历过的节点
```
func hasCycle(head *ListNode) bool {    // hash表
    hash := make(map[*ListNode]int)     // 开一个哈希表记录该节点是否已经遍历过，值记录节点索引
    for head != nil {
        if _,ok := hash[head]; ok {     // 该节点遍历过，形成了环
            return true
        }
        hash[head] = head.Val           // 记录该节点已经遍历过
        head = head.Next
    }
    return false
}
```

[24. 两两交换链表中的节点](https://leetcode-cn.com/problems/swap-nodes-in-pairs/)

给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

思路:使用前置指针
重点代码，由于golang与python的等号赋值的特殊性，我们可以一行代码完成整个交换过程
这一行代码做了4件事，以prev->1->2->3->4->nil为例

prev.Next 也就是1，赋值给2的下一个元素，也就是把2接在1的后面
把2的下一个元素，也就是3，赋值给1的下一个元素，也就是把1接到3后面
把1的下一个元素，也就是2，赋值给prev的下一个元素，因为2和1已经调换顺序，需要把prev重新连接到2的前面
把prev的下一个元素，也就是1，赋值给prev。注意！这里的prev还没有发生改变，这一行是同时生效，所以现在prev的下一个元素还是1，所以看似把prev.Next赋值给了prev是一跳，其实是两跳！因为1和2换了位置，我们又跳到了1，所以是两跳。
最后，出循环的条件是，prev的下一个不是空，并且下下个也不是空值。我们可以发现，思路就是，通过prev指针长了个长胳膊，调换他的下一个和下下个元素，而不是利用三个指针来完成


```
func swapPairs(head *ListNode) *ListNode {
    var prev *ListNode = &ListNode{0, head}  // 构造指向head的prev指针
    var hint *ListNode = prev				//因为prev指针需要移动，留存一个backup
    for prev.Next != nil && prev.Next.Next != nil {
        prev.Next.Next.Next, prev.Next.Next, prev.Next, prev = prev.Next, prev.Next.Next.Next, prev.Next.Next, prev.Next
    }
    return hint.Next
}

```

[25. K 个一组翻转链表](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/)
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

```
func reverseKGroup(head *ListNode, k int) *ListNode {
	var dummy ListNode
	dummy.Next = head
	pre, start, end := &dummy, &dummy, &dummy
	for end != nil {
		start = pre.Next
		for i := 0; i < k; i++ {
			end = end.Next
			if end == nil{
				return dummy.Next
			}
		}
		next := end.Next
		end.Next = nil
		pre.Next = reverse(start)
		pre, end = start, start
		start.Next = next
	}
	return dummy.Next
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

```

[面试 题50](https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/)

在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。
思路: 使用一个int数组来记录, 将每个字符出现次数记在这个数组中
最后遍历字符串, 查找数组中记录的字符,返回第一个值 为1的字符

```
func firstUniqChar(s string) byte {
  m := [256]int{}
    for _, r := range s {
        m[int(r)]++
    }
    for i, r := range s {
        if m[int(r)] == 1 {
            return s[i]
        }
    }
    return ' '
}
```