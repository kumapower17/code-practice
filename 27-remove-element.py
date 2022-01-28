class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        i = 0
        for v in nums:
            if v != val:
                nums[i] = v
                i += 1
        return i
