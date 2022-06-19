class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        start = maxlen = 0
        map = {}
        for i, c in enumerate(s):
            if c in map and start <= map[c]:
                start = map[c]+1
            else:
                maxlen = max(maxlen, i-start+1)
            map[c] = i
        return maxlen
    

sol = Solution()

test1 = "abcabcbb"
test2 = "bbbbb"
test3 = "pwwkew"
test4 = "tmmzuxt"


print(sol.lengthOfLongestSubstring(test1))
print(sol.lengthOfLongestSubstring(test2))
print(sol.lengthOfLongestSubstring(test3))
print(sol.lengthOfLongestSubstring(test4))