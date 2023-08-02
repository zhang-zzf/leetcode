//给你一个字符串 columnTitle ，表示 Excel 表格中的列名称。返回 该列名称对应的列序号 。 
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
// 示例 1: 
//
// 
//输入: columnTitle = "A"
//输出: 1
// 
//
// 示例 2: 
//
// 
//输入: columnTitle = "AB"
//输出: 28
// 
//
// 示例 3: 
//
// 
//输入: columnTitle = "ZY"
//输出: 701 
//
// 
//
// 提示： 
//
// 
// 1 <= columnTitle.length <= 7 
// columnTitle 仅由大写英文组成 
// columnTitle 在范围 ["A", "FXSHRXW"] 内 
// 
//
// Related Topics 数学 字符串 
// 👍 390 👎 0


import org.junit.jupiter.api.Test;

import static org.assertj.core.api.BDDAssertions.then;


class ExcelSheetColumnNumberTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.titleToNumber("ZY")).isEqualTo(701);
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {

        public int titleToNumber(String columnTitle) {
            int ans = 0;
            int length = columnTitle.length();
            for (int i = length - 1; i >= 0; i--) {
                ans += (columnTitle.charAt(i) - 'A' + 1) * Math.pow(26, length - 1 - i);
            }
            return ans;
        }

    }
//leetcode submit region end(Prohibit modification and deletion)


}