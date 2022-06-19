from typing import List
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        front = 0
        
        res = 0
        for i in range(1, len(prices)):
            if prices[i] < prices[front]:
                front = i
            
            res = max(res, prices[i]-prices[front])
            print(prices[i], front, res)
        
        
        return res
        
    
sol = Solution()

test1 = [7,1,5,3,6,4]
test2 = [7,6,4,3,1]


print(sol.maxProfit(test1))
print(sol.maxProfit(test2))
