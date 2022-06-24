from typing import List
class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        sorted_list = sorted(intervals)
        res = []
        first = sorted_list[0][0]
        second = sorted_list[0][1]
        for i in range(1, len(sorted_list)):
            if sorted_list[i][0] <= second:
                if sorted_list[i][1] > second:
                    second = sorted_list[i][1]
            else:
                res.append([first, second])
                first = sorted_list[i][0]
                second = sorted_list[i][1]
        
        res.append([first, second])
        return res
    
sol = Solution()

test1 = [[2,6],[1,3],[8,10],[15,18]]
test2 = [[1,4],[4,5]]
test3 = [[1,4],[2,3]]

print(sol.merge(test3))