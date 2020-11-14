# 周二算法心得

## 二叉树
遍历:
1.前序  (Pre-order): 根-左-右
```
def preOrder(self,root):
	if root:
		self.traverse_path.append(root.val)
		preOrder(root,root.left)
		preOrder(root,root.right)
```
2.中序(In-order): 左-根-右
```
def inOrder(self,root):
	if root:
		inOrder(root,root.left)
		self.traverse_path.append(root.val)
		inOrder(root,root.right)
```

3. 后序(Post-order): 左-右-根
```
def postOrder(self,root):
	if root:
		postOrder(root,root.left)
		postOrder(root,root.right)
		self.traverse_path.append(root.val)
```

## 二叉搜索树
搜索,插入,删除 时间复杂度 O(logN)

## 堆和二叉堆

堆 Heap
可以迅速找到最大或最小值的数据结构
根节点最大的堆 叫大顶堆或大根堆, 根节点最小的堆叫小顶堆或小根堆
常见的有二叉堆或非波那契堆

性质一: 是一棵完全树(除了最下面一层可能不满,其它各层都是满)
性质二: 任一节点>=其子节点的值

大顶堆常见操作:
find-max : O(1)
delete-max: O(logN)
insert(create):O(logN) or O(1)

## 哈希算法
关于哈希的应用场景, 一个典型的应用是布隆过滤器, 用来避免过多的查询直接怼到数据库中, 原理是 比如初始化一个20亿的键值对,全部置 0, 然后在用户注册的时候在这个key-value库中通过哈希算法算出个key插入置1, 这样在查询用户是否存在时可以直接查询这个KV库而 不用跑到后面的数据库中, 此处关键是hash算法要足够unique, 避免不同用户生成同一个key导致错误

[299. 猜数字游戏](https://leetcode-cn.com/problems/bulls-and-cows/)

你正在和你的朋友玩 猜数字（Bulls and Cows）游戏：你写下一个数字让你的朋友猜。每次他猜测后，你给他一个提示，告诉他有多少位数字和确切位置都猜对了（称为“Bulls”, 公牛），有多少位数字猜对了但是位置不对（称为“Cows”, 奶牛）。你的朋友将会根据提示继续猜，直到猜出秘密数字。

请写出一个根据秘密数字和朋友的猜测数返回提示的函数，用 A 表示公牛，用 B 表示奶牛。

请注意秘密数字和朋友的猜测数都可能含有重复数字。
你可以假设秘密数字和朋友的猜测数都只包含数字，并且它们的长度永远相等。

思路:
1.  求A遍历一次秘密串比较猜测串即可 O(n)
     求B:
	 将求A过程中猜对之外的秘密串放入hashMap,由于可能有重复项, 存放方式是 map[char]int , 如果有重复,值++
	 遍历猜测串,如果hashMap命中则删除项并B+1
	 
2. 使用一个数组记录secret中数字的出现次数,在第二个循环中测试,
	1) 如果 猜 中位置和数字, 公牛加一,奶牛要减一,因为重复数字存在可能
	2) 如果没猜中
	    a)如果数字存在,奶牛加一
		
```
func getHint(secret string, guess string) string {                                             //best NO.1                                                                           m := make([]int, 128)                                                                 l := len(secret)                                                                     for i := 0; i < l; i++ {                                                                   m[secret[i]]++                                                                 }                                                                               fmt.Println("count array:",m)                                                           aN, bN := 0, 0                                                                     for i := 0; i < l; i++ {                                                                   n := m[guess[i]]                                                                     fmt.Println("n:",n)                                                                   if secret[i] == guess[i] {                                                                 if n==0{                                                                             bN--                                                                             fmt.Printf("guess right,n==0,bN--:%d\n",bN)                                               }else{                                                                               m[guess[i]]--                                                                       fmt.Printf("guess right,n=%d,%s count --:%d\n",n,string(guess[i]),m[guess[i]])                       }                                                                               aN++                                                                         } else {                                                                             if n > 0 {                                                                             bN++                                                                             m[guess[i]]--                                                                       fmt.Printf("guess wrong,n=%d,%s count --:%d , bN:%d\n",n,string(guess[i]),m[guess[i]],bN)               }                                                                           }                                                                           }                                                                               return  fmt.Sprintf("%dA%dB", aN, bN)                 
```

[350. 两个数组的交集 II](https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/)
给定两个数组，编写一个函数来计算它们的交集。

1. 思路 第一个循环将第一个数组使用一个hashMap计数,记录每个数出现的次数, 在第二个循环遍历第二个数组 ,看有没有记录数组中出现的数与之相同,有则记入结果数组

[242. 有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram/)

给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

```
func Intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)                                                            
	res := make([]int,0)                                                              
	for _, v := range nums1 {                                                                   m1[v] += 1                                                                     }                                                                          
	for _, v := range nums2 {
		if m1[v] > 0 {
			res = append(res, v)
			m1[v]--                                                                       	}
	}                                                                             	return res
}
```

