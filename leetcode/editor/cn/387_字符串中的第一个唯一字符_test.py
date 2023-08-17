import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertEqual(2, self.solution.firstUniqChar("loveleetcode"))


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def firstUniqChar(self, s: str) -> int:
        ans = -1
        cnt_map = [0] * 26
        for x in s:
            cnt_map[ord(x) - ord('a')] += 1
        for idx, x in enumerate(s):
            if cnt_map[ord(x) - ord('a')] == 1:
                ans = idx
                break
        return ans
    # leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
