from ast import While
from typing import Optional
# Definition for singly-linked list.
class ListNode:   
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
        
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        current = head
        previous = None
        next = None
        while (current != None):
            next =  current.next
            current.next = previous
            previous = current
            current = next
            
        return previous
    
sol =  Solution()

test1 = ListNode(1, ListNode(2, ListNode(3, ListNode(4, ListNode(4, None)))))
test2 = ListNode(1, ListNode(2, None))
test3 = None

print(sol.reverseList(test1))
print(sol.reverseList(test2))
print(sol.reverseList(test3))
