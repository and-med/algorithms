package com.devandrew._279;

public class _279_dp {
    private int[] dp;

    private int numSquaresRec(int n) {
        if (n < 4) {
            return n;
        } else if (dp[n] != 0) {
            return dp[n];
        }

        int ans = n;

        for (int i = 1; i * i <= n; i++) {
            ans = Math.min(ans, 1 + numSquaresRec(n - i * i));
        }

        dp[n] = ans;

        return ans;
    }

    public int numSquares(int n) {
        this.dp = new int[n + 1];
        return numSquaresRec(n);
    }
}
