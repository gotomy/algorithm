from typing import List

class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        lastNum = digits[len(digits)-1]
        newData = []
        if lastNum != 9:
            # 无进位
            for i in range(0, len(digits)):
                if i == len(digits)-1:
                    newData.append(digits[i]+1)
                else:
                    newData.append(digits[i])
        else:
            # 有进位
            carry = 1
            newData.append(0)
            for i in range(len(digits)-2, -1, -1):
               newData.append((digits[i]+carry)%10)
               carry = (digits[i]+carry) // 10
            if carry == 1:
                newData.append(carry)
            newData.reverse()
        return newData


solution = Solution()
print(solution.plusOne([1,9,9,8,9]))