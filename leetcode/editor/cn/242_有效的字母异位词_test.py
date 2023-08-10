import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertTrue(self.solution.isAnagram('anagram', 'nagaram'))
        self.assertFalse(self.solution.isAnagram('anagram', 'nagarm'))


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        cnt = [0] * 26
        for c in s:
            cnt[ord(c) - ord('a')] += 1
        for c in t:
            cnt[ord(c) - ord('a')] -= 1
        for c in cnt:
            if c != 0:
                return False
        return True


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
