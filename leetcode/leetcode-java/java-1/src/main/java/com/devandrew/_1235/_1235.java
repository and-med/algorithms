package com.devandrew._1235;

import org.apache.commons.lang3.tuple.Pair;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.IntStream;

public class _1235 {
    private class SmartAssDp {
        private List<Pair<Integer, Integer>> dp;

        public SmartAssDp(int n) {
            this.dp = new ArrayList<>(n);
        }

        private int binarySearch(int l, int r, int v) {
            if (l == r) {
                return l - 1;
            }

            int m = l + (r - l) / 2;
            if (dp.get(m).getKey() == v) {
                return m;
            } else if (dp.get(m).getKey() < v) {
                return binarySearch(m + 1, r, v);
            }
            return binarySearch(l, m, v);
        }

        private void sortLast() {
            for (int i = dp.size() - 2; i >= 0; i--) {
                if (dp.get(i).getKey() > dp.get(i + 1).getKey()) {
                    Pair<Integer, Integer> temp = dp.get(i + 1);
                    dp.set(i + 1, dp.get(i));
                    dp.set(i, temp);
                }
            }
        }

        public int findLowerBound(int v) {
            return binarySearch(0, dp.size(), v);
        }

        public int getProfitOrZero(int it) {
            if (it != -1 && it != dp.size()) {
                return dp.get(it).getValue();
            }

            return 0;
        }

        public void update(int key, int value, int it) {
            if (it != -1 && it != dp.size()) {
                if (dp.get(it).getKey() == key) {
                    dp.set(it, Pair.of(key, value));
                    return;
                }
            }
            dp.add(Pair.of(key, value));
            sortLast();
        }

        public int getMaxProfit() {
            int max = Integer.MIN_VALUE;
            for (int i = 0; i < dp.size(); i++) {
                max = Math.max(max, dp.get(i).getValue());
            }
            return max;
        }
    }

    public int jobScheduling(int[] startTime, int[] endTime, int[] profit) {
        int[] sorted = IntStream.range(0, startTime.length)
                .boxed()
                .sorted((l, r) -> {
                    int et = Integer.compare(endTime[l], endTime[r]);
                    if (et != 0) {
                        return et;
                    }
                    int st = Integer.compare(startTime[l], startTime[r]);
                    if (st != 0) {
                        return st;
                    }
                    return Integer.compare(profit[l], profit[r]);
                })
                .mapToInt(i -> i)
                .toArray();

        SmartAssDp dp = new SmartAssDp(startTime.length);

        for (int i : sorted) {
            int eidx = dp.findLowerBound(endTime[i]);
            int eProfit = dp.getProfitOrZero(eidx);
            int sProfit = dp.getProfitOrZero(dp.findLowerBound(startTime[i]));

            dp.update(endTime[i], Math.max(eProfit, sProfit + profit[i]), eidx);
        }

        return dp.getMaxProfit();
    }
}
