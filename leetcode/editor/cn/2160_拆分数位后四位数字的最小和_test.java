import java.util.PriorityQueue;
import org.junit.jupiter.api.Test;
import static org.assertj.core.api.BDDAssertions.then;


class MinimumSumOfFourDigitNumberAfterSplittingDigitsTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.minimumSum(2932)).isEqualTo(52);
    }
    
    //leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public int minimumSum(int num) {
        PriorityQueue<Integer> queue = new PriorityQueue<>();
        while (num > 0) {
            queue.add(num % 10);
            num /= 10;
        }
        return (queue.poll() + queue.poll()) * 10 + queue.poll() + queue.poll();
    }
}
//leetcode submit region end(Prohibit modification and deletion)

    
}