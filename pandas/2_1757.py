"""
1757. 可回收且低脂的产品
简单
119
相关企业
SQL Schema
Pandas Schema
表：Products

+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| product_id  | int     |
| low_fats    | enum    |
| recyclable  | enum    |
+-------------+---------+
product_id 是该表的主键（具有唯一值的列）。
low_fats 是枚举类型，取值为以下两种 ('Y', 'N')，其中 'Y' 表示该产品是低脂产品，'N' 表示不是低脂产品。
recyclable 是枚举类型，取值为以下两种 ('Y', 'N')，其中 'Y' 表示该产品可回收，而 'N' 表示不可回收。


编写解决方案找出既是低脂又是可回收的产品编号。

返回结果 无顺序要求 。

返回结果格式如下例所示：



示例 1：

输入：
Products 表：
+-------------+----------+------------+
| product_id  | low_fats | recyclable |
+-------------+----------+------------+
| 0           | Y        | N          |
| 1           | Y        | Y          |
| 2           | N        | Y          |
| 3           | Y        | Y          |
| 4           | N        | N          |
+-------------+----------+------------+
输出：
+-------------+
| product_id  |
+-------------+
| 1           |
| 3           |
+-------------+
解释：
只有产品 id 为 1 和 3 的产品，既是低脂又是可回收的产品。

"""

import pandas as pd

def find_products(products: pd.DataFrame) -> pd.DataFrame:
    return products.loc[(products['low_fats'] == 'Y') & (products['recyclable'] == 'Y'),['product_id']]


if __name__ == "__main__":
    lst = [0,1,2,3,4]
    lst1 = ['Y','Y','N','Y','N']
    lst2 = ['N','Y','Y','Y','N']
    df = pd.DataFrame(list(zip(lst,lst1,lst2)),columns=['product_id','low_fats','recyclable'])
    print(find_products(df))

