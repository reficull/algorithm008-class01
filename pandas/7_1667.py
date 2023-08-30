"""
1667. 修复表中的名字
简单
106
相关企业
SQL Schema
Pandas Schema
表： Users

+----------------+---------+
| Column Name    | Type    |
+----------------+---------+
| user_id        | int     |
| name           | varchar |
+----------------+---------+
user_id 是该表的主键(具有唯一值的列)。
该表包含用户的 ID 和名字。名字仅由小写和大写字符组成。


编写解决方案，修复名字，使得只有第一个字符是大写的，其余都是小写的。

返回按 user_id 排序的结果表。

返回结果格式示例如下。



示例 1：

输入：
Users table:
+---------+-------+
| user_id | name  |
+---------+-------+
| 1       | aLice |
| 2       | bOB   |
+---------+-------+
输出：
+---------+-------+
| user_id | name  |
+---------+-------+
| 1       | Alice |
| 2       | Bob   |
+---------+-------+
"""
import pandas as pd

def fix_names(users: pd.DataFrame) -> pd.DataFrame:
    # 取第一个字符转大写， 加上后面字符 转小写

    #users["name"] = users["name"].str[0].str.upper() + users["name"].str[1:].str.lower()
    # title 方法只存在于 pandas
    users["name"] = users["name"].str.title()
    return users.sort_values("user_id")




