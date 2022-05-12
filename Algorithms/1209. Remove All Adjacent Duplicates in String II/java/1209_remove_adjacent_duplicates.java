import java.util.Stack;

class Solution {
    
    public class charStack {
        public char c;
        public int count;
    }
    public String removeDuplicates(String s, int k) {
        
        Stack<charStack> stack = new Stack<>();

        
        for (int i = 0; i < s.length(); i++) {
            // System.out.println(s.charAt(i));
            if (!stack.isEmpty() && stack.peek().c == s.charAt(i)){

                charStack temp = stack.pop();
                temp.count += 1;

                if (temp.count != k) {
                    stack.push(temp);
                }
            } else {
            charStack temp = new charStack();
            temp.c = s.charAt(i);
            temp.count = 1;
            stack.push(temp);
            }
        }

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
