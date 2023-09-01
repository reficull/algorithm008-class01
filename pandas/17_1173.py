"""
1173. 即时食物配送 I
简单
51
相关企业
SQL Schema
Pandas Schema
配送表: Delivery

+-----------------------------+---------+
| Column Name                 | Type    |
+-----------------------------+---------+
| delivery_id                 | int     |
| customer_id                 | int     |
| order_date                  | date    |
| customer_pref_delivery_date | date    |
+-----------------------------+---------+
delivery_id 是表的主键(具有唯一值的列)。
该表保存着顾客的食物配送信息，顾客在某个日期下了订单，并指定了一个期望的配送日期（和下单日期相同或者在那之后）。


如果顾客期望的配送日期和下单日期相同，则该订单称为 「即时订单」，否则称为「计划订单」。

编写解决方案统计即时订单所占的百分比， 保留两位小数。

返回结果如下所示。



示例 1:

输入：
Delivery 表:
+-------------+-------------+------------+-----------------------------+
| delivery_id | customer_id | order_date | customer_pref_delivery_date |
+-------------+-------------+------------+-----------------------------+
| 1           | 1           | 2019-08-01 | 2019-08-02                  |
| 2           | 5           | 2019-08-02 | 2019-08-02                  |
| 3           | 1           | 2019-08-11 | 2019-08-11                  |
| 4           | 3           | 2019-08-24 | 2019-08-26                  |
| 5           | 4           | 2019-08-21 | 2019-08-22                  |
| 6           | 2           | 2019-08-11 | 2019-08-13                  |
+-------------+-------------+------------+-----------------------------+
输出：
+----------------------+
| immediate_percentage |
+----------------------+
| 33.33                |
+----------------------+
解释：2 和 3 号订单为即时订单，其他的为计划订单。
"""

import pandas as pd

def food_delivery(delivery: pd.DataFrame) -> pd.DataFrame:
    immediate = delivery[delivery["order_date"] == delivery["customer_pref_delivery_date"]]
    ans = pd.DataFrame({"immediate_percentage": [round(len(immediate)/len(delivery) * 100,2)]})
    return ans