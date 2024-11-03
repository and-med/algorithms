package com.devandrew._872;

import com.devandrew.helpers.Parser;

public class _872Test {
    public static void run() {
        Parser parser = new Parser("872");

        parser.withInputAndOutputProblemAdjusted(scanner -> {
            Integer[] tree1 = parser.parseNullableArray(scanner.nextLine());
            Integer[] tree2 = parser.parseNullableArray(scanner.nextLine());

            _872 solution = new _872();
            return solution.leafSimilar(_872.fromArray(tree1, 0), _872.fromArray(tree2, 0));
        });
    }
}
