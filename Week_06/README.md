# 周六动态规划

https://leetcode-cn.com/problems/climbing-stairs/description/
https://leetcode-cn.com/problems/triangle/description/
https://leetcode.com/problems/triangle/discuss/38735/Python-easy-to-understand-solutions-(top-down-bottom-up)
https://leetcode-cn.com/problems/maximum-subarray/
https://leetcode-cn.com/problems/maximum-product-subarray/description/
https://leetcode-cn.com/problems/coin-change/description/


https://leetcode-cn.com/problems/minimum-path-sum/
https://leetcode-cn.com/problems/decode-ways
https://leetcode-cn.com/problems/maximal-square/
https://leetcode-cn.com/problems/task-scheduler/
https://leetcode-cn.com/problems/palindromic-substrings/
困难
https://leetcode-cn.com/problems/longest-valid-parentheses/
https://leetcode-cn.com/problems/edit-distance/
https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k/
https://leetcode-cn.com/problems/frog-jump/
https://leetcode-cn.com/problems/split-array-largest-sum
https://leetcode-cn.com/problems/student-attendance-record-ii/
https://leetcode-cn.com/problems/minimum-window-substring/
https://leetcode-cn.com/problems/burst-balloons/


实战
https://leetcode-cn.com/problems/house-robber/
https://leetcode-cn.com/problems/house-robber-ii/description/
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/#/description
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/yi-ge-fang-fa-tuan-mie-6-dao-gu-piao-wen-ti-by-l-3/

高级DP
https://leetcode-cn.com/problems/perfect-squares/
https://leetcode-cn.com/problems/edit-distance/ （重点）
https://leetcode-cn.com/problems/jump-game/
https://leetcode-cn.com/problems/jump-game-ii/
https://leetcode-cn.com/problems/unique-paths/
https://leetcode-cn.com/problems/unique-paths-ii/
https://leetcode-cn.com/problems/unique-paths-iii/
https://leetcode-cn.com/problems/coin-change/
https://leetcode-cn.com/problems/coin-change-2/


各位周末好，加两道给大家周末的练习题：比较典型，也算很难，大家练习一下！
https://leetcode-cn.com/problems/contains-duplicate/
https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/


编辑距离
给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')

单词A插入一个字符和单词B删除一个字符是等价的
单词A替换一个字符和单词B替换一个字符等价

dp[i][j] 定义 word1 在i 位和word2在j位时使用了多少次编辑操作

如果word1在i位删除,则之前的子问题的距离是word1[i-1] , word2[j]
如果 word1在i位插入,则之前的代价是word1[i], word2[j-1]
如果word1在i位替换,则比较了word1[i-1],word2[j-1]

选择操作量最小的那个+1, 就是记录本次操作

如果 word1和 word2 相同 在第i位相同
则dp[i][j]的编辑距离是
Min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])

如果不同则距离是
Min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1]) + 1



|w     |""|r   |ro |ros|
|---    |---|---|---|--- |
|""     |---|1   |2  | 3  |
|h      | 1  |1  | 2 |3  |
|ho    | 2  | 2  |1 |2   | 
|hor   | 3  |  2 |2  | 2 |  
|hors | 4  |  3 |3  |3   | 
|horse| 5  | 4  |4  |3    | 



https://leetcode-cn.com/problems/house-robber/

198. 打家劫舍
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

 

示例 1：
steal  = 1  dp[i][j] = nums[i]
not steal = 0
|---|      ||
|---|--0--|-----1-|-2-|
| 0 |  0  |      0  |0  |
| 1 |  1  |      0  |1  |
| 2 |  0  |      2  |2  |   Max(nums[i],nums[i-1])
| 3 |  3  |      0  |4  |   nums[i]+nums[i-2]
| 1 |  0  |      1  |3  |
| f |  0  |           |4>3  |


输入：[1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。
	 	
		
	
	 
	 
	dp[i][j] 表示 第i的房子,j=1表示 偷,0表示 不偷
	不偷
	dp[i][0] = Max(dp[i-1][0] , dp[i-1][1])
	偷
	dp[i][1] =  nums[i] + Max(dp[i-1][0] , dp[i-1][1])
	
	




213 打家劫舍2

与1的区别是房屋环状,第一家和最后一家相连
f(0) = nums[0]
f(1) = Max(nums[0],nums[1])
f(2)=  Max(nums[2] + f(0), nums[1])

f(n) = Max( nums[n] + f(n-2), f(n-1))

for i:=2;i<len(nums);i++{
	dp[i]=Max(nums[i] + dp[i-2],dp[i-1])
}

return Max( maxInAry(dp[:-1]),maxInAry(dp[1:len(nums)]

