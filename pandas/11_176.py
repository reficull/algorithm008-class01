"""
176. 第二高的薪水
中等
1.4K
相关企业
SQL Schema
Pandas Schema
Employee 表：
+-------------+------+
| Column Name | Type |
+-------------+------+
| id          | int  |
| salary      | int  |
+-------------+------+
在 SQL 中，id 是这个表的主键。
表的每一行包含员工的工资信息。


查询并返回 Employee 表中第二高的薪水 。如果不存在第二高的薪水，查询应该返回 null(Pandas 则返回 None) 。

查询结果如下例所示。



示例 1：

输入：
Employee 表：
+----+--------+
| id | salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+
输出：
+---------------------+
| SecondHighestSalary |
+---------------------+
| 200                 |
+---------------------+
示例 2：

输入：
Employee 表：
+----+--------+
| id | salary |
+----+--------+
| 1  | 100    |
+----+--------+
输出：
+---------------------+
| SecondHighestSalary |
+---------------------+
| null                |
+---------------------+
"""
import pandas as pd


def second_highest_salary(employee: pd.DataFrame) -> pd.DataFrame:
    employee = employee.drop_duplicates(["salary"])
    if len(employee["salary"].unique()) < 2:
        return pd.DataFrame({'SecondHighestSalary': [np.NaN]})
    employee = employee.sort_values("salary", ascending=False).head(2).tail(1)
    employee.drop("id", axis=1, inplace=True)
    employee.rename({"salary": "SecondHighestSalary"}, axis=1, inplace=True)
    return employee.head(2).tail(1)
