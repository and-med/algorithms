package com.devandrew._200NumberOfIslands;

import com.devandrew.helpers.Parser;

public class _200Test {
    public static void run() {
        Parser parser = new Parser("200NumberOfIslands");

        parser.withInputAndOutputProblemAdjusted(scanner -> {
            char[][] grid = parser.parseCharArrays(scanner.nextLine());

            _200UnionFind solution = new _200UnionFind();
            return solution.numIslands(grid);
        });
    }
}
