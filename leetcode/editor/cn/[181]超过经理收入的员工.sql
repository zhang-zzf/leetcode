#表：Employee 
#
# 
#+-------------+---------+
#| Column Name | Type    |
#+-------------+---------+
#| id          | int     |
#| name        | varchar |
#| salary      | int     |
#| managerId   | int     |
#+-------------+---------+
#Id是该表的主键。
#该表的每一行都表示雇员的ID、姓名、工资和经理的ID。
# 
#
# 
#
# 编写一个SQL查询来查找收入比经理高的员工。 
#
# 以 任意顺序 返回结果表。 
#
# 查询结果格式如下所示。 
#
# 
#
# 示例 1: 
#
# 
#输入: 
#Employee 表:
#+----+-------+--------+-----------+
#| id | name  | salary | managerId |
#+----+-------+--------+-----------+
#| 1  | Joe   | 70000  | 3         |
#| 2  | Henry | 80000  | 4         |
#| 3  | Sam   | 60000  | Null      |
#| 4  | Max   | 90000  | Null      |
#+----+-------+--------+-----------+
#输出: 
#+----------+
#| Employee |
#+----------+
#| Joe      |
#+----------+
#解释: Joe 是唯一挣得比经理多的雇员。 
#
# Related Topics 数据库 
# 👍 634 👎 0

create table t_20230803_1
(
    id        int,
    name      varchar(128),
    salary    int,
    managerId int,
    primary key (id)
);

insert into t_20230803_1
values (1, 'Joe', 7000, 3),
       (2, 'Henry', 8000, 4),
       (3, 'Sam', 6000, Null),
       (4, 'Max', 9000, Null);


#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below
select E1.name as Employee
from t_20230803_1 E1
         left join t_20230803_1 E2 on E1.managerId = E2.id
where E1.salary > E2.salary;
#leetcode submit region end(Prohibit modification and deletion)
