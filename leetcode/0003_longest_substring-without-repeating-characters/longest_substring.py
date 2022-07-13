class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        mp = {}
        left, maxVal = 0, 0
        i = 0
        for val in s :
          if val in mp:
                left =  max(left, mp[val] + 1)
          mp[val] = i
          maxVal = max(maxVal, i - left +1)
          i = i + 1

        return maxVal;

solution = Solution()
print(solution.lengthOfLongestSubstring("abcabcbb"))