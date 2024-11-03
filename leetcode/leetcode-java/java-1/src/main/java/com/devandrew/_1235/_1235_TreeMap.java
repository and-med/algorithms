package com.devandrew._1235;

import java.util.Collections;
import java.util.Map;
import java.util.TreeMap;
import java.util.stream.IntStream;

public class _1235_TreeMap {
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

        TreeMap<Integer, Integer> dp = new TreeMap<>();

        for (int i : sorted) {
            Map.Entry<Integer, Integer> endProfitEntry = dp.floorEntry(endTime[i]);
            int eProfit = endProfitEntry == null ? 0 : endProfitEntry.getValue();
            Map.Entry<Integer, Integer> startProfitEntry = dp.floorEntry(startTime[i]);
            int sProfit = startProfitEntry == null ? 0 : startProfitEntry.getValue();

            dp.put(endTime[i], Math.max(eProfit, sProfit + profit[i]));
        }

        return Collections.max(dp.values());
    }
}
