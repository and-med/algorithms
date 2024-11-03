package com.devandrew._2281;

import java.util.Arrays;
import java.util.Stack;

public class _2281SpeedUp {
    public static long MOD = 1000000007;

    class SumOfSubarraysCalculator {
        private int[] strength;

        public SumOfSubarraysCalculator(int[] strength) {
            this.strength = strength;
        }

        public long calculate(int l, int r, int m) {
            int leftLen = m - l + 1;
            int rightLen = r - m + 1;

            long sum = 0;
            for (int i = l, curr = rightLen; i <= m; i++, curr += rightLen) {
                sum = (sum + (long) curr * strength[i]) % MOD;
            }

            for (int i = r, curr = leftLen; i > m; i--, curr += leftLen) {
                sum = (sum + (long) curr * strength[i]) % MOD;
            }

            return sum;
        }
    }

    public int totalStrength(int[] strength) {
        int[] left_less = new int[strength.length + 1];
        Arrays.fill(left_less, -1);
        int[] right_less = new int[strength.length + 1];
        Arrays.fill(right_less, strength.length);

        Stack<Integer> st = new Stack<>();
        for (int i = 0; i < strength.length; i++) {
            while (!st.isEmpty() && strength[st.peek()] >= strength[i]) {
                st.pop();
            }
            if (!st.isEmpty()) {
                left_less[i] = st.peek();
            }

            st.push(i);
        }

        // 3 1 5 1 4 3 1 3 3

        st.clear();
        for (int i = strength.length - 1; i >= 0; i--) {
            while (!st.isEmpty() && strength[st.peek()] > strength[i]) {
                st.pop();
            }
            if (!st.isEmpty()) {
                right_less[i] = st.peek();
            }

            st.push(i);
        }

        long ans = 0;

        SumOfSubarraysCalculator calculator = new SumOfSubarraysCalculator(strength);

        for (int i = 0; i < strength.length; i++) {
            int left_cap = left_less[i];
            int right_cap = right_less[i];

            long sum = calculator.calculate(left_cap + 1, right_cap - 1, i);
            ans = (ans + sum * strength[i]) % MOD;
        }

        return (int)ans;
    }
}
