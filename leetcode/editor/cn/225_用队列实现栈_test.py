import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = MyStack()

    def test_givenNormal_when_thenSuccess(self):
        self.assertTrue(True, "ShouldBeTrue")


# leetcode submit region begin(Prohibit modification and deletion)
class MyStack:

    def __init__(self):
        self.high = []
        self.low = []

    def push(self, x: int) -> None:
        self.high.append(x)
        if len(self.high) > 1:
            self.low.append(self.high.pop(0))

    def pop(self) -> int:
        ans = self.high.pop(0)
        self.high, self.low = self.low, self.high
        while len(self.high) > 1:
            self.low.append(self.high.pop(0))
        return ans
        pass

    def top(self) -> int:
        return self.high[0]

    def empty(self) -> bool:
        return len(self.high) == 0


# Your MyStack object will be instantiated and called as such:
# obj = MyStack()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.top()
# param_4 = obj.empty()
# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
