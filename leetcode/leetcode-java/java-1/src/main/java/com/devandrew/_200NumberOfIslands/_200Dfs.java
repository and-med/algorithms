package com.devandrew._200NumberOfIslands;

public class _200Dfs {
    private char[][] grid;
    private boolean[][] visited;
    private int[][] directions = new int[][] { { 0, 1 }, { 0, - 1}, { 1, 0 }, { -1, 0 } };
    private int m;
    private int n;

    private boolean dfs(int i, int j) {
        if (grid[i][j] == '0' || visited[i][j]) {
            return false;
        }

        visited[i][j] = true;

        for (int[] dir: directions) {
            int new_i = i + dir[0];
            int new_j = j + dir[1];

            if (new_i >= 0 && new_i < this.m && new_j >= 0 && new_j < this.n) {
                dfs(new_i, new_j);
            }
        }

        return true;
    }

    public int numIslands(char[][] grid) {
        this.grid = grid;
        this.m = grid.length;
        this.n = grid[0].length;
        this.visited = new boolean[m][n];

        int cnt = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (dfs(i, j)) {
                    cnt++;
                }
            }
        }

        return cnt;
    }
}
