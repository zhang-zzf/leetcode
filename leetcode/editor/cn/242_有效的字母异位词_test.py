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
        for x in range(ord('a'), ord('z') + 1):
            if s.count(chr(x)) != t.count(chr(x)):
                return False
        return True


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
