圣诞活动预热开始啦，汉堡店推出了全新的汉堡套餐。为了避免浪费原料，请你帮他们制定合适的制作计划。

给你两个整数 tomatoSlices 和 cheeseSlices，分别表示番茄片和奶酪片的数目。不同汉堡的原料搭配如下：

巨无霸汉堡：4 片番茄和 1 片奶酪
小皇堡：2 片番茄和 1 片奶酪
请你以 [total_jumbo, total_small]（[巨无霸汉堡总数，小皇堡总数]）的格式返回恰当的制作方案，使得剩下的番茄片 tomatoSlices 和奶酪片 cheeseSlices 的数量都是 0。

如果无法使剩下的番茄片 tomatoSlices 和奶酪片 cheeseSlices 的数量为 0，就请返回 []。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-burgers-with-no-waste-of-ingredients


x: jumbo 
y: small 

4x + 2y = tomatoSlices 
x + y = cheeseSlices

x = (tomatoSlices - 2y)/4
x = cheeseSlices - y

cheeseSlices = (tomatoSlices - 2y)/4 + y 
cheeseSlices = tomatoSlices/4 + y/2

so: make small first:
y = (cheeseSlices - tomatoSlices/4) * 2
x = cheeseSlices - y 

make jumbo first:

y = cheeseSlices - x

cheeseSlices - x = (cheeseSlices - tomatoSlices/4) * 2 
 x = cheeseSlices/2 + tomatoSlices/2
 y = cheeseSlices - x

