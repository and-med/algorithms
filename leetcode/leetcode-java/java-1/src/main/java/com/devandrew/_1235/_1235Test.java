package com.devandrew._1235;

import com.devandrew._1567._1567;
import com.devandrew.helpers.Parser;

public class _1235Test {
    public static void run() {
        Parser parser = new Parser("1235");

        parser.withInputAndOutputProblemAdjusted(scanner -> {
            int[] startTime = parser.parseArray(scanner.nextLine());
            int[] endTime = parser.parseArray(scanner.nextLine());
            int[] profit = parser.parseArray(scanner.nextLine());

            _1235_TreeMap solution = new _1235_TreeMap();
            return solution.jobScheduling(startTime, endTime, profit);
        });
    }
}
