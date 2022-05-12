from typing import Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
        
class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        """
        Do not return anything, modify head in-place instead.
        """
        if not head or not head.next:
            return
        
        slow, fast = head, head
        
        while fast and fast.next:
            slow = slow.next
            fast= fast.next.next
        
        
        # Reverse second half of the list
        
        # Set current to the node after where slow ended 
        current = slow.next
        
        # Set the next node of slow to null
        slow.next = None
    
    
        # Reverse the list
        previous = None
        next = None
        while (current != None):
            next =  current.next
            current.next = previous
            previous = current
            current = next
            
        temp = head
        while previous:
            temp.next, temp = previous, temp.next
            previous.next, previous = temp, previous.next
            
        
        
    
sol =  Solution()

test1 = ListNode(1, ListNode(2, ListNode(3, ListNode(4, ListNode(4, None)))))
test2 = ListNode(1, ListNode(2, None))
test3 = None

sol.reorderList(test1)
# sol.reverseList(test2)
# sol.reverseList(test3)