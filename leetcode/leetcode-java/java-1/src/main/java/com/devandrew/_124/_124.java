package com.devandrew._124;

import com.devandrew.utils.TreeNode;

import java.util.LinkedList;
import java.util.Queue;

public class _124 {
    private static int MAX_SIZE = 30000;
    private static int MAX_EDGES = 3;
    private static int EMPTY_EDGE = -1;
    private int[][] graph;
    private int[] weights;
    private int size;

    public int maxPathSum(TreeNode root) {
        this.graph = new int[MAX_SIZE][MAX_EDGES];
        this.weights = new int[MAX_SIZE];
        this.size = 0;
        for (int i = 0; i < MAX_SIZE; i++) {
            for (int j = 0; j < MAX_EDGES; j++) {
                graph[i][j] = EMPTY_EDGE;
            }
        }
        buildGraph(root, 0);
        System.out.println(size);
        int max = Integer.MIN_VALUE;
        for (int i = 0; i < size; i++) {
            max = Math.max(max, dfsMaxPath(i));
        }
        return max;
    }

    private int dfsMaxPath(int s) {
        Queue<int[]> queue = new LinkedList<>();
        boolean[] visited = new boolean[size];
        queue.add(new int[] { s, weights[s] });
        int max = Integer.MIN_VALUE;

        while (!queue.isEmpty()) {
            int[] edge = queue.poll();
            visited[edge[0]] = true;
            max = Math.max(max, edge[1]);

            for (int e: graph[edge[0]]) {
                if (e != EMPTY_EDGE && !visited[e]) {
                    queue.add(new int[] { e, edge[1] + weights[e] });
                }
            }
        }

        return max;
    }

    private void buildGraph(TreeNode root, int parentId) {
        if (root != null) {
            int rootId = size++;
            weights[rootId] = root.val;
            addEdge(parentId, rootId);
            addEdge(rootId, parentId);
            buildGraph(root.left, rootId);
            buildGraph(root.right, rootId);
        }
    }

    private void addEdge(int leftId, int rightId) {
        for (int i = 0; i < MAX_EDGES; i++) {
            if (graph[leftId][i] == EMPTY_EDGE) {
                graph[leftId][i] = rightId;
                break;
            }
        }
    }
}
