"""
1148.
文章浏览
I
简单
85
相关企业
SQL
Schema
Pandas
Schema
Views
表：

+---------------+---------+
| Column
Name | Type |
+---------------+---------+
| article_id | int |
| author_id | int |
| viewer_id | int |
| view_date | date |
+---------------+---------+
此表可能会存在重复行。（换句话说，在
SQL
中这个表没有主键）
此表的每一行都表示某人在某天浏览了某位作者的某篇文章。
请注意，同一人的
author_id
和
viewer_id
是相同的。


请查询出所有浏览过自己文章的作者

结果按照
id
升序排列。

查询结果的格式如下所示：



示例
1：

输入：
Views
表：
+------------+-----------+-----------+------------+
| article_id | author_id | viewer_id | view_date |
+------------+-----------+-----------+------------+
| 1 | 3 | 5 | 2019 - 0
8 - 01 |
| 1 | 3 | 6 | 2019 - 0
8 - 02 |
| 2 | 7 | 7 | 2019 - 0
8 - 01 |
| 2 | 7 | 6 | 2019 - 0
8 - 02 |
| 4 | 7 | 1 | 2019 - 07 - 22 |
| 3 | 4 | 4 | 2019 - 07 - 21 |
| 3 | 4 | 4 | 2019 - 07 - 21 |
+------------+-----------+-----------+------------+

输出：
+------+
| id |
+------+
| 4 |
| 7 |
+------+
"""

import pandas as pd

def article_views(views: pd.DataFrame) -> pd.DataFrame:
    df = views[views['author_id'] == views['viewer_id']]
    df.drop_duplicates(subset=['author_id'], inplace=True)

    # 按 'author_id' 列升序排序 DataFrame。
    df.sort_values(by=['author_id'], inplace=True)

    # 将 'author_id' 列重命名为 'id'。
    df.rename(columns={'author_id':'id'}, inplace=True)
    df = df[['id']]
    return df