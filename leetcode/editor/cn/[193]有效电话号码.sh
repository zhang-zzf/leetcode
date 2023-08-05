#!/usr/bin/env bash
#给定一个包含电话号码列表（一行一个电话号码）的文本文件 file.txt，写一个单行 bash 脚本输出所有有效的电话号码。
#
# 你可以假设一个有效的电话号码必须满足以下两种格式： (xxx) xxx-xxxx 或 xxx-xxx-xxxx。（x 表示一个数字） 
#
# 你也可以假设每行前后没有多余的空格字符。 
#
# 
#
# 示例： 
#
# 假设 file.txt 内容如下： 
#
# 
#987-123-4567
#123 456 7890
#(123) 456-7890
# 
#
# 你的脚本应当输出下列有效的电话号码： 
#
# 
#987-123-4567
#(123) 456-7890
# 
#
# Related Topics Shell 
# 👍 148 👎 0


#leetcode submit region begin(Prohibit modification and deletion)
# Read from the file file.txt and output all valid phone numbers to stdout.
grep -P "^\(\d{3}\)\s\d{3}-\d{4}$|^\d{3}-\d{3}-\d{4}$" file.txt

#leetcode submit region end(Prohibit modification and deletion)
