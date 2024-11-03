package com.devandrew._200NumberOfIslands;


public class _200UnionFind {
    public class UnionFind {
        private int[] parent;
        private int[] rank;

        public UnionFind(int n) {
            this.parent = new int[n];
            this.rank = new int[n];

            for (int i = 0; i < n; i++) {
                parent[i] = i;
                rank[i] = 0;
            }
        }

        public int find(int i) {
            if (i != parent[i]) {
                parent[i] = find(parent[i]);
            }

            return parent[i];
        }

        public boolean union(int l, int r) {
            int parentLeft = find(l);
            int parentRight = find(r);

            if (parentLeft == parentRight) {
                return false;
            }

            if (rank[parentLeft] < rank[parentRight]) {
                parent[parentLeft] = parentRight;
            } else if (rank[parentRight] < rank[parentLeft]) {
                parent[parentRight] = parentLeft;
            } else {
                parent[parentLeft] = parentRight;
                rank[parentRight]++;
            }

            return true;
        }
    }


    private int[][] directions = new int[][] { { 0, 1 }, { 0, - 1}, { 1, 0 }, { -1, 0 } };

    public int numIslands(char[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        UnionFind union = new UnionFind(m * n);

        int cnt = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == '1') {
                    cnt++;
                    grid[i][j] = '0';

                    for (int[] dir : directions) {
                        int new_i = i + dir[0];
                        int new_j = j + dir[1];

                        if (new_i >= 0 && new_i < m && new_j >= 0 && new_j < n && grid[new_i][new_j] == '1') {
                            if (union.union((i * n) + j, (new_i * n) + new_j)) {
                                cnt--;
                            }
                        }
                    }
                }
            }
        }

        return cnt;
    }
}
