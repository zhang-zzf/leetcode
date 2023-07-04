import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        # given
        prefix_of_word = self.solution.isPrefixOfWord('i love eating burger',
                                                      'burge')
        # when
        self.assertEqual(4, prefix_of_word)


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def isPrefixOfWord(self, sentence: str, searchWord: str) -> int:
        ans: int = -1
        words: list[str] = sentence.split(" ")
        for idx in range(len(words)):
            if words[idx].startswith(searchWord):
                ans = idx + 1
                break
        return ans


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
