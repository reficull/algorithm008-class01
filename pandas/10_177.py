"""
177. 第N高的薪水
中等
742
相关企业
SQL Schema
表: Employee

+-------------+------+
| Column Name | Type |
+-------------+------+
| id          | int  |
| salary      | int  |
+-------------+------+
在 SQL 中，id 是该表的主键。
该表的每一行都包含有关员工工资的信息。


查询 Employee 表中第 n 高的工资。如果没有第 n 个最高工资，查询结果应该为 null 。

查询结果格式如下所示。



示例 1:

输入:
Employee table:
+----+--------+
| id | salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+
n = 2
输出:
+------------------------+
| getNthHighestSalary(2) |
+------------------------+
| 200                    |
+------------------------+
示例 2:

输入:
Employee 表:
+----+--------+
| id | salary |
+----+--------+
| 1  | 100    |
+----+--------+
n = 2
输出:
+------------------------+
| getNthHighestSalary(2) |
+------------------------+
| null                   |
+------------------------+
"""

import pandas as pd


def nth_highest_salary(employee: pd.DataFrame, N: int) -> pd.DataFrame:
    df = employee[["salary"]].drop_duplicates()
    if N > len(df):
        return pd.DataFrame({'getNthHighestSalary(2)': [None]})
    return df.sort_values("salary", ascending=False).head(N).tail(1)
