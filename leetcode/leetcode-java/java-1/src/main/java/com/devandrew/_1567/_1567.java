package com.devandrew._1567;

import java.util.Arrays;
import java.util.stream.Collectors;

public class _1567 {
    public int getMaxLen(int[] nums) {
        int ans = 0;
        int neg = 0;
        int pos = 0;

        System.out.println(Arrays.stream(nums).mapToObj(Integer::toString).collect(Collectors.joining(" ")));

        for (int i = 0; i < nums.length; i++) {
            if (nums[i] < 0) {
                int temp = neg;
                neg = pos + 1;
                pos = temp == 0 ? 0 : temp + 1;
            } else if (nums[i] > 0) {
                neg = neg == 0 ? 0 : neg + 1;
                pos = pos + 1;
            } else {
                neg = 0;
                pos = 0;
            }

            System.out.println(neg + " " + pos);

            ans = Math.max(ans, pos);
        }

        return ans;
    }
}
