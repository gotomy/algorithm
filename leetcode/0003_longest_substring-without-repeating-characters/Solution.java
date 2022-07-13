class Solution {

    public int lengthOfLongestSubstring(String s) {        
        int longestSubLength = 0;
        String longestSubstring = "";

        int tempLength = 0;
        StringBuilder tempSubstring = new StringBuilder();

        for(int index=0; index < s.length(); index ++) {
           String strItem = String.valueOf(s.charAt(index));
           int happenIndex = tempSubstring.indexOf(strItem);
           if(happenIndex!= -1) {
                if (happenIndex + 1 == tempSubstring.length()) {
                    tempSubstring = new StringBuilder(strItem);
                } else {
                    tempSubstring = new StringBuilder(tempSubstring.substring(happenIndex + 1)).append(strItem);
                }
                tempLength = tempSubstring.length();
           } else {
                tempSubstring.append(strItem);
                tempLength++;
           }

           if(tempLength > longestSubLength) {
                longestSubLength = tempLength;
                longestSubstring  = tempSubstring.toString();
           }
        }
        System.out.println(longestSubstring);
        return longestSubLength;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.lengthOfLongestSubstring("pwwkew"));
    }
    
}
