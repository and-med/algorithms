package com.devandrew;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;

public class _489Robot {
    // This is the robot's control interface.
    // You should not implement it, or speculate about its implementation
    public static interface Robot {
        // Returns true if the cell in front is open and robot moves into the cell.
        // Returns false if the cell in front is blocked and robot stays in the current cell.
        boolean move();
        // Robot will stay in the same cell after calling turnLeft/turnRight.
        // Each turn will be 90 degrees.
        void turnLeft();
        void turnRight();
        // Clean the current cell.
        void clean();
    }

    public enum Direction {
        UP(0, -1, 0),
        RIGHT(1, 0, 1),
        BOTTOM(2, 1, 0),
        LEFT(3, 0, -1);

        private static final int Length = 4;

        private final int code;
        private final int xChange;
        private final int yChange;

        Direction(int code, int xChange, int yChange) {
            this.code = code;
            this.xChange = xChange;
            this.yChange = yChange;
        }

        public static Direction turnLeft(Direction direction) {
            int newCode = (direction.code - 1) < 0 ? Length - 1 : direction.code - 1;
            return Direction.values()[newCode];
        }

        public static Direction turnRight(Direction direction) {
            int newCode = (direction.code + 1) == Length ? 0 : direction.code + 1;
            return Direction.values()[newCode];
        }

        public int getX() {
            return xChange;
        }

        public int getY() {
            return yChange;
        }
    }

    private int[][] memo;
    private int m = 202;
    private int n = 402;
    private int[][] directions = new int[][] { { -1, 0 }, { 0, 1 }, { 1, 0 }, { 0, -1 } };

    private void cleanRoomDfs(Robot robot, int i, int j, int direction) {
        robot.clean();
        memo[i][j] = 1;

        for (int k = 0; k < directions.length; k++) {
            int d = (direction + k) % directions.length;
            int newI = i + directions[d][0];
            int newJ = j + directions[d][1];

            if (memo[newI][newJ] == -1 && robot.move()) {
                cleanRoomDfs(robot, newI, newJ, d);

                robot.turnLeft();
                robot.turnLeft();
                robot.move();
                robot.turnLeft();
                robot.turnLeft();
            } else {
                memo[newI][newJ] = 0;
            }

            robot.turnRight();
        }
    }

    public void cleanRoom(Robot robot) {
        memo = new int[m][n];
        for (int i = 0; i < m; i++) {
            Arrays.fill(memo[i], -1);
        }

        new ArrayList<Integer>().sort(Collections.reverseOrder());

        cleanRoomDfs(robot, m/2, n/2, 0);
    }
}
