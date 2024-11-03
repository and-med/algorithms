package com.devandrew;

import com.devandrew.helpers.Parser;

public class _1696JumpGameTest {
    public static void run() {
        Parser parser = new Parser("1696");

        parser.withInput(scanner -> {
            parser.createOutputFile();

            while (scanner.hasNext()) {
                int[] nums = parser.parseArray(scanner.nextLine());
                int k = Integer.parseInt(scanner.nextLine());

                _1696JumpGameDeque solution = new _1696JumpGameDeque();
                int sol = solution.maxResult(nums, k);
                parser.appendToOutput(sol);
            }
        });
    }
}
