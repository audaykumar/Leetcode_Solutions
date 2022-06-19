class Solution:
    def climbStairs(self, n: int) -> int:
        if n<=3: return n
        
        n1 = 2
        n2 = 3
        for _ in range(4, n+1):
            temp = n1 + n2
            n1 = n2
            n2 = temp
        
        return n2
            
sol = Solution()

test1 = 2
test2 = 4


print(sol.climbStairs(test1))
print(sol.climbStairs(test2))