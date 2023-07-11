import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertTrue(self.solution.isValid("()"))
        self.assertFalse(self.solution.isValid("[[]"))
        self.assertFalse(self.solution.isValid("[]]"))
        self.assertFalse(self.solution.isValid("]]]"))


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def isValid(self, s: str) -> bool:
        ans: bool = True
        pairs: dict[str, str] = {
            "(": ")",
            "[": ']',
            "{": "}"
        }
        stack: list[str] = []
        for c in s:
            # watch out: can not use pairs[c]
            if pairs.get(c) is not None:
                stack.append(c)
            elif len(stack) == 0 or c != pairs[stack.pop()]:
                ans = False
                break
        if len(stack) != 0:
            ans = False
        return ans


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
