import lombok.extern.slf4j.Slf4j;
import org.junit.jupiter.api.Test;

import java.util.HashSet;
import java.util.Set;

import static java.util.stream.Collectors.toList;


@Slf4j
class HappyNumberTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        Set<String> result = new HashSet<>();
        for (int i = 1; i < Integer.MAX_VALUE; i++) {
            String aCycle = happyCycle(i);
            result.add(aCycle);
        }
        log.info("all cycles -> {}", result);
    }

    public String happyCycle(int n) {
        String ans = "[1]";
        Set<Integer> set = new HashSet<>();
        boolean inCycle = false;
        while (true) {
            n = solution.nextN(n);
            if (n == 1) {
                break;
            }
            if (set.contains(n)) {
                // n 是环内的一个数字
                // clear set
                if (inCycle) {
                    ans = set.stream().sorted().collect(toList()).toString();
                    break;
                }
                inCycle = true;
                set.clear();
            }
            set.add(n);
        }
        return ans;
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {

        public boolean isHappy(int n) {
            boolean ans = false;
            int slow = nextN(n), fast = nextN(nextN(n));
            for (; true; ) {
                int fn = nextN(fast);
                if (slow == 1 || fast == 1 || fn == 1) {
                    ans = true;
                    break;
                }
                if (slow == fast) {
                    // 环
                    break;
                }
                slow = nextN(slow);
                fast = nextN(fn);
            }
            return ans;
        }

        private int nextN(int n) {
            int ans = 0;
            for (; n != 0; ) {
                int num = n % 10;
                ans += num * num;
                n /= 10;
            }
            return ans;
        }

    }
//leetcode submit region end(Prohibit modification and deletion)


}