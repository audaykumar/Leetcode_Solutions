from typing import List
class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        result = []
        sorted_nums = sorted(nums)
        i=0
        while(i < len(sorted_nums)):
            target = -1 * sorted_nums[i]

            front = i+1
            back = len(sorted_nums)-1
            
            if target < 0:
                break
            
            while front < back:
                sum = sorted_nums[front] + sorted_nums[back]
                
                if sum < target:
                    front += 1
                elif sum > target:
                    back -= 1
                else:
                    triplet = [sorted_nums[i], sorted_nums[front], sorted_nums[back]]
                    result.append(triplet)
                    
                    while ((front < back) and (sorted_nums[front] == triplet[1])):
                        front += 1
                    
                    while ((back > front) and (sorted_nums[back] == triplet[2])):
                        back -= 1
                    
            while((i+1 < len(sorted_nums)) and (sorted_nums[i+1]==sorted_nums[i])):
                i += 1
            i+=1
        return result
    
sol = Solution()


test1 = [-1,0,1,2,-1,-4]
test2 = []
test3 = [0]
test4 = [34, 55, 79, 28, 46, 33, 2, 48, 31, -3, 84, 71, 52, -3, 93, 15, 21, -43, 57, -6, 86, 56, 94, 74, 83, -14, 28, -66, 46, -49, 62, -11, 43, 65, 77, 12, 47, 61, 26, 1, 13, 29, 55, -82, 76, 26, 15, -29, 36, -29, 10, -70, 69, 17, 49]


print(sol.threeSum(test1))
print(sol.threeSum(test2))
print(sol.threeSum(test3))
print(sol.threeSum(test4))