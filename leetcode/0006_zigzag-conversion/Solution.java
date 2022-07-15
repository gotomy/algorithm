class Solution {

    public String convert(String s, int numRows) {
        if(numRows == 1) {
            return s;
        }

        if(numRows >= s.length()) {
            return s;
        }

        // 矩阵填写时，向下填写r个字符，然后向右上填写r-2个字符后回到第一行，周期为t=r+r-2 = 2r-2，
        // 每周期使用1+r-2 = r-1列
        int t = numRows * 2 -2;
        int c = (s.length() + t -1)/t * (numRows -1);
        char[][] mat = new char[t][c];
        for(int i=0, x =0, y=0; i < s.length(); ++i) {
            mat[x][y] = s.charAt(i);
            if (i % t < numRows - 1) {
                // 下移
                ++x;
            } else {
                --x;
                ++y; //右上移
            }
        }

        StringBuilder sb = new StringBuilder();
        for(char[] row : mat) {
            for(char ch : row) {
                if(ch != 0) {
                    sb.append(ch);
                }
            }
        }
        return sb.toString();
    }
    
    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.convert("PAYPALISHIRING", 4));
    }
}
