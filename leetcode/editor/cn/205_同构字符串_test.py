import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertTrue(True, "ShouldBeTrue")


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        ans = True
        m1: dict[str, str] = {}
        m2: dict[str, str] = {}
        # todo idx iterate
        for idx in range(len(s)):
            sc, tc = s[idx], t[idx]
            if sc not in m1:
                if (tc not in m2) or (m2[tc] == sc):
                    m1[sc], m2[tc] = tc, sc
                else:
                    ans = False
                    break
            elif m1[sc] != tc:
                ans = False
                break
        return ans


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
