package com.devandrew;

import com.devandrew.helpers.Parser;

public class _489Test {
    public static void run() {
        Parser parser = new Parser("489");

        parser.withInputAndOutput(scanner -> {
            int[][] matrix = parser.parseArrays(scanner.nextLine());
            int i = Integer.parseInt(scanner.nextLine());
            int j = Integer.parseInt(scanner.nextLine());

            _489Robot.Robot robot = new RobotImpl(matrix, i, j);

            _489Robot solution = new _489Robot();
            solution.cleanRoom(robot);

            return matrix;
        });
    }

    public static class RobotImpl implements _489Robot.Robot {
        private final int[][] matrix;
        private int i;
        private int j;

        private _489Robot.Direction direction = _489Robot.Direction.UP;

        public RobotImpl(int[][] matrix, int i, int j) {
            this.matrix = matrix;
            this.i = i;
            this.j = j;
        }

        @Override
        public boolean move() {
            int newI = i + direction.getX();
            int newJ = j + direction.getY();

            if (newI >= 0 && newI < matrix.length
                    && newJ >= 0 && newJ < matrix[0].length
                    && matrix[newI][newJ] != 0) {
                i = newI;
                j = newJ;
                return true;
            }

            return false;
        }

        @Override
        public void turnLeft() {
            direction = _489Robot.Direction.turnLeft(direction);
        }

        @Override
        public void turnRight() {
            direction = _489Robot.Direction.turnRight(direction);
        }

        @Override
        public void clean() {
            matrix[i][j] = 2;
        }
    }
}
