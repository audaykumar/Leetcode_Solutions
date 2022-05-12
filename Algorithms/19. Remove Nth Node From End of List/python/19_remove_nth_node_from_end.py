from typing import Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        
        len = 1
        temp = head
        while temp!= None:
            len +=1
            temp=temp.next
        
        if len == 1:
            return None
        
        size = len-n
        
        if size == 1:
            return head.next
        
        temp = head
        while size != 1:
            previous = temp
            temp = temp.next
            size -= 1
        
        
        previous.next = temp.next
        return head
    

        
    
sol = Solution()

test1 = ListNode(1, ListNode(2, ListNode(3, ListNode(4, ListNode(5, None)))))
test2 = ListNode(1, ListNode(2, None))
test3 = ListNode(1, None)

sol.removeNthFromEnd(test1,2)
sol.removeNthFromEnd(test2,2)
sol.removeNthFromEnd(test3,1)