package com.devandrew._279;

import com.devandrew.helpers.Parser;

public class _279Test {
    public static void run() {
        Parser parser = new Parser("279");

        parser.withInputAndOutputProblemAdjusted(scanner -> {
            int n = Integer.parseInt(scanner.nextLine());

            _279_dp_bottom_up solution = new _279_dp_bottom_up();
            return solution.numSquares(n);
        });
    }
}
