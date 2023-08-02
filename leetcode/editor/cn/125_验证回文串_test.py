import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        panama = "A man, a plan, a canal: Panama"
        is_palindrome = self.solution.isPalindrome(panama)
        self.assertTrue(is_palindrome)
        not_palindrome = self.solution.isPalindrome("race a car")
        self.assertFalse(not_palindrome)


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def isPalindrome(self, s: str) -> bool:
        ans = True
        left, right = 0, len(s) - 1
        while left < right:
            if not s[left].isalnum():
                left += 1
                continue
            if not s[right].isalnum():
                right -= 1
                continue
            if s[left].lower() != s[right].lower():
                ans = False
                break
            # move pointer
            left, right = left + 1, right - 1
        return ans


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
