package com.devandrew._937;

import com.devandrew.helpers.Parser;

public class _937Test {
    public static void run() {
        Parser parser = new Parser("937");

        parser.withInputAndOutputProblemAdjusted(scanner -> {
            String[] logs = parser.parseStringArray(scanner.nextLine());

            _937 solution = new _937();
            return solution.reorderLogFiles(logs);
        });
    }
}
