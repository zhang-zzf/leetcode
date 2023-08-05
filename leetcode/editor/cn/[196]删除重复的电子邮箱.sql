#表: Person 
#
# 
#+-------------+---------+
#| Column Name | Type    |
#+-------------+---------+
#| id          | int     |
#| email       | varchar |
#+-------------+---------+
#在 SQL 中，id 是该表的主键列。
#该表的每一行包含一封电子邮件。电子邮件将不包含大写字母。
# 
#
# 
#
# 删除 所有重复的电子邮件，只保留一个具有最小 id 的唯一电子邮件。 
#
# （对于 SQL 用户，请注意你应该编写一个 DELETE 语句而不是 SELECT 语句。） 
#
# （对于 Pandas 用户，请注意你应该直接修改 Person 表。） 
#
# 运行脚本后，显示的答案是 Person 表。驱动程序将首先编译并运行您的代码片段，然后再显示 Person 表。Person 表的最终顺序 无关紧要 。 
#
# 返回结果格式如下示例所示。 
#
# 
#
# 示例 1: 
#
# 
#输入: 
#Person 表:
#+----+------------------+
#| id | email            |
#+----+------------------+
#| 1  | john@example.com |
#| 2  | bob@example.com  |
#| 3  | john@example.com |
#+----+------------------+
#输出: 
#+----+------------------+
#| id | email            |
#+----+------------------+
#| 1  | john@example.com |
#| 2  | bob@example.com  |
#+----+------------------+
#解释: john@example.com重复两次。我们保留最小的Id = 1。 
#
# Related Topics 数据库 
# 👍 752 👎 0

create table t_20230805_1
(
    id    int,
    email varchar(128),
    primary key (id)
);

alter table t_20230805_1
    add index idx_email (`email`);


insert into t_20230805_1
values (1, 'a'),
       (2, 'b'),
       (3, 'a'),
       (4, 'a');

truncate t_20230805_1;

select *
from t_20230805_1;

select MIN(id)
from t_20230805_1
group by email;

# ERROR
delete
from t_20230805_1
where id not in (select MIN(id)
                 from t_20230805_1
                 group by email);

SELECT p1.*, p2.*
FROM t_20230805_1 p1,
     t_20230805_1 p2
WHERE p1.Email = p2.Email;

SELECT p1.*, p2.*
FROM t_20230805_1 p1,
     t_20230805_1 p2
WHERE p1.Email = p2.Email
  and p1.id > p2.id;


delete p1
from t_20230805_1 p1,
     t_20230805_1 p2
where p1.email = p2.email
  and p1.id > p2.id;

#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below
delete p1
from Person p1,
     Person p2
where p1.email = p2.email
  and p1.id > p2.id

#leetcode submit region end(Prohibit modification and deletion)
