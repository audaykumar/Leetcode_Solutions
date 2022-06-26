from typing import List
class Solution:
    def numPairsDivisibleBy60(self, time: List[int]) -> int:
        songMap = {}
        count = 0
        for t in time:
            check = t%60

            target = 60 - check
            if target == 60:
                target = 0
            # print(songMap, check, target)
            if target in songMap:
                count += songMap[target]
        
            if check in songMap:
                songMap[check] +=1
            else:
                songMap[check] = 1
        return count
    
sol = Solution()

test1 = [30,20,30,150,100,40]

test2 = [60,60,60]

print(sol.numPairsDivisibleBy60(test1))

print(sol.numPairsDivisibleBy60(test2))