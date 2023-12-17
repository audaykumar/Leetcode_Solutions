class Solution:
    def longestPalindrome(self, s: str) -> str:
        res_len = 0
        for i in range(len(s)):
            # print(i)
            odd_check = self.check(s,i,i)
            
            even_check = self.check(s,i,i+1)
            if len(odd_check) > len(even_check):
                if len(odd_check) > res_len:
                    res = odd_check
                    res_len = len(odd_check)
            else:
                if len(even_check) > res_len:
                    res = even_check
                    res_len = len(even_check)
                
        return res
    
    def check(Self, s, lower, upper):
            
        while lower >= 0 and upper < len(s) and s[lower] == s[upper]:
            lower-=1
            upper+=1
            # print("current str: ", s[lower+1:upper])
            # print("lower: ", lower)
            # print("upper: ", upper)
            
        # print("return str: ", s[lower+1:upper])
        return s[lower+1:upper]

sol = Solution()

test1 = "babad"
test2 = "cbbd"

print(sol.longestPalindrome(test1))

print(sol.longestPalindrome(test2))