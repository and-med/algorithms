package com.devandrew._2281;

import com.devandrew.helpers.Parser;

public class _2281Test {
    public static void run() {
        Parser parser = new Parser("2281");

        parser.withInputAndOutputProblemAdjusted(scanner -> {
            int[] strength = parser.parseArray(scanner.nextLine());

            _2281SpeedUp solution = new _2281SpeedUp();
            return solution.totalStrength(strength);
        });
    }
}
