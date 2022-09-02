from typing import List

class Solution:
    def maxArea(self, height: List[int]) -> int:
        max_area = 0
        left, right = 0, len(height)-1
        while left < right:
            min_height = 0
            width = right - left
            if height[left] <= height[right]:
                min_height = height[left]
                left += 1
            else:
                min_height = height[right]
                right -= 1
            temp_area = min_height * (width)
            if temp_area > max_area:
                max_area = temp_area
           
        return max_area

solution = Solution()
print(solution.maxArea([1,8,6,2,5,4,8,3,7]))
print(solution.maxArea([1,1]))