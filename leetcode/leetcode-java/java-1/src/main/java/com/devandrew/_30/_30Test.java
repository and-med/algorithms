package com.devandrew._30;

import com.devandrew._79._79;
import com.devandrew.helpers.Parser;

public class _30Test {
    public static void run() {
        Parser parser = new Parser("30");

        parser.withInputAndOutputProblemAdjusted(scanner -> {
            String s = scanner.nextLine();
            String[] words = parser.parseStringArray(scanner.nextLine());

            _30 solution = new _30();
            return solution.findSubstring(s, words);
        });
    }
}
