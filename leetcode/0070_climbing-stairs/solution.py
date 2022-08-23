class Solution:
    def climbStairs(self, n: int) -> int:
        def dfs(target, step, total):
            left = target - step
            if left == 0:
                total[0] += 1
                return
            if left < 0:
                return
            dfs(left, 1, total)
            dfs(left, 2, total)
        total = [0]
        dfs(n, 1, total)
        dfs(n, 2, total)
        return total[0]


solution = Solution()
print(solution.climbStairs(35))