package com.devandrew;

import com.devandrew.helpers.Parser;

public class _253Test {
    public static void run() {
        Parser parser = new Parser("253");

        parser.withInput(scanner -> {
            parser.createOutputFile();

            while(scanner.hasNext()) {
                int[][] intervals = parser.parseArrays(scanner.nextLine());

                _253MeetingRooms solution = new _253MeetingRooms();
                int sol = solution.minMeetingRooms(intervals);
                parser.appendToOutput(sol);
            }
        });
    }
}
