// 一家商店正在打折销售糖果。每购买 两个 糖果，商店会 免费 送一个糖果。
//
// 免费送的糖果唯一的限制是：它的价格需要小于等于购买的两个糖果价格的 较小值 。 
//
// 
// 比方说，总共有 4 个糖果，价格分别为 1 ，2 ，3 和 4 ，一位顾客买了价格为 2 和 3 的糖果，那么他可以免费获得价格为 1 的糖果，但不能获得
// 价格为 4 的糖果。
// 
//
// 给你一个下标从 0 开始的整数数组 cost ，其中 cost[i] 表示第 i 个糖果的价格，请你返回获得 所有 糖果的 最小 总开销。 
//
// 
//
// 示例 1： 
//
// 输入：cost = [1,2,3]
// 输出：5
// 解释：我们购买价格为 2 和 3 的糖果，然后免费获得价格为 1 的糖果。
// 总开销为 2 + 3 = 5 。这是开销最小的 唯一 方案。
// 注意，我们不能购买价格为 1 和 3 的糖果，并免费获得价格为 2 的糖果。
// 这是因为免费糖果的价格必须小于等于购买的 2 个糖果价格的较小值。
// 
//
// 示例 2： 
//
// 输入：cost = [6,5,7,9,2,2]
// 输出：23
// 解释：最小总开销购买糖果方案为：
//- 购买价格为 9 和 7 的糖果
//- 免费获得价格为 6 的糖果
//- 购买价格为 5 和 2 的糖果
//- 免费获得价格为 2 的最后一个糖果
// 因此，最小总开销为 9 + 7 + 5 + 2 = 23 。
// 
//
// 示例 3： 
//
// 输入：cost = [5,5]
// 输出：10
// 解释：由于只有 2 个糖果，我们需要将它们都购买，而且没有免费糖果。
// 所以总最小开销为 5 + 5 = 10 。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= cost.length <= 100 
// 1 <= cost[i] <= 100 
// 
//
// Related Topics 贪心 数组 排序 
// 👍 24 👎 0


import static org.assertj.core.api.BDDAssertions.then;

import java.util.Comparator;
import java.util.PriorityQueue;
import org.junit.jupiter.api.Test;


class MinimumCostOfBuyingCandiesWithDiscountTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.minimumCost(new int[]{1, 2, 3})).isEqualTo(5);
        then(solution.minimumCost(new int[]{6, 5, 7, 9, 2, 2})).isEqualTo(23);
        then(solution.minimumCost(new int[]{5, 5})).isEqualTo(10);
    }

    @Test
    void givenFailedCase1_when_thenSuccess() {
        then(solution.minimumCost(new int[]{10, 5, 9, 4, 1, 9, 10, 2, 10, 8})).isEqualTo(48);
    }


    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        /**
         * 1 <= cost.length <= 100 1 <= cost[i] <= 100 基数算法
         */
        public int minimumCost(int[] cost) {
            return optimizeAfterReferencingAnswer(cost);
        }

        private int optimizeAfterReferencingAnswer(int[] cost) {
            int[] counter = new int[101];
            for (int c : cost) {
                counter[c] += 1;
            }
            int ans = 0;
            for (int j = 0, c = 100; ; ) {
                while (c > 0 && counter[c] == 0) {
                    c -= 1;
                }
                if (c == 0) {
                    break;
                }
                counter[c] -= 1;
                //
                if (j == 2) {
                    j = 0;
                    continue;
                }
                ans += c;
                j += 1;
            }
            return ans;
        }

        private int success1ByMyself(int[] cost) {
            int ans = 0;
            PriorityQueue<Integer> pq = new PriorityQueue<>(Comparator.reverseOrder());
            for (int i : cost) {
                pq.add(i);
            }
            for (int i = 0; pq.size() > 0; ) {
                Integer c = pq.poll();
                if (i == 2) {
                    i = 0;
                    continue;
                }
                ans += c;
                i += 1;
            }
            return ans;
        }

        private int failedCase1(int[] cost) {
            int ans = 0;
            PriorityQueue<Integer> pq = new PriorityQueue<>(Comparator.reverseOrder());
            for (int i : cost) {
                pq.add(i);
            }
            for (int i = 0; pq.size() > 0; i++) {
                Integer c = pq.poll();
                if (i < 2) {
                    ans += c;
                }
                else {
                    i = 0; // 问题出在这里 -> i++
                }
            }
            return ans;
        }
    }
// leetcode submit region end(Prohibit modification and deletion)


}