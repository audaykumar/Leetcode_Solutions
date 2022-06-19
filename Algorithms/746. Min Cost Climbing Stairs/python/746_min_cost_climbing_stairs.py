from cgi import test
from typing import List
class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        n1 = cost[0]
        n2 = cost[1]
        for i in range(2, len(cost)):
            temp = min(n1, n2)
            n1 = n2
            n2 = temp + cost[i]
        final = min(n1,n2)
        return final
    
sol = Solution()

test1 = [10,15,20]
test2 = [1,100,1,1,1,100,1,1,100,1]

print(sol.minCostClimbingStairs(test1))
print(sol.minCostClimbingStairs(test2))