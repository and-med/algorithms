package com.devandrew;

import com.devandrew.helpers.Parser;

public class _329Test {
    public static void run() {
        Parser parser = new Parser("329");

        parser.withInputAndOutput(scanner -> {
            int[][] matrix = parser.parseArrays(scanner.nextLine());

            _329LongestIncreasingPath solution = new _329LongestIncreasingPath();
            return solution.longestIncreasingPath(matrix);
        });
    }
}
