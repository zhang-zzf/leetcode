import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class CategorizeBoxAccordingToCriteriaTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.categorizeBox(2909, 3968, 3272, 727))
            .isEqualTo("Both");
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public String categorizeBox(int length, int width, int height, int mass) {
            boolean bulky = length >= 10000 ||
                width >= 10000 ||
                height >= 10000 ||
                // watch out: out of int range
                // length * width * height >= 1000000000;
                (long) length * width * height >= 1000000000;
            boolean heavy = mass >= 100;
            if (bulky && heavy) {
                return "Both";
            }
            else if (bulky) {
                return "Bulky";
            }
            else if (heavy) {
                return "Heavy";
            }
            else {
                return "Neither";
            }
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}