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
        if len(s) != len(t):
            return False
        cnt_map: dict[str, int] = {}
        for c in s:
            cnt = cnt_map.get(c)
            cnt_map[c] = 1 if cnt is None else cnt + 1
        for c in t:
            cnt = cnt_map.get(c)
            if cnt is None or cnt - 1 < 0:
                return False
            if cnt - 1 == 0:
                del cnt_map[c]
            else:
                cnt_map[c] = cnt - 1
        return len(cnt_map) == 0


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
