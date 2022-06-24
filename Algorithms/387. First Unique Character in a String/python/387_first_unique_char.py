class Solution:
    def firstUniqChar(self, s: str) -> int:
        map = {}
        for i, c in enumerate(s):
            if c not in map:
                map[c] = [1, i]
            else:
                # map.pop(c)
                map[c][0] += 1
        
        val = -1
        for k, v in map.items():
            if v[0] == 1:
                if val != -1:
                    val = min(val, v[1])
                else:
                    val = v[1]
                
                
        return val
        
    
sol = Solution()

test1 = "leetcode"
test2 = "loveleetcode"
# test3 = [2,2]
test3 = "aabb"

test4 = "aadadaad"

# print(sol.firstUniqChar(test1))
# print(sol.firstUniqChar(test2))
# print(sol.firstUniqChar(test3))
print(sol.firstUniqChar(test4))