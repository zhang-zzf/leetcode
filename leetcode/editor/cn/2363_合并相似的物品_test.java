import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.List;


class MergeSimilarItemsTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public List<List<Integer>> mergeSimilarItems(int[][] items1, int[][] items2) {
            int[] cnt = new int[1001];
            for (int[] item : items1) {
                cnt[item[0]] += item[1];
            }
            for (int[] item : items2) {
                cnt[item[0]] += item[1];
            }
            List<List<Integer>> ans = new ArrayList<>();
            for (int i = 0; i < cnt.length; i++) {
                if (cnt[i] > 0) {
                    ans.add(List.of(i, cnt[i]));
                }
            }
            return ans;
        }
    }
//leetcode submit region end(Prohibit modification and deletion)


}