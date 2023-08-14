import math
import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        left, right = 0, 6
        while left <= right:
            mid = left + ((right - left) >> 1)
            print(mid)
            if mid == 1:
                break
            elif mid > 1:
                right = mid - 1
            else:
                left = mid + 1
        ans = self.solution.nthUglyNumber(3, 2, 3, 5)
        self.assertEqual(4, ans)


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def nthUglyNumber(self, n: int, a: int, b: int, c: int) -> int:
        ans = 0
        left, right = 0, min(a, b, c) * n
        ab, ac, bc = math.lcm(a, b), math.lcm(a, c), math.lcm(b, c)
        abc = math.lcm(ab, c)
        while left <= right:
            # TODO >> 的优先级比 + 低
            mid = left + ((right - left) >> 1)
            ugly_num = mid // a + mid // b + mid // c
            ugly_num = ugly_num - mid // ab - mid // ac - mid // bc
            ugly_num += mid // abc
            if ugly_num == n:
                if mid % a == 0 or mid % b == 0 or mid % c == 0:
                    ans = mid
                    break
                else:
                    right = mid - 1
            elif ugly_num > n:
                right = mid - 1
            else:
                left = mid + 1
        return ans


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
