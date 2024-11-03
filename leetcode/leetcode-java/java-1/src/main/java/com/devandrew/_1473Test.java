package com.devandrew;

import com.devandrew.helpers.Parser;

public class _1473Test {
    public static void run() {
        Parser parser = new Parser("1473");

        parser.withInput(scanner -> {
            parser.createOutputFile();

            while(scanner.hasNext()) {
                int[] houses = parser.parseArray(scanner.nextLine());
                int[][] cost = parser.parseArrays(scanner.nextLine());
                int m = Integer.parseInt(scanner.nextLine());
                int n = Integer.parseInt(scanner.nextLine());
                int target = Integer.parseInt(scanner.nextLine());

                _1473RowOfHouses solution = new _1473RowOfHouses();
                int sol = solution.minCost(houses, cost, m, n, target);
                parser.appendToOutput(sol);
            }
        });
    }
}
