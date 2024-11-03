package com.devandrew;

import com.devandrew.helpers.Parser;

public class _218Test {
    public static void run() {
        Parser parser = new Parser("218");

        parser.withInputAndOutput(scanner -> {
            int[][] buildings = parser.parseArrays(scanner.nextLine());

            _218TheCitySkyline solution = new _218TheCitySkyline();
            return solution.getSkyline(buildings);
        });
    }
}
