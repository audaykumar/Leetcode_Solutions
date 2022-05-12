from cgi import test


class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        
        for char in s:
            if char=='(' or char=='[' or char=='{':
                stack.append(char)
            
            try:
                if char==')':
                    if stack.pop() != '(':
                        return False
                elif char==']':
                    if stack.pop() != '[':
                        return False
                elif char=='}':
                    if stack.pop() != '{':
                        return False
            except:
                return False
        
        if len(stack) != 0:
            return False
        return True
        
    
sol = Solution()

test1 = '()'
test2 = '()[]{}'
test3 = '(]'
test4 = '['

print(sol.isValid(test1))
print(sol.isValid(test2))
print(sol.isValid(test3))
print(sol.isValid(test4))