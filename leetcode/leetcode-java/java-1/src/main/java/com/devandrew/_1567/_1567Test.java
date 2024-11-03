package com.devandrew._1567;

import com.devandrew._937._937;
import com.devandrew.helpers.Parser;

public class _1567Test {
    public static void run() {
        Parser parser = new Parser("1567");

        parser.withInputAndOutputProblemAdjusted(scanner -> {
            int[] nums = parser.parseArray(scanner.nextLine());

            _1567 solution = new _1567();
            return solution.getMaxLen(nums);
        });
    }
}
