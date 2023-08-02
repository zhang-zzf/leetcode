import unittest
from functools import reduce


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        ans = self.solution.titleToNumber("Z")
        self.assertEqual(26, ans)


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def titleToNumber(self, columnTitle: str) -> int:
        # TODO ord(str) => 'a' -> 97 / chr(int) => 97 -> 'a'
        def func(a: int, c: str) -> int:
            return a * 26 + ord(c) - ord('A') + 1

        # 必须传入 initial 参数
        return reduce(func, columnTitle, 0)


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
