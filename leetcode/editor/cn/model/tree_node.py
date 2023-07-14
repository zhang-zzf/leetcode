#!/usr/bin/env python3
# -*- coding: utf-8 -*-

# 类定义过程中使用注解引用类本身
from __future__ import annotations

from typing import List, Optional


class TreeNode:

    def __init__(self, val: int = 0,
                 left: TreeNode = None,
                 right: TreeNode = None):
        self.val = val
        self.left = left
        self.right = right

    # 类方法
    @classmethod
    def decode(cls, nodes: List[int]) -> Optional[TreeNode]:
        if len(nodes) == 0:
            return None
        root = cls(nodes.pop(0))
        level = [root]
        while len(level) > 0:
            r = level.pop(0)
            if not r:
                continue
            left_filled: bool = False
            for i in range(min(2, len(nodes))):
                n = nodes.pop(0)
                # 三目运算符
                node = cls(n) if n else None
                if not left_filled:
                    r.left = node
                    left_filled = True
                else:
                    r.right = node
                level.append(node)
        return root
