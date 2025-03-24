import static org.assertj.core.api.BDDAssertions.then;

import java.util.ArrayList;
import java.util.List;
import org.junit.jupiter.api.Test;


class CellsInARangeOnAnExcelSheetTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.cellsInRange("K1:L2")).contains("K1", "K2", "L1", "L2");
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public List<String> cellsInRange(String s) {
            List<String> ans = new ArrayList<>();
            for (char c = s.charAt(0); c <= s.charAt(3); c++) {
                for (char r = s.charAt(1); r <= s.charAt(4); r++) {
                    ans.add(String.valueOf(c) + r);
                }
            }
            return ans;
        }

        private static List<String> v1(String s) {
            List<String> ans = new ArrayList<>();
            for (char c = s.charAt(0); c <= s.charAt(3); c++) {
                // 为何要转 int
                for (int r = s.charAt(1) - '0'; r <= s.charAt(4) - '0'; r++) {
                    ans.add(c + "" + r);
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}