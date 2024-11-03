package com.devandrew;

import java.util.Deque;
import java.util.LinkedList;

public class _1696JumpGameDeque {
    public int maxResult(int[] nums, int k) {
        int[] maxScores = new int[nums.length];
        Deque<Integer> deque = new LinkedList<>();

        maxScores[0] = nums[0];
        deque.offerLast(0);

        for (int i = 1; i < nums.length; i++) {
            if (deque.peekFirst() < i - k) {
                deque.pollFirst();
            }

            maxScores[i] = nums[i] + maxScores[deque.peekFirst()];

            while (!deque.isEmpty() && maxScores[i] > maxScores[deque.peekLast()]) {
                deque.pollLast();
            }

            deque.offerLast(i);
        }

        return maxScores[nums.length - 1];
    }
}
