import java.util.*;

class Solution {
    public int nearestExit(char[][] maze, int[] entrance) {
        int m = maze.length;
        int n = maze[0].length;
        var stack = new Stack<List<Integer>>();
        var visited = new HashSet<Pair<Integer, Integer>>();
        stack.push(Arrays.asList(entrance[0], entrance[1], 0));
        var directions = Arrays.asList(
            Arrays.asList(1, 0),
            Arrays.asList(0, 1),
            Arrays.asList(0, -1),
            Arrays.asList(-1, 0)
        );

        while(!stack.isEmpty()) {
            var cell = stack.pop();
            int cellX = cell.get(0);
            int cellY = cell.get(1);
            int dist = cell.get(2);
            
            if (maze[cellX][cellY] != '+') {
                maze[cellX][cellY] = '+';

                if (cellX == 0 || cellX == m - 1 || cellY == 0 || cellY == n - 1 && dist != 0) {
                    return dist;
                }

                for (var d: directions) {
                    int targetX = cellX + d.get(0);
                    int targetY = cellY + d.get(1);

                    if (targetX >= 0 && targetX < m && targetY >= 0 && targetY < n) {
                        stack.push(Arrays.asList(targetX, targetY, dist + 1));
                    }
                }
            }
        }

        return -1;
    }
}