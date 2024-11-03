package com.devandrew;

import com.devandrew.helpers.Parser;

public class _1152Test {
    public static void run() {
        Parser parser = new Parser("1152");

        parser.withInputAndOutput(scanner -> {
            String[] username = parser.parseStringArray(scanner.nextLine());
            int[] timestamp = parser.parseArray(scanner.nextLine());
            String[] website = parser.parseStringArray(scanner.nextLine());

            _1152AnalyzeUserWebsiteVisit solution = new _1152AnalyzeUserWebsiteVisit();
            return solution.mostVisitedPattern(username, timestamp, website);
        });
    }
}
