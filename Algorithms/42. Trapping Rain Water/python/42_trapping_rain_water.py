from typing import List
class Solution:
    def trap(self, height: List[int]) -> int:
        left = -1
        right = -1
        max_height_left = []
        # Find largest heigh from left hand side
        for h in height:
            left = max(left, h)
            max_height_left.append(left)
        
        max_height_left.reverse()
        
        # Find largest height from the RHS.
        # At each level find the smaller height between LHS and RHS
        # Subtract the current position level from this height to get trapped water
        height.reverse()
        water = 0
        for i, val in enumerate(height):
            right = max(right, val)
            water += min(max_height_left[i], right) - val

        return water
    
sol = Solution()

test1 = [0,1,0,2,1,0,1,3,2,1,2,1]
test2 = [4,2,0,3,2,5]

print(sol.trap(test1))
print(sol.trap(test2))