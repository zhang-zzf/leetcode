#
# 
#
# 表: Person 
#
# 
#+-------------+---------+
#| Column Name | Type    |
#+-------------+---------+
#| id          | int     |
#| email       | varchar |
#+-------------+---------+
#id 是该表的主键列。
#此表的每一行都包含一封电子邮件。电子邮件不包含大写字母。
# 
#
# 
#
# 编写一个 SQL 查询来报告所有重复的电子邮件。 请注意，可以保证电子邮件字段不为 NULL。 
#
# 以 任意顺序 返回结果表。 
#
# 查询结果格式如下例。 
#
# 
#
# 示例 1: 
#
# 
#输入: 
#Person 表:
#+----+---------+
#| id | email   |
#+----+---------+
#| 1  | a@b.com |
#| 2  | c@d.com |
#| 3  | a@b.com |
#+----+---------+
#输出: 
#+---------+
#| Email   |
#+---------+
#| a@b.com |
#+---------+
#解释: a@b.com 出现了两次。 
#
# Related Topics 数据库 
# 👍 461 👎 0

create table t_20230803_2
(
    id    int,
    email varchar(128),
    primary key (id)
);

insert into t_20230803_2
values (1, 'a'),
       (2, 'b'),
       (3, 'a');
#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below
select email as Email
from (select email, count(email) as cnt
      from t_20230803_2
      group by email) a
where a.cnt > 1

#leetcode submit region end(Prohibit modification and deletion)
