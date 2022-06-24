from typing import Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        check1 = check2 = head 
        while check2 and check2.next:
            check1 = check1.next
            check2 = check2.next.next
            if check1 == check2:
                return True
        return False

sol = Solution()

test1 = ListNode(3)
test1.next = ListNode(2)
test1.next.next = ListNode(0)
test1.next.next.next = ListNode(-4)
test1.next.next.next.next = test1.next

test2 = ListNode(1)
test2.next = ListNode(2)
test2.next.next = test2
# test2.next.next = ListNode(3)

test3 = ListNode(1)

test4 = None


print(sol.hasCycle(test1))
print(sol.hasCycle(test2))
print(sol.hasCycle(test3))
print(sol.hasCycle(test4))