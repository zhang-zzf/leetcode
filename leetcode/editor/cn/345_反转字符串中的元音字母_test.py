import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.solution.reverseVowels("Hello")
        self.assertTrue(True, "ShouldBeTrue")


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def reverseVowels(self, s: str) -> str:
        # TODO str => list[str]
        split: list[str] = list(s)
        vowels = {
            'a', 'e', 'i', 'o', 'u',
            'A', 'E', 'I', 'O', 'U'
        }
        left, right = 0, len(s) - 1
        while left < right:
            if split[left] not in vowels:
                left += 1
                continue
            if split[right] not in vowels:
                right -= 1
                continue
            split[left], split[right] = split[right], split[left],
            left, right = left + 1, right - 1
        # TODO list to str
        return ''.join(split)


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