[49. 字母异位词分组](https://leetcode-cn.com/problems/group-anagrams/)
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
```
func nthUglyNumber(n int) int {                                                              if n == 1 {
		return 1
	}
	i, j, k := 0, 0, 0
	result := []int{1}
	res := 0
	for n > 1 {
		a := result[i] * 2
		b := result[j] * 3
		c := result[k] * 5
		res = min(min(a, b), c)
		if res == a {
			i++
		}
		if res == b {
			j++
		}
		if res == c {
			k++
		}
		result = append(result, res)
		n--
	}
	return result[len(result)-1]
}                                                       }              
```

[509. 斐波那契数](https://leetcode-cn.com/problems/fibonacci-number/)
斐波那契数，通常用 F(n) 表示，形成的序列称为斐波那契数列。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：

F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
给定 N，计算 F(N)。

```
var goldenRatio float64 = float64((1 + math.Sqrt(5)) / 2);
    return int(math.Round(math.Pow(goldenRatio, float64(N)) / math.Sqrt(5)));
```

[1021. 删除最外层的括号](https://leetcode-cn.com/problems/remove-outermost-parentheses/)
有效括号字符串为空 ("")、"(" + A + ")" 或 A + B，其中 A 和 B 都是有效的括号字符串，+ 代表字符串的连接。例如，""，"()"，"(())()" 和 "(()(()))" 都是有效的括号字符串。

如果有效字符串 S 非空，且不存在将其拆分为 S = A+B 的方法，我们称其为原语（primitive），其中 A 和 B 都是非空有效括号字符串。

给出一个非空有效字符串 S，考虑将其进行原语化分解，使得：S = P_1 + P_2 + ... + P_k，其中 P_i 是有效括号字符串原语。

对 S 进行原语化分解，删除分解中每个原语字符串的最外层括号，返回 S 。
```
func removeOuterParentheses(S string) string {                                                     stackState := make([]int,len(S)+1)                                                         stackState[0]=0                                                                     stackHigh := 0                                                                     var newS []rune                                                                     for i:=0;i<len(S);i++{                                                                     if S[i] == '('{                                                                         stackHigh++                                                                   }else if S[i]==')'{                                                                       stackHigh--                                                                   }                                                                               stackState[i+1] = stackHigh                                                             if stackHigh!=0 &&  !(stackHigh==1 && stackState[i]==0){                                             newS = append(newS,rune(S[i]))                                                       }                                                                           }                                                                               return string(newS)                                                               }               
```

[面试题59 - I. 滑动窗口的最大值](https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/)
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。
```
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

```

[589. N叉树的前序遍历](https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/)
给定一个 N 叉树，返回其节点值的前序遍历。

例如，给定一个 3叉树 :

 ![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/narytreeexample.png)
返回其前序遍历: [1,3,5,6,2,4]。

```
func preorderTraversal(root *Node) []int {                                                       res := make([]int,0)                                                                 doTraverse(root,&res)                                                                 return res                                                                     }                                                                               func doTraverse(root *Node,res *[]int){                                                         if root == nil{                                                                         return                                                                       }                                                                               *res = append(*res,root.Val)                                                             for i:=0;i<len(root.Children);i++{                                                             doTraverse(root.Children[i],res)                                                     }                                                                           }       
```

[590. N叉树的后序遍历](https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/)

给定一个 N 叉树，返回其节点值的后序遍历。

例如，给定一个 3叉树 :

 ![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/narytreeexample.png)
返回其后序遍历: [5,6,3,2,4,1].


[429. N叉树的层序遍历](https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/)
给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。

例如，给定一个 3叉树 :

 ![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/narytreeexample.png)

返回其层序遍历:
```
[
     [1],
     [3,2,4],
     [5,6]
]
```
[面试题40. 最小的k个数](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/)
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。



[面试题40. 最小的k个数](https://leetcode-cn.com/problems/sliding-window-maximum/)
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

[347. 前 K 个高频元素](https://leetcode-cn.com/problems/top-k-frequent-elements/)
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。


[enter description here](https://leetcode-cn.com/problems/chou-shu-lcof/)


[200. 岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)


[面试题51. 数组中的逆序对](https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/)
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
示例 1:

输入: [7,5,6,4]
输出: 5

[104. 二叉树的最大深度](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。


[面试题 01.08. 零矩阵](https://leetcode-cn.com/problems/zero-matrix-lcci/)
编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。


[543. 二叉树的直径](https://leetcode-cn.com/problems/diameter-of-binary-tree/)



[1200. 最小绝对差](https://leetcode-cn.com/problems/minimum-absolute-difference/)

[75. 颜色分类](https://leetcode-cn.com/problems/sort-colors/)
思路: 双指针,遇到0放前面,遇到2放后面, 这样会有很多边缘cases有问题,最好三指针, cur和p0分开
```
func sortColors(nums []int)  {
  l := len(nums)    
    for p0,p1,p2 := 0,0, l-1 ;(p1<=p2);{    
        if nums[p1]==0{    
            nums[p0] , nums[p1] = nums[p1],nums[p0]    
            p0++;p1++    
        }else if nums[p1]==2{    
            nums[p1],nums[p2] = nums[p2],nums[p1]    
            p2--;p0--;if p0<0 {p0 = 0};p1=p0
        }else{    
            p1++    
        }         
    }                                  
}
```

[23. 合并K个排序链表](https://leetcode-cn.com/problems/merge-k-sorted-lists/)
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6



