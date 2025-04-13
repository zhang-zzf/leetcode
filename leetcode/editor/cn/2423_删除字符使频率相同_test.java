import static org.assertj.core.api.BDDAssertions.then;

import java.util.HashSet;
import java.util.Set;
import org.junit.jupiter.api.Test;


class RemoveLetterToEqualizeFrequencyTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.equalFrequency("abcc")).isTrue();
        then(solution.equalFrequency("aazz")).isFalse();
    }

    @Test
    void givenFailedCase1_when_thenSuccess() {
        then(solution.equalFrequency("bac")).isTrue();
    }

    /**
     * 题目理解错误
     */
    @Test
    void givenFailedCase2_when_thenSuccess() {
        then(solution.equalFrequency("bbacc")).isTrue();
    }


    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public boolean equalFrequency(String word) {
            int[] freq = new int[26];
            for (char c : word.toCharArray()) {
                freq[c - 'a'] += 1;
            }
            // 暴力破解
            for (int i = 0; i < freq.length; i++) {
                if (freq[i] == 0) {
                    continue;
                }
                freq[i] -= 1;
                if (leftIsSame(freq)) {
                    return true;
                }
                // 还原
                freq[i] += 1;
            }
            return false;
        }

        private boolean leftIsSame(int[] freq) {
            Set<Integer> set = new HashSet<>();
            for (int f : freq) {
                if (f == 0) {
                    continue;
                }
                set.add(f);
                if (set.size() > 1) {
                    return false;
                }
            }
            return true;
        }

        private static boolean failedCase2(String word) {
            int[] freq = new int[26];
            for (char c : word.toCharArray()) {
                freq[c - 'a'] += 1;
            }
            // 正确思路：把最大值-1，判断剩下的不等于0的值是否相同
            int max = 0;
            for (int f : freq) {
                max = Math.max(f, max);
            }
            for (int i = 0; i < freq.length; i++) {
                if (freq[i] == max) {
                    freq[i] -= 1;
                    // 只把1个最大值-1
                    break;
                }
            }
            // 判断剩下的不等于0的值是否相同
            int same = 0;
            for (int f : freq) {
                if (f == 0) {
                    continue;
                }
                if (same == 0) {
                    same = f;
                }
                else if (same != f) {
                    return false;
                }
            }
            return true;
        }

        private static boolean failedCase1(String word) {
            int[] freq = new int[26];
            for (char c : word.toCharArray()) {
                freq[c - 'a'] += 1;
            }
            // 思路没有找对方向
            // 正确思路：把最大值-1，判断剩下的不等于0的值是否相同
            int max = 0;
            for (int f : freq) {
                max = Math.max(f, max);
            }
            // 最大值只能有1个
            int maxCnt = 0;
            for (int f : freq) {
                if (f == 0 || (f == max && maxCnt++ == 0) || f == max - 1) {
                    continue;
                }
                return false;
            }
            return true;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}