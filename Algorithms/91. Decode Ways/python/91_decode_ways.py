class Solution:
    def numDecodings(self, s: str) -> int:
        dp = [0 for _ in range(len(s)+1)]
        
        dp[len(s)] = 1
        for i in range(len(s)-1, -1,-1):
            if s[i] == "0":
                dp[i] = 0
            else:
                dp[i] = dp[i+1]
                
            if i+1 < len(s) and (s[i] == "1" or s[i] == "2" and s[i+1] in "0123456"):
                dp[i] += dp[i+2]
            
        return dp[0]
    
sol = Solution()

test1 = "12"
test2  = "226"
test3 = "06"

test4 = "11106"
test5 = "6"

print(sol.numDecodings(test1))
print(sol.numDecodings(test2))
print(sol.numDecodings(test5))
print(sol.numDecodings(test4))