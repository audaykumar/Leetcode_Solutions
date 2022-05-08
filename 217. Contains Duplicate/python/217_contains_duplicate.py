from typing import List

class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        check = False   
        num_map = {}
        for n in nums:
            if n not in num_map:
                num_map[n] = True
            else:
                check = True
        return check
        
        # Other possible solution
        # return len(nums) != len(set(nums))
    
sol = Solution()

test1 = [1,2,3,1]
test2 = [1,2,3,4]
test3 = [1,1,1,3,3,4,3,2,4,2]


print(sol.containsDuplicate(test1))
print(sol.containsDuplicate(test2))
print(sol.containsDuplicate(test3))