import org.junit.jupiter.api.Test;


class BestPokerHandTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public String bestHand(int[] ranks, char[] suits) {
            if (flush(suits)) {
                return "Flush";
            }
            int[] cnt = new int[14];
            for (int rank : ranks) {
                cnt[rank] += 1;
            }
            for (int count : cnt) {
                if (count >= 3) {
                    return "Three of a Kind";
                }
            }
            for (int count : cnt) {
                if (count == 2) {
                    return "Pair";
                }
            }
            // for (int rank : ranks) {
            //     cnt[rank] += 1;
            //     // [4,4,2,4,4]
            //     if (cnt[rank] >= 3) {
            //         return "Three of a Kind";
            //     }
            //     if (cnt[rank] >= 2) {
            //         return "Pair";
            //     }
            // }
            return "High Card";
        }

        private boolean flush(char[] suits) {
            char[] cnt = new char[4];
            for (char c : suits) {
                cnt[c - 'a'] += 1;
            }
            for (char c : cnt) {
                if (c >= 5) {
                    return true;
                }
            }
            return false;
        }
    }
//leetcode submit region end(Prohibit modification and deletion)


}