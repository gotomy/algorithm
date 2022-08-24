from typing import List

class Solution:
    def canBeEqual(self, target: List[int], arr: List[int]) -> bool:
        target.sort()
        arr.sort()
        for i in range(0, len(target)):
            if target[i] != arr[i]:
                return False
        return True

solution = Solution()
print(solution.canBeEqual([1,2,3,4],[2,2,1,3]))