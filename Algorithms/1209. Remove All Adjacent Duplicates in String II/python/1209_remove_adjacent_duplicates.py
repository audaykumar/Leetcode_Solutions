class Solution:
    def removeDuplicates(self, s: str, k: int) -> str:
        stack = []
        for char in s:
            if stack and stack[-1][0] == char:
                stack[-1][1] += 1
            else:
                stack.append([char,1])
            
            if stack[-1][1] == k:
                stack.pop()
            output = ""
            
        for char, count in stack:
            output += char * count
        return output
    
sol = Solution()

print(sol.removeDuplicates("abcd", 2))
print(sol.removeDuplicates("deeedbbcccbdaa", 3))
print(sol.removeDuplicates("pbbcggttciiippooaais", 2))