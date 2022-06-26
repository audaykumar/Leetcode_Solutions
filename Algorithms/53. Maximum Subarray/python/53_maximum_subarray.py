from typing import List
class Solution:
    def maxSubArray(self, nums: List[int]) -> int:    
        curr_max = nums[0]
        global_max = nums[0]
        
        for i in range(1, len(nums)):
            curr_max = max(curr_max + nums[i], nums[i])
            
            global_max = max(global_max, curr_max)
        
        return global_max
    
sol = Solution()

test1 = [-2,1,-3,4,-1,2,1,-5,4]

print(sol.maxSubArray(test1))