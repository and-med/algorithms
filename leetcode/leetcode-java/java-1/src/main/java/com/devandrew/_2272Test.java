package com.devandrew;

import com.devandrew.helpers.Parser;

public class _2272Test {
    public static void run() {
        Parser parser = new Parser("2272");

        parser.withInputAndOutput(scanner -> {
            String s = scanner.nextLine();
            s = s.substring(1, s.length() - 1);

            _2272SubstringWithLargestVariance solution = new _2272SubstringWithLargestVariance();
            return solution.largestVariance(s);
        });
    }
}
