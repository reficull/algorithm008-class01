"""
2082. 富有客户的数量
简单
6
相关企业
SQL Schema
Pandas Schema
表： Store

+-------------+------+
| Column Name | Type |
+-------------+------+
| bill_id     | int  |
| customer_id | int  |
| amount      | int  |
+-------------+------+
bill_id 是这个表的主键(具有唯一值的列)。
每一行包含一个订单的金额及相关客户的信息。


编写解决方案找出 至少有一个 订单的金额 严格大于 500 的客户的数量。

返回结果格式如下示例所示：



示例 1:

输入：
Store 表:
+---------+-------------+--------+
| bill_id | customer_id | amount |
+---------+-------------+--------+
| 6       | 1           | 549    |
| 8       | 1           | 834    |
| 4       | 2           | 394    |
| 11      | 3           | 657    |
| 13      | 3           | 257    |
+---------+-------------+--------+
输出：
+------------+
| rich_count |
+------------+
| 2          |
+------------+
解释：
客户 1 有 2 个订单金额严格大于 500。
客户 2 没有任何订单金额严格大于 500。
客户 3 有 1 个订单金额严格大于 500。
"""

import pandas as pd

def count_rich_customers(store: pd.DataFrame) -> pd.DataFrame:
    rich = store[store['amount']>500]
    count = rich["customer_id"].nunique()
    ans = pd.DataFrame({'rich_count': [count]})
    return ans