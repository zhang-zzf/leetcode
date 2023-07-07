import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertEqual(self.solution.calculate('AB'), 4)


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def calculate(self, s: str) -> int:
        x, y = 1, 0
        for op in s:
            if op == 'A':
                x = x * 2 + y
            else:
                # op == 'B'
                y = y * 2 + x
        return x + y


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
