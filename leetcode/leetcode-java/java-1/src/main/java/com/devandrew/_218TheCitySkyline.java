package com.devandrew;

import java.lang.reflect.Array;
import java.util.*;

public class _218TheCitySkyline {
    public class BuildingPoint {
        private int x;
        private int y;
        private boolean isStart;

        public BuildingPoint(int x, int y, boolean isStart) {
            this.x = x;
            this.y = y;
            this.isStart = isStart;
        }

        public int getX() {
            return x;
        }

        public int getY() {
            return y;
        }

        public boolean getIsStart() {
            return isStart;
        }

        public int compareTo(BuildingPoint b) {
            int xCompare = Integer.compare(x, b.x);

            if (xCompare == 0) {
                if (this.isStart != b.isStart) {
                    return this.isStart ? -1 : 1;
                }

                return this.isStart ? Integer.compare(b.y, this.y) : Integer.compare(this.y, b.y);
            }

            return xCompare;
        }

        @Override
        public String toString() {
            return "BuildingPoint{" +
                    "x=" + x +
                    ", y=" + y +
                    ", isStart=" + isStart +
                    '}';
        }
    }

    public List<List<Integer>> getSkyline(int[][] buildings) {
        List<List<Integer>> response = new ArrayList<>();
        List<BuildingPoint> buildingPoints = new ArrayList<>(buildings.length * 2);

        for (int i = 0; i < buildings.length; i++) {
            buildingPoints.add(new BuildingPoint(buildings[i][0], buildings[i][2], true));
            buildingPoints.add(new BuildingPoint(buildings[i][1], buildings[i][2], false));
        }

        buildingPoints.sort(BuildingPoint::compareTo);

        TreeMap<Integer, Integer> heights = new TreeMap<>();

        heights.put(0, 1);

        for (BuildingPoint curr : buildingPoints) {
            if (curr.isStart) {
                if (curr.y > heights.lastKey()) {
                    response.add(Arrays.asList(curr.x, curr.y));
                }

                heights.merge(curr.y, 1, Integer::sum);
            } else {
                int prevHeight = heights.lastKey();

                heights.computeIfPresent(curr.y, (key, count) -> {
                    if (count == 1) {
                        return null;
                    }
                    return count - 1;
                });

                if (heights.lastKey() != prevHeight) {
                    response.add(Arrays.asList(curr.x, heights.lastKey()));
                }
            }
        }

        return response;
    }
}
