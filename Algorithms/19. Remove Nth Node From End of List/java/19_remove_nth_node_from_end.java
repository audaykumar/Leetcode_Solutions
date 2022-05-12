    // Definition for singly-linked list.
class ListNode {
        int val;
        ListNode next;
        ListNode() {}
        ListNode(int val) { this.val = val; }
        ListNode(int val, ListNode next) { this.val = val; this.next = next; }
}

class Solution {

    public ListNode removeNthFromEnd(ListNode head, int n) {
    ListNode start = new ListNode();
    ListNode slow = start, fast = start;
    start.next = head;
    
    for(int i=1; i<=n+1; i++)   {
        fast = fast.next;
    }

    while (fast != null) {
        slow = slow.next;
        fast= fast.next;
    }
    
    slow.next = slow.next.next;
    return start.next;
    }

    public static void main(String[] args) {
        Solution sol = new Solution();

        ListNode head = new ListNode(1);
        head.next = new ListNode(2);
        head.next.next = new ListNode(3);
        head.next.next.next = new ListNode(4);
        head.next.next.next.next = new ListNode(5);
        
        sol.removeNthFromEnd(head, 2);
    }
}