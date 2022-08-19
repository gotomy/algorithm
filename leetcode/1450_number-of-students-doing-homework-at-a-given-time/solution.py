from typing import List

class Solution:
    def busyStudent(self, startTime: List[int], endTime: List[int], queryTime: int) -> int:
        busy_student = 0
        for i in range(0, len(startTime)):
            if startTime[i] <= queryTime and endTime[i] >= queryTime:
                busy_student += 1
        return busy_student
        

solution = Solution()
print(solution.busyStudent([1,2,3], [3,2,7], 4))
print(solution.busyStudent([9,8,7,6,5,4,3,2,1], [10,10,10,10,10,10,10,10,10], 5))
print(solution.busyStudent([1,1,1,1], [1,3,2,4], 7))