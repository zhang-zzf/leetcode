import static org.assertj.core.api.BDDAssertions.then;

import java.util.Comparator;
import java.util.Optional;
import java.util.PriorityQueue;
import org.junit.jupiter.api.Test;


class SortEvenAndOddIndicesIndependentlyTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.sortEvenOdd(new int[]{4, 1, 2, 3}))
            .isEqualTo(new int[]{2, 3, 4, 1});
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int[] sortEvenOdd(int[] nums) {
            int[] ans = new int[nums.length];
            PriorityQueue<Integer> odd = new PriorityQueue<>(Comparator.reverseOrder());
            PriorityQueue<Integer> even = new PriorityQueue<>();
            for (int i = 0; i < nums.length; i++) {
                if (i % 2 == 0) {
                    even.add(nums[i]);
                }
                else {
                    odd.add(nums[i]);
                }
            }
            for (int i = 0; i < nums.length; i++) {
                Integer num;
                if (i % 2 == 0) {
                    num = even.poll();
                }
                else {
                    num = odd.poll();
                }
                ans[i] = Optional.ofNullable(num).orElse(0);
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}