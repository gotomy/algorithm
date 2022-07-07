from typing import List

class Solution:
    def twoSum(self, nums : List[int], target: int) -> List[int]:
        dict = {}
        for idx in range(len(nums)):
            other = target - nums[idx]
            if other in dict:
                return [dict[other], idx]
            dict[nums[idx]] = idx
        return None
    
s = Solution()
nums = [2, 7, 11, 15]
target = 18
print(s.twoSum(nums, target))