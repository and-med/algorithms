package com.devandrew;

public class _10RegularExpressionMatching {
    private int[][] storage;

    public boolean isMatch(String s, String p) {
        storage = new int[s.length() + 1][p.length() + 1];
        return isMatch(s, p, 0, 0);
    }

    private boolean isMatch(String s, String p, int i, int j) {
        if (storage[i][j] != 0) {
            return storage[i][j] == 1;
        }

        boolean answer;

        if (j == p.length()) {
            answer = i == s.length();
        } else {
            boolean areEqual = i != s.length() && (s.charAt(i) == p.charAt(j) || p.charAt(j) == '.');

            if (j + 1 < p.length() && p.charAt(j + 1) == '*') {
                if (areEqual) {
                    answer = isMatch(s, p, i, j + 2) || isMatch(s, p, i + 1, j);
                } else {
                    answer = isMatch(s, p, i, j + 2);
                }
            } else {
                answer = areEqual && isMatch(s, p, i + 1, j + 1);
            }
        }

        storage[i][j] = answer ? 1 : 2;
        return answer;
    }
}
