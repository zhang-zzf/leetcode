import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class CountDaysSpentTogetherTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.countDaysTogether("08-15", "08-18", "08-16", "08-19"))
            .isEqualTo(3);
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        final int[] daysOfMonth = new int[]{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};

        public int countDaysTogether(String arriveAlice, String leaveAlice, String arriveBob, String leaveBob) {
            int[] alice = new int[]{getDays(arriveAlice), getDays(leaveAlice)};
            int[] bob = new int[]{getDays(arriveBob), getDays(leaveBob)};
            return Math.max(Math.min(alice[1], bob[1]) - Math.max(alice[0], bob[0]) + 1, 0);
        }

        /**
         * 转换成天数
         */
        private int getDays(String mmdd) {
            String[] split = mmdd.split("-");
            int month = Integer.parseInt(split[0]);
            int days = Integer.parseInt(split[1]);
            while (month-- > 0) {
                days += daysOfMonth[month];
            }
            return days;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}