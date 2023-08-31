"""
570. 至少有5名直接下属的经理
提示
中等
84
相关企业
SQL Schema
Pandas Schema
表: Employee

+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| name        | varchar |
| department  | varchar |
| managerId   | int     |
+-------------+---------+
在 SQL 中，id 是该表的主键列。
该表的每一行都表示雇员的名字、他们的部门和他们的经理的id。
如果managerId为空，则该员工没有经理。
没有员工会成为自己的管理者。


查询至少有5名直接下属的经理 。

以 任意顺序 返回结果表。

查询结果格式如下所示。



示例 1:

输入:
Employee 表:
+-----+-------+------------+-----------+
| id  | name  | department | managerId |
+-----+-------+------------+-----------+
| 101 | John  | A          | None      |
| 102 | Dan   | A          | 101       |
| 103 | James | A          | 101       |
| 104 | Amy   | A          | 101       |
| 105 | Anne  | A          | 101       |
| 106 | Ron   | B          | 101       |
+-----+-------+------------+-----------+
输出:
+------+
| name |
+------+
| John |
+------+
"""

import pandas as pd

def find_managers(employee: pd.DataFrame) -> pd.DataFrame:
    #print(employee)
    df = employee.groupby('managerId').size().reset_index(name='count')

    df = df[df['count'] >= 5]

    df = df.merge(employee, left_on="managerId", right_on="id", how="inner")
    return df[["name"]]

if __name__ == "__main__":
    employee = pd.DataFrame({
        "id": [101,102,103,104,105,106],
        "name": ["John", "Dan", "James", "Amy", "Anne", "Ron"],
        "department": ["A", "A", "A", "A", "A", "B"],
        "managerId": [None, 101, 101, 101, 101, 101]
    })

    df = find_managers(employee)
    print(df)