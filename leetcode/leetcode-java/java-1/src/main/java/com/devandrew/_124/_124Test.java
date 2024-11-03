package com.devandrew._124;

import com.devandrew.helpers.Parser;
import com.devandrew.utils.TreeNode;

public class _124Test {
    public static void run() {
        Parser parser = new Parser("124");

        parser.withInputAndOutputProblemAdjusted(scanner -> {
            Integer[] tree = parser.parseNullableArray(scanner.nextLine());

            _124 solution = new _124();
            return solution.maxPathSum(TreeNode.fromArray(tree, 0));
        });
    }
}
