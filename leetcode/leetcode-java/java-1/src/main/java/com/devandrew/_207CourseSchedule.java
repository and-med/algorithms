package com.devandrew;

import java.util.*;

public class _207CourseSchedule {
    public boolean canFinish(int numCourses, int[][] prerequisites) {
        int[] indegree = new int[numCourses];
        Arrays.fill(indegree, 0);
        Map<Integer, List<Integer>> map = new HashMap<>();

        for (int[] prerequisite : prerequisites) {
            int i = prerequisite[0];
            int j = prerequisite[1];

            indegree[j]++;

            if (map.containsKey(i)) {
                map.get(i).add(j);
            } else {
                List<Integer> list = new LinkedList<>();
                list.add(j);
                map.put(i, list);
            }
        }

        Queue<Integer> queue = new LinkedList<>();

        for (int i = 0; i < indegree.length; i++) {
            if (indegree[i] == 0) {
                queue.add(i);
            }
        }

        int topologicalOrderCnt = 0;

        while (!queue.isEmpty()) {
            int vertex = queue.poll();

            ++topologicalOrderCnt;
            for (int connectedVertex : map.getOrDefault(vertex, new ArrayList<>())) {
                if (--indegree[connectedVertex] == 0) {
                    queue.add(connectedVertex);
                }
            }
        }

        return topologicalOrderCnt == numCourses;
    }
}
