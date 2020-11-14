# 周八高级搜索

dfs复习 

visited = set()

def dfs(node,visited):
	if node in visited:
		return
	
	visited.add(node)
	#do current task
	
	for next_node in node.Children():
		if not next_node in visted:
			dfs(next_node,visited)
			
非递归写法 

def DFS(self,tree)
	if tree.root is None:
		return
		
	visited,stack = [],[tree.root]
	while stack:
		node = stack.pop()
		visted.add(node)
			
		//do process node
		nodes = generate_nodes()
		stack.push(nodes)

BFS 模板

def BFS(graph,start,end)
	visited = set()
	queue = []
	queue.append([start])
	
	while queue:
		node = queue.pop()
		visited.add(node)
		
		//do process
		nodes = generate_new_nodes()
		queue.append([nodes])
		

			
高级搜索主要方法:
1. 减少重复 fibonacci 使用cache , 剪枝(生成括号问题)
2. 搜索方向:
	dfs 深度优先 
	bfs 广度优先
	
3. 双向搜索,启发式搜索


位运算

XOR:
x ^ 0 = x
x ^ 1s = ~x  ( 1s = ~0)
x ^ (~x) = 1s
x ^ x = 0
a = b ^ c => a ^ b = c => a ^c = b    //交换两个数

指定位置 的位运算

1.将最右边N位清0:  x & (~0 << n)
2.获取X的第n位值 : 1 & (x >> n)
3.获取X的第n位幂值: x & (x << n)
4.仅将第n位置1: x | (1 <<n)
5.仅将第n位置0: x&(~(1<<n)  )
6.将x最高位到第n位清0：x & ( (1<< n) -1)

判断奇偶：
x % 2 ==> x & 1    判断最低位是否为1

x >> 1 ==> x/2
mid = (left + right)/2  ==> (left + right) >> 1

x & (x-1)  清0最低位的1
x & -x 得到最低位的1


实战题目 / 课后作业
https://leetcode-cn.com/problems/number-of-1-bits/
https://leetcode-cn.com/problems/power-of-two/
https://leetcode-cn.com/problems/reverse-bits/
https://leetcode-cn.com/problems/n-queens/description/
https://leetcode-cn.com/problems/n-queens-ii/description/
https://leetcode-cn.com/problems/counting-bits/description/