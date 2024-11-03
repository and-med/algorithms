package com.devandrew._1307;

import com.devandrew.helpers.Parser;

public class _1307Test {
    public static void run() {
        Parser parser = new Parser("1307");

        parser.withInputAndOutputProblemAdjusted(scanner -> {
            String[] words = parser.parseStringArray(scanner.nextLine());
            String result = scanner.nextLine();

            _1307 solution = new _1307();
            return solution.isSolvable(words, result);
        });
    }
}
