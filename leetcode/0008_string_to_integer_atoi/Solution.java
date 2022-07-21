import java.util.HashMap;
import java.util.Map;

class Solution{
    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.myAtoi("20000000000000000000"));
        System.out.println(solution.myAtoi("+-03043"));
        System.out.println(solution.myAtoi("3434"));
        System.out.println(solution.myAtoi("abd3od39094"));
    }

    public int myAtoi(String str) {
        Automaton automaton = new Automaton();
        int length = str.length();
        for(int i=0; i< length; i++) {
            automaton.get(str.charAt(i));
        }
        return (int)(automaton.sign * automaton.answer);
    }
}

class Automaton {
    public int sign = 1;
    public long answer = 0;
    private String state = "start";
    private Map<String, String[]> table = new HashMap<String, String[]>(){
        {
            put("start", new String[]{"start", "signed", "in_number", "end"});
            put("signed", new String[]{"end", "end", "in_number", "end"});
            put("in_number", new String[]{"end", "end", "in_number", "end"});
            put("end", new String[]{"end", "end", "end", "end"});
        }
    };

    public void get(char c) {
        state = table.get(state)[getCol(c)];
        if ("in_number".equals(state)) {
            answer = answer * 10 + c - '0';
            answer = sign == 1 ? Math.min(answer, (long)Integer.MAX_VALUE) : Math.min(answer, -(long)Integer.MIN_VALUE);
        } else if ("signed".equals(state)){
            sign = c == '+' ? 1 : -1;
        }
    }

    private int getCol(char c) {
        if (c == ' ') {
            return 0;
        }
        if (c == '+' || c== '-') {
            return 1;
        }
        if (Character.isDigit(c)) {
            return 2;
        }
        return 3;
    }
}