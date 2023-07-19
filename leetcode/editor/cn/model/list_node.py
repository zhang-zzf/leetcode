#!/usr/bin/env python3
# -*- coding: utf-8 -*-

# 类定义过程中使用注解引用类本身
from __future__ import annotations


class ListNode:
    def __init__(self, val=0, nextNode: ListNode = None):
        self.val = val
        self.next = nextNode
