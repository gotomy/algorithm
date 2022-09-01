from typing import List

class Solution:
    def finalPrices(self, prices: List[int]) -> List[int]:
        ret = []
        for i in range(0, len(prices)):
            final_price = prices[i]
            for j in range(i+1, len(prices)):
                if final_price >= prices[j]:
                    final_price = final_price - prices[j]
                    break
            ret.append(final_price)
        return ret


solution = Solution()
print(solution.finalPrices([8,4,6,2,3]))
print(solution.finalPrices([1,2,3,4,5]))
print(solution.finalPrices([10,1,1,6]))