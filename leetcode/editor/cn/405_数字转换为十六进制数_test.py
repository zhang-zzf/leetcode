import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertEqual('1a', self.solution.toHex(26))
        self.assertEqual('0', self.solution.toHex(0))
        self.assertEqual('ffffffff', self.solution.toHex(-1))


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def toHex(self, num: int) -> str:
        if num == 0:
            return '0'
        hex_tab = '0123456789abcdef'
        ans = ''
        for i in range(7, -1, -1):
            ans += hex_tab[num >> (4 * i) & 0x0f]
        # 移除前导 0
        return ans.lstrip('0')

    # 此解法是错误的
    def toHexError(self, num: int) -> str:
        if num == 0:
            return '0'
        hex_tab = '0123456789abcdef'
        ans = ''
        while num:
            # TODO -1 >> 4 的结果依旧是 -1
            ans = hex_tab[num & 0x0f] + ans
            num >>= 4
        return ans


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
