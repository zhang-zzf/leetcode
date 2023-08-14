import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = MyQueue()

    def test_givenNormal_when_thenSuccess(self):
        self.assertTrue(True, "ShouldBeTrue")


# leetcode submit region begin(Prohibit modification and deletion)
class MyQueue:

    def __init__(self):
        self.tail = []
        self.head = []

    def push(self, x: int) -> None:
        self.tail.append(x)

    def pop(self) -> int:
        self.tail_to_head()
        return self.head.pop()

    def tail_to_head(self):
        if len(self.head) == 0:
            while len(self.tail) > 0:
                self.head.append(self.tail.pop())

    def peek(self) -> int:
        self.tail_to_head()
        return self.head[-1]

    def empty(self) -> bool:
        self.tail_to_head()
        return len(self.head) == 0


# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()
# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
