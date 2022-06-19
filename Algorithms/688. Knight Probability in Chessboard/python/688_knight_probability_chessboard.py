# WORK IN PROGRESS
class Solution:
    def knightProbability(self, n: int, k: int, row: int, column: int) -> float:
        prob = 1.0
        
        prob_list = [[[-1]]*n]*n
        
        prob_list = self.calc_prob(n, prob_list)
        print(prob_list)
        for i in range(k):
            pass

        return prob
    
    def calc_prob(self, n, prob_list):
        for r in range (n):
             count = 0
             for c in range(n):
                if r-2 >= 0 and c-1 >=0:
                    count +=1
                if r-1 >= 0 and c-2 >=0:
                    count +=1
                if r-2 >= 0 and c+1 <n:
                    count +=1
                if r-1 >= 0 and c+2 <n:
                    count +=1
                if r+2 < n and c-1 >= 0:
                    count +=1
                if r+1 < n and c-2 >= 0:
                    count +=1
                if r+2 < n and c+1 < n:
                    count +=1
                if r+1 < n and c+2 < n:
                    count +=1
                
                # print(r,c)
                prob_list[r][c] = count/8
                    
        return prob_list     
    
sol = Solution()

test1 = [3,2,0,0]
test2 = [1,0,0,0]

print(sol.knightProbability(test1[0],test1[1],test1[2],test1[3]))
# print(sol.knightProbability(test2[0],test2[1],test2[2],test2[3]))