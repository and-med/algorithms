package com.devandrew._79;

public class _79 {
    private static final int[][] directions = new int[][] { { 1, 0 }, { 0, 1 }, { -1, 0 }, { 0, -1 } };
    private char[][] board;
    private int m;
    private int n;

    public boolean backtrack(String word, int i, int j, int idx) {
        if (idx == word.length() || board[i][j] != word.charAt(idx)) {
            return false;
        } else if (idx == word.length() - 1) {
            return true;
        }

        boolean res = false;
        board[i][j] = '#';
        for (int[] d: directions) {
            int x = i + d[0];
            int y = j + d[1];
            if (x >= 0 && x < m && y >= 0 && y < n) {
                res = res || backtrack(word, x, y, idx + 1);
            }
        }
        board[i][j] = word.charAt(idx);

        return res;
    }

    public boolean exist(char[][] board, String word) {
        this.board = board;
        this.m = board.length;
        this.n = board[0].length;

        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[i].length; j++) {
                if (backtrack(word, i, j, 0)) {
                    return true;
                }
            }
        }

        return false;
    }
}
