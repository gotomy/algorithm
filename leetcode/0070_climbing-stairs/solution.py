class Solution:
    def climbStairs(self, n: int) -> int:
        # dfs暴力，时间超时
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

    def climbStairs2(self, n: int) -> int:
        # 动态规则，转移数组
        p, q, r = 0, 0 ,1
        for i in range(1, n+1):
            p = q
            q = r
            r = p + q
        return r

solution = Solution()
print(solution.climbStairs2(35))