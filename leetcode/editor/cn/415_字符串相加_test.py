import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertEqual('134', self.solution.addStrings("11", '123'))
        self.assertEqual('0', self.solution.addStrings("0", '0'))

    def test_givenFailedCase1_when_thenSuccess(self):
        self.assertEqual('10', self.solution.addStrings("1", '9'))


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def addStrings(self, num1: str, num2: str) -> str:
        ans: str = ''
        num1 = num1[::-1]
        num2 = num2[::-1]
        if len(num1) >= len(num2):
            num1, num2 = num2, num1
        promotion: int = 0
        for i in range(0, len(num2)):
            n: int = int(num2[i])
            nn: int = 0
            if i < len(num1):
                nn = int(num1[i])
            s: int = n + nn + promotion
            if s < 10:
                promotion = 0
            else:
                promotion = 1
                s -= 10
            ans = str(s) + ans
        if promotion:
            ans = str(promotion) + ans
        return ans


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
