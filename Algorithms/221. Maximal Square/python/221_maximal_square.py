from typing import List
class Solution:
    def maximalSquare(self, matrix: List[List[str]]) -> int:
        rows = len(matrix)
        cols = len(matrix[0])
        
        s = []
        max_size = 0
        for i in range(rows):
            temp = []
            for j in range(cols):
                if i==0 or j==0:
                    temp.append(int(matrix[i][j]))
                    max_size = max(max_size, int(matrix[i][j]))
                else:
                    temp.append(0)
            s += [temp]
        
        for i in range(1, rows):
            for j in range(1, cols):
                if matrix[i][j] == "1":
                    s[i][j] = min(s[i][j-1],s[i-1][j],s[i-1][j-1]) + 1
                    max_size = max(max_size, s[i][j])
    
        return (max_size**2)
    
sol = Solution()

test1 = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]

test2 = [["0","1"],["1","0"]]

test3 = [["0"]]

test4 = [["1"]]
print(test1, sol.maximalSquare(test1))
print(test2, sol.maximalSquare(test2))
print(test3, sol.maximalSquare(test3))
print(test4, sol.maximalSquare(test4))