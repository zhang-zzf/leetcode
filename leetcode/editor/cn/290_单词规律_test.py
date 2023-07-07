import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertTrue(self.solution.wordPattern("abba", "dog cat cat dog"))

    def test_givenFailCase1_when_thenSuccess(self):
        self.assertFalse(self.solution.wordPattern("abba", "dog dog dog dog"))


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def wordPattern(self, pattern: str, s: str) -> bool:
        words: list[str] = s.split(" ")
        if len(pattern) != len(words):
            return False
        ans: bool = True
        p2w: dict[str, str] = {}
        w2p: dict[str, str] = {}
        for i in range(len(pattern)):
            p: str = pattern[i]
            w: str = words[i]
            m1: str = p2w.get(p)
            m2: str = w2p.get(w)
            # python None
            if m1 is None and m2 is None:
                p2w[p] = w
                w2p[w] = p
            elif m1 != w or m2 != p:
                ans = False
                break
        return ans

    # leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
