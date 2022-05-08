from typing import List
class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        result = []
        sorted_candidates = sorted(candidates)
        self.find_list(sorted_candidates, target, [], result)
        return result
    
    def find_list(self, candidates, target, cur_path, result):
        if target == 0:
            result.append(cur_path)
            return
        for i in range(len(candidates)):
            if candidates[i] > target:
                break
            self.find_list(candidates[i:], target-candidates[i], cur_path+[candidates[i]], result)
        
sol = Solution()

test1 = [2,3,6,7]
test2 = [2,3,5]
test3 = [2]


print(sol.combinationSum(test1, 7))
print(sol.combinationSum(test2, 8))
print(sol.combinationSum(test3, 1))