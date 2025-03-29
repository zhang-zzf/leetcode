import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import org.junit.jupiter.api.Test;


class FindResultantArrayAfterRemovingAnagramsTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public List<String> removeAnagrams(String[] words) {
            List<String> ans = new ArrayList<>();
            String lastSorted = null;
            for (int i = 0; i < words.length; i++) {
                String word = words[i];
                String sorted = sortString(word);
                if (!sorted.equals(lastSorted)) {
                    ans.add(word);
                    lastSorted = sorted;
                }
            }
            return ans;
        }

        private static String sortString(String word) {
            char[] charArray = word.toCharArray();
            Arrays.sort(charArray);
            return new String(charArray);
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}