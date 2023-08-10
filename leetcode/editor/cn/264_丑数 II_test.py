import queue
import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertEqual(12, self.solution.nthUglyNumber(10))


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def nthUglyNumber(self, n: int) -> int:
        dp = [0, 1]
        pq = queue.PriorityQueue(maxsize=3)
        pq.put((2, [1, 2]))
        pq.put((3, [1, 3]))
        pq.put((5, [1, 5]))
        x = 2
        while x <= n:
            ugly, item = pq.get()
            if ugly != dp[-1]:
                dp.append(ugly)
                x += 1
            item[0] += 1
            pq.put((dp[item[0]] * item[1], item))
        return dp[n]


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
