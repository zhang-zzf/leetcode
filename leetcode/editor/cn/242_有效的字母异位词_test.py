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
        s_cnt: dict[str, int] = {}
        t_cnt: dict[str, int] = {}
        for c in s:
            cnt = s_cnt.get(c)
            s_cnt[c] = 1 if cnt is None else cnt + 1
        for c in t:
            cnt = t_cnt.get(c)
            t_cnt[c] = 1 if cnt is None else cnt + 1
        if len(s_cnt) != len(t_cnt):
            return False
        ans = True
        for k, v in s_cnt.items():
            if t_cnt.get(k) != v:
                ans = False
                break
        return ans


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
