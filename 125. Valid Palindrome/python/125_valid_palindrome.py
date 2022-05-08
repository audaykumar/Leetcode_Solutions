class Solution:
    def isPalindrome(self, s: str) -> bool:
        stripped_str = ''.join(ch.lower() for ch in s if ch.isalnum())
        length = len(stripped_str)
        if length==0:
            return True
        
        if length%2!=0:
            front = back = int(length/2)
        else:
            front = int(length/2)-1
            back = int(front + 1)

        while (front >=0 and back < length and stripped_str[front]==stripped_str[back]):
            front -=1
            back +=1
            
        if front == -1 and back == length:
            return True
        
        return False
        
sol = Solution()

test1 = "A man, a plan, a canal: Panama"
test2 = "race a car"
test3 = " "


print(sol.isPalindrome(test1))
print(sol.isPalindrome(test2))
print(sol.isPalindrome(test3))