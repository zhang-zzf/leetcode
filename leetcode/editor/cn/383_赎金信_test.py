import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertTrue(self.solution.canConstruct('aa', 'aaa'), "ShouldBeTrue")

    def test_givenFailCase1_when_thenSuccess(self):
        self.assertFalse(self.solution.canConstruct('a', 'b'))


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        ans = True
        mapping = {}
        for c in magazine:
            if mapping.get(c) is None:
                mapping[c] = 1
            else:
                mapping[c] += 1
        for c in ransomNote:
            cnt = mapping.get(c)
            if cnt is None or cnt < 1:
                ans = False
                break
            mapping[c] -= 1
        return ans


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
