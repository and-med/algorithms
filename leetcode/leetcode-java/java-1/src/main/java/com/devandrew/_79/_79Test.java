package com.devandrew._79;

import com.devandrew.helpers.Parser;

public class _79Test {
    public static void run() {
        Parser parser = new Parser("79");

        parser.withInputAndOutputProblemAdjusted(scanner -> {
            char[][] board = parser.parseCharArrays(scanner.nextLine());
            String word = scanner.nextLine();

            _79 solution = new _79();
            return solution.exist(board, word);
        });
    }
}
