package com.devandrew._1926;

import com.devandrew.helpers.Parser;

public class _1926Test {
    public static void run() {
        Parser parser = new Parser("1926");

        parser.withInputAndOutputProblemAdjusted(scanner -> {
            char[][] maze = parser.parseCharArrays(scanner.nextLine());
            int[] entrance = parser.parseArray(scanner.nextLine());

            _1926 solution = new _1926();
            return solution.nearestExit(maze, entrance);
        });
    }
}
