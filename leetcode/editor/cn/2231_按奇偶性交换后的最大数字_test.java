import static java.util.Optional.ofNullable;
import static org.assertj.core.api.BDDAssertions.then;

import java.util.Optional;
import java.util.PriorityQueue;
import org.junit.jupiter.api.Test;


class LargestNumberAfterDigitSwapsByParityTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.largestInteger(1234)).isEqualTo(3412);
        then(solution.largestInteger(65875)).isEqualTo(87655);
    }

    /**
     * <pre>
     *  Progress： 整数高位到低位遍历
     *    1. 字符串
     *    1. 栈
     *    1. 总长度 + 除法
     *    1. 递归
     * </pre>
     */

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int largestInteger(int num) {
            return initVersion(num);
        }

        private static int initVersion(int num) {
            PriorityQueue<Integer> oddQueue = new PriorityQueue<>();
            PriorityQueue<Integer> evenQueue = new PriorityQueue<>();
            for (int n = num; n > 0; n /= 10) {
                int digit = n % 10;
                if (digit % 2 == 1) {
                    oddQueue.add(digit);
                }
                else {
                    evenQueue.add(digit);
                }
            }
            int ans = 0;
            // 从低位向高位遍历, 如何从高位向低位遍历？
            for (int n = num, cnt = 0; n > 0; n /= 10, cnt++) {
                Integer min = (n % 10) % 2 == 1 ? oddQueue.poll() : evenQueue.poll();
                ans = ans + (int) Math.pow(10, cnt) * Optional.ofNullable(min).orElse(0);
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}