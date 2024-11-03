package com.devandrew._279;

public class _279_dp_bottom_up {
    public int numSquares(int n) {
        int[] dp = new int[n + 1];

        for(int i = 1; i <= n; i++) {
            dp[i] = i;

            for(int j = 1; j * j <= i; j++) {
                dp[i] = Math.min(dp[i], 1 + dp[i - (j * j)]);
            }
        }

        return dp[n];
    }
}
