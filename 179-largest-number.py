#!/usr/bin/env python3
from functools import cmp_to_key


class Solution:
    @staticmethod
    def cmp(a, b):
        if a + b > b + a:
            return -1
        return 1

    def largestNumber(self, nums):
        largest = "".join(sorted(map(str, nums), key=cmp_to_key(self.cmp)))
        if largest < '1':
            largest = '0'
        return largest
