import java.util.Arrays;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;
import org.junit.jupiter.api.Test;
import static org.assertj.core.api.BDDAssertions.then;


class FindTheDifferenceOfTwoArraysTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
       
    }
    
    //leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public List<List<Integer>> findDifference(int[] nums1, int[] nums2) {
        Set<Integer> num1Set = Arrays.stream(nums1).boxed().collect(Collectors.toSet());
        Set<Integer> num2Set = Arrays.stream(nums2).boxed().collect(Collectors.toSet());
        return Arrays.asList(
                num1Set.stream().filter(n -> !num2Set.contains(n)).collect(Collectors.toList()),
                num2Set.stream().filter(n -> !num1Set.contains(n)).collect(Collectors.toList())
        );
    }
}
//leetcode submit region end(Prohibit modification and deletion)

    
}