import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        ans = self.solution.divisorSubstrings(430043, 2)
        self.assertEqual(2, ans)


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def divisorSubstrings(self, num: int, k: int) -> int:
        ans = 0
        # TODO num to string
        # TODO string to num
        # TODO string as slice
        num_str: str = str(num)
        for idx in range(len(num_str) - k + 1):
            val: int = int(num_str[idx:idx + k])
            if val == 0:
                continue
            if num % val == 0:
                ans += 1
        return ans


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
