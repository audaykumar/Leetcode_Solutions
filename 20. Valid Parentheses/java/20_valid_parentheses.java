import java.util.ArrayList;
import java.util.Stack;

class Solution {
    public boolean isValid(String s) {
        // Stack< x = new ArrayList<>()
        Stack<Character> stack = new Stack<>();
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == '(') {
                stack.push(')');
            }else if (s.charAt(i) == '['){
                stack.push(']');
            }else if (s.charAt(i) == '{'){
                stack.push('}');
            } else if (stack.isEmpty() || stack.pop() != s.charAt(i)){
                return false;
            }
        }
        // System.out.println(stack);
        return stack.isEmpty();
    }

public static void main(String[] args) {
    Solution sol = new Solution();

    String test1 = new String("()");
    String test2 = new String("()[]{}");
    String test3 = new String("(]");
    String test4 = new String("[");
    String test5 = new String("]");

    System.out.println(sol.isValid(test1));
    System.out.println(sol.isValid(test2));
    System.out.println(sol.isValid(test3));
    System.out.println(sol.isValid(test4));
    System.out.println(sol.isValid(test5));
    }
}