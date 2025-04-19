import org.junit.jupiter.api.Test;

import java.util.HashSet;
import java.util.Set;

import static org.assertj.core.api.BDDAssertions.then;


class NumberOfDistinctAveragesTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.distinctAverages(new int[]{4, 1, 4, 0, 3, 5})).isEqualTo(2);
    }

    @Test
    void givenFailedCase1_when_thenSuccess() {
        then(solution.distinctAverages(new int[]{9, 5, 7, 8, 7, 9, 8, 2, 0, 7})).isEqualTo(5);
    }


    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int distinctAverages(int[] nums) {
            int[] cnt = new int[101];
            for (int num : nums) {
                cnt[num] += 1;
            }
            int ans = 0;
            Set<Integer> avgSet = new HashSet<>();
            for (int i = 0, j = cnt.length - 1; i < j; ) {
                while (cnt[i] == 0 && i < j) {
                    i++;
                }
                while (cnt[j] == 0 && i < j) {
                    j--;
                }
                // failed case1
                // if (i == j) {
                //     continue;
                // }
                if (i == j && cnt[i] == 0) {
                    continue;
                }
                int min = Math.min(cnt[i], cnt[j]);
                cnt[i] -= min;
                cnt[j] -= min;
                if (avgSet.add(i + j)) {
                    ans += 1;
                }
            }
            return ans;
        }
    }
//leetcode submit region end(Prohibit modification and deletion)


}