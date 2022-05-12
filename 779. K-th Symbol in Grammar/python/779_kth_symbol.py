class Solution:
    def kthGrammar(self, n: int, k: int) -> int:
        if n==1:
            return 0
        size =  2**(n-1)
        if k <= size/2:
            return self.kthGrammar(n-1, k)
        else:
            return int(not self.kthGrammar(n-1, int(k-(size/2))))

sol = Solution()

test1 = [1,1]
test2 = [2,1]
# test3 = [2,2]
test3 = [3,4]


print(sol.kthGrammar(test1[0],test1[1]))
print(sol.kthGrammar(test2[0],test2[1]))
print(sol.kthGrammar(test3[0],test3[1]))
