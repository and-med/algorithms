package com.devandrew;

public class _97InterleavingStrings {
    private String s1;
    private String s2;
    private String s3;
    private Boolean memo[][];

    private boolean isInterleaveInternal(int left, int right) {
        if (left == s1.length() && right == s2.length()) {
            return true;
        }

        if (memo[left][right] != null) {
            return memo[left][right];
        }

        boolean answer = false;

        if (left < s1.length()) {
            answer = answer || (s1.charAt(left) == s3.charAt(left + right)
                    && isInterleaveInternal(left + 1, right));
        }

        if (right < s2.length()) {
            answer = answer || (s2.charAt(right) == s3.charAt(left + right)
                    && isInterleaveInternal(left, right + 1));
        }

        memo[left][right] = answer;

        return answer;
    }

    public boolean isInterleave(String s1, String s2, String s3) {
        if (s3.length() != s1.length() + s2.length()) {
            return false;
        }

        this.s1 = s1;
        this.s2 = s2;
        this.s3 = s3;

        this.memo = new Boolean[s1.length() + 1][s2.length() + 1];

        return isInterleaveInternal(0, 0);
    }
}
