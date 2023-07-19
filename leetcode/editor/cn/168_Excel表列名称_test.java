//给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。 
//
// 例如： 
//
// 
//A -> 1
//B -> 2
//C -> 3
//...
//Z -> 26
//AA -> 27
//AB -> 28 
//...
// 
//
// 
//
// 示例 1： 
//
// 
//输入：columnNumber = 1
//输出："A"
// 
//
// 示例 2： 
//
// 
//输入：columnNumber = 28
//输出："AB"
// 
//
// 示例 3： 
//
// 
//输入：columnNumber = 701
//输出："ZY"
// 
//
// 示例 4： 
//
// 
//输入：columnNumber = 2147483647
//输出："FXSHRXW"
// 
//
// 
//
// 提示： 
//
// 
// 1 <= columnNumber <= 2³¹ - 1 
// 
//
// Related Topics 数学 字符串 
// 👍 636 👎 0


import org.junit.jupiter.api.Test;

import static org.assertj.core.api.BDDAssertions.then;


class ExcelSheetColumnTitleTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.convertToTitle(2147483647)).isEqualTo("FXSHRXW");
        then(solution.convertToTitle(701)).isEqualTo("ZY");
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {

        public String convertToTitle(int columnNumber) {
            StringBuilder buf = new StringBuilder();
            final int A = 'A' - 1;
            final int BASE = 26;
            while (columnNumber > 0) {
                int idx = columnNumber % BASE;
                char c = (char) (A + idx);
                columnNumber = columnNumber / BASE;
                // 核心点
                // 0-25 分别对应 Z,A,B,...Y
                if (idx == 0) {
                    c = 'Z';
                    columnNumber -= 1;
                }
                buf.insert(0, c);
            }
            return buf.toString();
        }

    }
//leetcode submit region end(Prohibit modification and deletion)


}