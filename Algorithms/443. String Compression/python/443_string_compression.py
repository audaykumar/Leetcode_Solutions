from typing import List

class Solution:
    def compress(self, chars: List[str]) -> int: 
        i = index = 0
        
        n = len(chars)
        while i < n:
            j=i
            while j<n and chars[j] == chars[i]:
                j += 1
            
            freq = j-i
            
            chars[index] = chars[i]
            index +=1
            
            if freq > 1:
                freq = str(freq)
                for val in freq:
                    chars[index] = val
                    index +=1
            i = j
            
        print(chars[:index])
        return index
    
sol = Solution()

test1 = ["a","a","b","b","c","c","c"]

test2 = ["a"]

test3 = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]

print(sol.compress(test1))

print(sol.compress(test2))

print(sol.compress(test3))

