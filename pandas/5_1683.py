"""
1683.
无效的推文
简单
37
相关企业
SQL
Schema
Pandas
Schema
表：Tweets

+----------------+---------+
| Column
Name | Type |
+----------------+---------+
| tweet_id | int |
| content | varchar |
+----------------+---------+
在
SQL
中，tweet_id
是这个表的主键。
这个表包含某社交媒体
App
中所有的推文。


查询所有无效推文的编号（ID）。当推文内容中的字符数严格大于
15
时，该推文是无效的。

以任意顺序返回结果表。

查询结果格式如下所示：



示例
1：

输入：
Tweets
表：
+----------+----------------------------------+
| tweet_id | content |
+----------+----------------------------------+
| 1 | Vote
for Biden |
    | 2 | Let
    us
    make
    America
    great
    again! |
    +----------+----------------------------------+

    输出：
    +----------+
    | tweet_id |
    +----------+
    | 2 |
    +----------+
    解释：
    推文
    1
    的长度
    length = 14。该推文是有效的。
    推文
    2
    的长度
    length = 32。该推文是无效的。
"""


import pandas as pd


def invalid_tweets(tweets: pd.DataFrame) -> pd.DataFrame:
    is_valid = tweets['content'].str.len() > 15
    df = tweets[is_valid]
    df = df[['tweet_id']]
    return df


