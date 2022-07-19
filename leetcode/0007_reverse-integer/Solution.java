class Solution {
    public int reverse(int x) {
       int ret = 0;
       while(x != 0) {
         int temp = x % 10;
         if(ret > 214748364 || (ret == 214748364 && temp > 7)) {
            return 0;
         }
         if (ret < -214748364 || (ret == 214748364 && temp < -8)) {
            return 0;
         }
         ret = ret * 10 + temp;
         x = x / 10;
       }
       return ret;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.reverse(39049));
    }
}