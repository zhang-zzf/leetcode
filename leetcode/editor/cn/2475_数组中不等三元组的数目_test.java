import org.junit.jupiter.api.Test;

import static org.assertj.core.api.BDDAssertions.then;


class NumberOfUnequalTripletsInArrayTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.unequalTriplets(new int[]{4, 4, 2, 4, 3})).isEqualTo(3);
    }

    @Test
    void givenFailedCase1_when_thenSuccess() {
        then(solution.unequalTriplets(new int[]{1, 3, 1, 2, 4})).isEqualTo(7);
    }


    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int unequalTriplets(int[] nums) {
            return forceBreak(nums);
        }

        /**
         * 放弃
         */
        private static int failedCase1(int[] nums) {
            int max = 0;
            for (int num : nums) {
                max = Math.max(max, num);
            }
            int[] cnt = new int[max + 1];
            for (int num : nums) {
                cnt[num]++;
            }
            int product = 1, diffCnt = 0;
            for (int c : cnt) {
                if (c == 0) {
                    continue;
                }
                diffCnt += 1;
                product *= c;
            }
            return diffCnt > 2 ? product : 0;
        }

        private static int forceBreak(int[] nums) {
            int ans = 0;
            for (int i = 0; i < nums.length - 2; i++) {
                for (int j = i + 1; j < nums.length - 1; j++) {
                    for (int k = j + 1; k < nums.length; k++) {
                        if (nums[i] != nums[j] && nums[j] != nums[k] && nums[i] != nums[k]) {
                            ans++;
                        }
                    }
                }
            }
            return ans;
        }
    }
//leetcode submit region end(Prohibit modification and deletion)


}