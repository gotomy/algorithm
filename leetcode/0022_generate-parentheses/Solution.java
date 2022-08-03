import java.util.*;

class Solution {
    public List<String> generateParenthesis(int n) {
        List<String> res = new ArrayList<>();
        generateOne(res, "", n, n, "L");
        return res;
    }
    
    private void generateOne(List<String> list, String string, int left, int right, String flag) {
        // left, rigth 分别代表可用的左括号数和可用的右括号数，初始都是 n个可用
        System.out.println(string + ", left:" + left +", right:" + right + ", flag:" + flag);
        if (left == 0 && right == 0) {
            list.add(string);
            System.out.println();
            return;
        }
        
        if (left > 0) {
            generateOne(list, string + "(", left - 1, right, "L");
        }
        
        // 可用的括号 右括号大于左括号时，说明有 左括号先放置，才会是有效的括号组合
        if (right > left) {
            generateOne(list, string + ")", left, right - 1, "R");    
        }        
        
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.generateParenthesis(3));
    }
}