import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.in;
import static org.assertj.core.api.BDDAssertions.then;


class MinimumHoursOfTrainingToWinACompetitionTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.minNumberOfHours(5, 3, new int[]{1, 4, 3, 2}, new int[]{2, 6, 3, 1}))
                .isEqualTo(8);
    }

    @Test
    void givenFailedCase1_when_thenSuccess() {
        then(solution.minNumberOfHours(1, 1, new int[]{1, 1, 1, 1}, new int[]{1, 1, 1, 50}))
                .isEqualTo(51);
    }

    @Test
    void givenFailedCase2_when_thenSuccess() {
        then(solution.minNumberOfHours(5, 1, new int[]{1, 3, 3}, new int[]{1, 3, 7}))
                .isEqualTo(6);
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int minNumberOfHours(int initialEnergy, int initialExperience, int[] energy, int[] experience) {
            int ans = 0;
            int totalEnergy = 1;
            for (int e : energy) {
                totalEnergy += e;
            }
            ans += Math.max(0, totalEnergy - initialEnergy);
            int curExperience = initialExperience;
            for (int i = 0; i < experience.length; i++) {
                if (curExperience > experience[i]) {
                    curExperience += experience[i];
                } else {
                    // 打败当前的对手需要的练习时长
                    int neededExperience = experience[i] + 1 - curExperience;
                    curExperience = curExperience + neededExperience + experience[i];
                    ans += neededExperience;
                }
            }
            return ans;
        }

        private int failedCase2(int initialEnergy, int initialExperience, int[] energy, int[] experience) {
            int totalEnergy = 1;
            for (int e : energy) {
                totalEnergy += e;
            }
            int neededExperience = 0;
            for (int i = 0; i < experience.length; i++) {
                if (initialExperience > experience[i]) {
                    initialExperience += experience[i];
                } else {
                    neededExperience += experience[i] + 1 - initialExperience;
                    // 错误的算法
                    initialExperience = initialExperience + neededExperience + experience[i];
                }
            }
            return Math.max(0, totalEnergy - initialEnergy) + neededExperience;
        }
        private int failedCase1(int initialEnergy, int initialExperience, int[] energy, int[] experience) {
            int totalEnergy = 1;
            for (int e : energy) {
                totalEnergy += e;
            }
            int neededExperience = 0;
            for (int i = 0; i < experience.length; i++) {
                if (initialExperience > experience[i]) {
                    initialExperience += experience[i];
                } else {
                    neededExperience += experience[i] + 1 - initialExperience;
                    // 错误的算法
                    initialExperience = experience[i] + 1;
                }
            }
            return Math.max(0, totalEnergy - initialEnergy) + neededExperience;
        }
    }
//leetcode submit region end(Prohibit modification and deletion)


}