import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class DetermineIfTwoEventsHaveConflictTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.haveConflict(new String[]{"01:15", "02:00"}, new String[]{"02:00", "03:00"})).isTrue();
        then(solution.haveConflict(new String[]{"01:00", "02:00"}, new String[]{"01:20", "03:00"})).isTrue();
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public boolean haveConflict(String[] event1, String[] event2) {
            return event1[0].compareTo(event2[1]) <= 0 && event2[0].compareTo(event1[1]) <= 0;
        }

        private boolean v1(String[] event1, String[] event2) {
            int[] e1 = new int[]{toInstant(event1[0]), toInstant(event1[1])};
            int[] e2 = new int[]{toInstant(event2[0]), toInstant(event2[1])};
            return Math.min(e1[1], e2[1]) >= Math.max(e1[0], e2[0]);
        }

        private int toInstant(String s) {
            String[] split = s.split(":");
            return 60 * Integer.parseInt(split[0]) + Integer.parseInt(split[1]);
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}