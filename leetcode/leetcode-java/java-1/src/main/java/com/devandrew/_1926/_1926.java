package com.devandrew._1926;

import java.util.*;

public class _1926 {
    public int nearestExit(char[][] maze, int[] entrance) {
        int m = maze.length;
        int n = maze[0].length;
        Queue<List<Integer>> queue = new LinkedList<>();
        queue.add(Arrays.asList(entrance[0], entrance[1], 0));
        List<List<Integer>> directions = Arrays.asList(
                Arrays.asList(1, 0),
                Arrays.asList(0, 1),
                Arrays.asList(0, -1),
                Arrays.asList(-1, 0)
        );

        while(!queue.isEmpty()) {
            List<Integer> cell = queue.poll();
            int cellX = cell.get(0);
            int cellY = cell.get(1);
            int dist = cell.get(2);
            if (maze[cellX][cellY] == '.') {
                maze[cellX][cellY] = '+';

                if (dist != 0 && (cellX == 0 || cellX == m - 1 || cellY == 0 || cellY == n - 1)) {
                    return dist;
                }

                for (List<Integer> d : directions) {
                    int targetX = cellX + d.get(0);
                    int targetY = cellY + d.get(1);

                    if (targetX >= 0 && targetX < m && targetY >= 0 && targetY < n) {
                        queue.add(Arrays.asList(targetX, targetY, dist + 1));
                    }
                }
            }
        }

        return -1;
    }
}
