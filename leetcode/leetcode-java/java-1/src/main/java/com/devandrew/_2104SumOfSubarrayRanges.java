package com.devandrew;

import java.util.Arrays;
import java.util.Stack;

public class _2104SumOfSubarrayRanges {
    public long subArrayRanges(int[] nums) {
        Stack<Integer> st = new Stack<>();
        int[] prevLessThan = new int[nums.length];
        Arrays.fill(prevLessThan, -1);

        for (int i = 0; i < nums.length; i++) {
            while (!st.isEmpty() && nums[st.peek()] >= nums[i]) {
                st.pop();
            }
            if (!st.isEmpty()) {
                prevLessThan[i] = st.peek();
            }

            st.push(i);
        }

        st.clear();

        long sumMin = 0;

        for (int i = nums.length - 1; i >= 0; i--) {
            int nextLessThan = nums.length;

            while (!st.isEmpty() && nums[st.peek()] > nums[i]) {
                st.pop();
            }
            if (!st.isEmpty()) {
                nextLessThan = st.peek();
            }

            sumMin += ((long) (i - prevLessThan[i]) * (nextLessThan - i) - 1) * (long)nums[i];

            st.push(i);
        }

        st.clear();

        int[] prevMoreThan = new int[nums.length];
        Arrays.fill(prevMoreThan, -1);

        for (int i = 0; i < nums.length; i++) {
            while (!st.isEmpty() && nums[st.peek()] <= nums[i]) {
                st.pop();
            }
            if (!st.isEmpty()) {
                prevMoreThan[i] = st.peek();
            }

            st.push(i);
        }

        st.clear();

        long sumMax = 0;

        for (int i = nums.length - 1; i >= 0; i--) {
            int nextMoreThan = nums.length;

            while (!st.isEmpty() && nums[st.peek()] < nums[i]) {
                st.pop();
            }

            if (!st.isEmpty()) {
                nextMoreThan = st.peek();
            }

            sumMax += ((long) (i - prevMoreThan[i]) * (nextMoreThan - i) - 1) * (long)nums[i];

            st.push(i);
        }

        return sumMax - sumMin;
    }
}
