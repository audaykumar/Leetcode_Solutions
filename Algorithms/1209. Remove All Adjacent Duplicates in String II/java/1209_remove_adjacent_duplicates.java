import java.util.Stack;

class Solution {
    // Creating a charStack to hold character and its count value
    public class charStack {
        public char c;
        public int count;
    }
    public String removeDuplicates(String s, int k) {
        
        Stack<charStack> stack = new Stack<>();

        // Loop over the input string
        for (int i = 0; i < s.length(); i++) {

            // If the stack is not empty and the input character is the same as the top of the stack
            // Update count
            if (!stack.isEmpty() && stack.peek().c == s.charAt(i)){

                charStack temp = stack.pop();
                temp.count += 1;

                // If the count is not equal to k value push it back to the stack. Else dont
                if (temp.count != k) {
                    stack.push(temp);
                }
            } else {
                // Encountering first element or new element so push it to stack with count as 1.
                charStack temp = new charStack();
                temp.c = s.charAt(i);
                temp.count = 1;
                stack.push(temp);
            }
        }

        // Build the final output string based on the stack
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < stack.size(); i++) {
            String temp =String.valueOf(stack.elementAt(i).c).repeat(stack.elementAt(i).count);  
            sb.append(temp);
        }
     
        return sb.toString();
        
    }

    public static void main(String[] args) {
        Solution sol = new Solution();
        
        System.out.println(sol.removeDuplicates("abcd", 2));
        System.out.println(sol.removeDuplicates("deeedbbcccbdaa", 3));
        System.out.println(sol.removeDuplicates("pbbcggttciiippooaais", 2));
    }
}
