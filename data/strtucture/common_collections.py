from __future__ import annotations

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
        #
        while len(a_set):
            # random pop
            a_set.pop()


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
