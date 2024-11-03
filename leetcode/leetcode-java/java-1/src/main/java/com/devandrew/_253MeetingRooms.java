package com.devandrew;

import java.util.*;

public class _253MeetingRooms {
    public int minMeetingRooms(int[][] intervals) {
        Map<Integer, Integer> map = new TreeMap<>();
        Arrays.sort(intervals, Comparator.comparingInt(a -> a[0]));

        PriorityQueue<Integer> endTimes = new PriorityQueue<>();

        for (int[] interval : intervals) {
            if (!endTimes.isEmpty() && endTimes.peek() <= interval[0]) {
                endTimes.poll();
            }

            endTimes.add(interval[1]);
        }

        return endTimes.size();
    }
}
