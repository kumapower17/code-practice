#!/usr/bin/env python3

# https://leetcode-cn.com/problems/compare-version-numbers/

from itertools import zip_longest


class Solution:
    @staticmethod
    def cmp(a, b):
        n = max(len(a), len(b))
        za = a.zfill(n)
        zb = b.zfill(n)
        if za < zb:
            return -1
        if za == zb:
            return 0
        return 1

    def compareVersion(self, version1, version2):
        v1 = version1.split('.')
        v2 = version2.split('.')
        pairs = zip_longest(v1, v2, fillvalue='0')
        for a, b in pairs:
            r = self.cmp(a, b)
            if r != 0:
                return r
        return 0
