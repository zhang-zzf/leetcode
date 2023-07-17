from __future__ import annotations

import queue
import unittest
from typing import List

A = "A"


class TestCommonCollections(unittest.TestCase):

    def setUp(self) -> None:
        pass

    # Stack
    # LIFO
    def test_givenStack_when_thenSuccess(self):
        stack: List[str | None] = []
        stack.append(A)
        stack.append(None)
        while len(stack) > 0:
            last = stack.pop()
            self.assertIn(last, [None, A])

    # Queue
    # FIFO
    def test_givenQueue_when_thenSuccess(self):
        queue: List[str | None] = []
        queue.append(A)
        queue.append(None)
        while len(queue) > 0:
            last = queue.pop(0)
            self.assertIn(last, [None, A])

    # PriorityQueue
    def test_givenPriorityQueue_when_thenSuccess(self):
        pq = queue.PriorityQueue()
        pq.put((1, "zzf"))
        pq.put((-1, "hhl"))
        self.assertEqual(pq.get(), (-1, "hhl"))
        pass

    # List
    # Random Access
    def test_givenList_when_thenSuccess(self):
        a_list: List[str | None] = []
        # ['A']
        a_list.append(A)
        # ['A', None]
        a_list.append(None)
        # random insert
        # ['A', 'B', None]
        a_list.insert(1, 'B')
        # random read
        self.assertEqual('B', a_list[1])
        # iterate
        for item in a_list:
            pass
        # random remove item
        self.assertEqual('B', a_list.pop(1))

    # dict
    def test_givenDict_when_thenSuccess(self):
        pairs: dict[str, str] = {
            "(": ")",
            "[": ']',
            "{": "}"
        }
        # put
        pairs['A'] = A
        # get
        # throw_error_when_key_not_exist = pairs['KeyNotExist']
        val_none: str | None = pairs.get('KeyNotExist')
        self.assertIsNone(val_none)
        # check dict contains key
        a_in_pairs: bool = A in pairs
        self.assertTrue(a_in_pairs)
        self.assertFalse(')' in pairs)
        # iterate
        for k, v in pairs.copy().items():
            print(f'{k}: {v}')

    # dict
    def test_givenSet_when_thenSuccess(self):
        # 空 set 必须使用 set() 创建
        a_empty_set: set[int] = set()
        a_set: set[str] = {A, 'B'}
        # add to set
        a_set.add(A)
        # check element in set
        a_in_set: bool = A in a_set
        self.assertTrue(a_in_set)
        # remove from set
        a_set.remove(A)
        self.assertFalse(A in a_set)
        # iterate
        for item in a_set:
            pass
        while len(a_set):
            # random pop
            a_set.pop()


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
