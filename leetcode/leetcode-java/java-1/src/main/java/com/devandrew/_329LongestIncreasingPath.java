package com.devandrew;

public class _329LongestIncreasingPath {
    private int[][] matrix;

    private Integer[][] memo;

    private int longestIncreasingPathAt(int i, int j) {
        if (memo[i][j] != null) {
            return memo[i][j];
        }

        int longestPath = 0;

        if (i + 1 < matrix.length && matrix[i + 1][j] > matrix[i][j]) {
            longestPath = Math.max(longestPath, 1 + longestIncreasingPathAt(i + 1, j));
        }

        if (i - 1 >= 0 && matrix[i - 1][j] > matrix[i][j]) {
            longestPath = Math.max(longestPath, 1 + longestIncreasingPathAt(i - 1, j));
        }

        if (j + 1 < matrix[i].length && matrix[i][j + 1] > matrix[i][j]) {
            longestPath = Math.max(longestPath, 1 + longestIncreasingPathAt(i, j  + 1));
        }

        if (j - 1 >= 0 && matrix[i][j - 1] > matrix[i][j]) {
            longestPath = Math.max(longestPath, 1 + longestIncreasingPathAt(i, j - 1));
        }

        memo[i][j] = longestPath;

        return longestPath;
    }

    public int longestIncreasingPath(int[][] matrix) {
        int longest = 0;

        this.matrix = matrix;
        this.memo = new Integer[matrix.length][matrix[0].length];

        for (int i = 0; i < matrix.length; i++) {
            for (int j = 0; j < matrix[i].length; j++) {
                longest = Math.max(longest, longestIncreasingPathAt(i, j) + 1);
            }
        }

        return longest;
    }
}
