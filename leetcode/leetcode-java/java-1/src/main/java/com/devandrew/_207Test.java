package com.devandrew;

import com.devandrew.helpers.Parser;

public class _207Test {
    public static void run() {
        Parser parser = new Parser("207");

        parser.withInput((scanner) -> {
            int numCourses = Integer.parseInt(scanner.nextLine());
            int[][] prerequisites = parser.parseArrays(scanner.nextLine());
            _207CourseSchedule solution = new _207CourseSchedule();
            boolean res = solution.canFinish(numCourses, prerequisites);

            parser.outputSolution(res);
        });
    }
}
