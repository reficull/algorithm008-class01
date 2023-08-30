"""
183. 从不订购的客户
Customers
表：

+-------------+---------+
| Column
Name | Type |
+-------------+---------+
| id | int |
| name | varchar |
+-------------+---------+
在
SQL
中，id
是该表的主键。
该表的每一行都表示客户的
ID
和名称。
Orders
表：

+-------------+------+
| Column
Name | Type |
+-------------+------+
| id | int |
| customerId | int |
+-------------+------+
在
SQL
中，id
是该表的主键。
customerId
是
Customers
表中
ID
的外键(Pandas
中的连接键)。
该表的每一行都表示订单的
ID
和订购该订单的客户的
ID。


找出所有从不点任何东西的顾客。

以
任意顺序
返回结果表。

结果格式如下所示。



示例
1：

输入：
Customers
table:
+----+-------+
| id | name |
+----+-------+
| 1 | Joe |
| 2 | Henry |
| 3 | Sam |
| 4 | Max |
+----+-------+
Orders
table:
+----+------------+
| id | customerId |
+----+------------+
| 1 | 3 |
| 2 | 1 |
+----+------------+
输出：
+-----------+
| Customers |
+-----------+
| Henry |
| Max |
+-----------+
"""

import pandas as pd
def find_customers(customers: pd.DataFrame, orders: pd.DataFrame) -> pd.DataFrame:
    df = customers[~customers['id'].isin(orders['customerId'])]
    df = df[['name']]
    df.columns = ['Customers']
    return df

# left join method
def find_customers1(customers: pd.DataFrame, orders: pd.DataFrame) -> pd.DataFrame:
    df = customers.merge(orders, left_on='id', right_on='customerId', how='left')
    df = df[df['customerId'].isna()]
    df = df[['name']].rename(columns={'name':'Customers'})
    return df

if __name__ == "__main__":
    lst = [1,2,3,4]
    lst1 = ['Joe','Henry','Sam','Max']
    customers = pd.DataFrame(list(zip(lst,lst1)),columns=['id','name'])
    lst2 = [1,2]
    lst3 = [3,1]
    orders = pd.DataFrame(list(zip(lst2,lst3)),columns=['id','customerId'])
    df = find_customers1(customers, orders)
    print(df)