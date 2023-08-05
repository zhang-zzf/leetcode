#
# 
# 表： Weather 
# 
# 
#
# 
#+---------------+---------+
#| Column Name   | Type    |
#+---------------+---------+
#| id            | int     |
#| recordDate    | date    |
#| temperature   | int     |
#+---------------+---------+
#在 SQL 中，id 是该表的主键。
#该表包含特定日期的温度信息 
#
# 
#
# 找出与之前（昨天的）日期相比温度更高的所有日期的 id 。 
#
# 返回结果 无顺序要求 。 
#
# 结果格式如下例子所示。 
#
# 
#
# 示例 1： 
#
# 
#输入：
#Weather 表：
#+----+------------+-------------+
#| id | recordDate | Temperature |
#+----+------------+-------------+
#| 1  | 2015-01-01 | 10          |
#| 2  | 2015-01-02 | 25          |
#| 3  | 2015-01-03 | 20          |
#| 4  | 2015-01-04 | 30          |
#+----+------------+-------------+
#输出：
#+----+
#| id |
#+----+
#| 2  |
#| 4  |
#+----+
#解释：
#2015-01-02 的温度比前一天高（10 -> 25）
#2015-01-04 的温度比前一天高（20 -> 30） 
#
#
# Related Topics 数据库 
# 👍 467 👎 0

Create table If Not Exists t_20230805_2
(
    id          int,
    recordDate  date,
    temperature int,
    primary key (id),
    index idx_record_date (recordDate)
);
Truncate table t_20230805_2;
insert into t_20230805_2 (id, recordDate, temperature)
values ('1', '2015-01-01', '10');
insert into t_20230805_2 (id, recordDate, temperature)
values ('2', '2015-01-02', '25');
insert into t_20230805_2 (id, recordDate, temperature)
values ('3', '2015-01-03', '20');
insert into t_20230805_2 (id, recordDate, temperature)
values ('4', '2015-01-04', '30');

select *
from t_20230805_2;

select id,
       recordDate,
       adddate(recordDate, interval -1 day) as yestoday,
       temperature
from t_20230805_2;

select t2.id
from t_20230805_2 as t1
         left join (select id,

                           adddate(recordDate, interval -1 day) as yestoday,
                           temperature
                    from t_20230805_2) as t2 on t1.recordDate = t2.yestoday
where t1.temperature < t2.temperature;

# Equal
select *
from t_20230805_2 w1
         left join t_20230805_2 w2
                   on w1.recordDate = DATE_ADD(w2.recordDate, INTERVAL 1 DAY)
where w1.Temperature > w2.Temperature;

#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below
select t2.id
from Weather as t1
         left join (select id,
                           adddate(recordDate, interval -1 day) as yestoday,
                           temperature
                    from Weather) as t2 on t1.recordDate = t2.yestoday
where t1.temperature < t2.temperature;

#leetcode submit region end(Prohibit modification and deletion)
