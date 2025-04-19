import org.junit.jupiter.api.Test;


class ConvertTheTemperatureTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public double[] convertTemperature(double celsius) {
            return new double[]{celsius + 273.15, celsius * 1.80 + 32.00};
        }
    }
//leetcode submit region end(Prohibit modification and deletion)


}