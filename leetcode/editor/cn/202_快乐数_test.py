import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        is_happy = self.solution.isHappy(19)
        self.assertTrue(is_happy)


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def isHappy(self, n: int) -> bool:
        ans: bool = False

        def next_n(a_num: int) -> int:
            result: int = 0
            while a_num:
                a_num, mod = divmod(a_num, 10)
                result += mod * mod
            return result

        slow, fast = next_n(n), next_n(next_n(n)),
        while True:
            fn = next_n(fast)
            if slow == 1 or fast == 1 or fn == 1:
                ans = True
                break
            if slow == fast:
                break
            slow, fast = next_n(slow), next_n(fn)
        return ans


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
