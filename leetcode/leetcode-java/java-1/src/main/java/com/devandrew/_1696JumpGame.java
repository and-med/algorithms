package com.devandrew;

import java.util.*;

public class _1696JumpGame {
    public class MaxHeap {
        private final int[] heap;
        private int n;

        public MaxHeap(int n) {
            this.heap = new int[n];
            Arrays.fill(heap, Integer.MIN_VALUE);

            this.n = 0;
        }

        public boolean isEmpty() {
            return n == 0;
        }

        public void add(int val) {
            this.insertLast(val);
            heapBubbleUp(n - 1);
        }

        public void remove(int val) {
            int index = 0;

            while (heap[index] != val) {
                index++;
            }

            swap(index, n - 1);
            removeLast();

            heapBubbleUp(index);
            heapify(index);
        }

        public int getMax() {
            if (isEmpty()) {
                throw new IllegalStateException("Heap is empty.");
            }

            return heap[0];
        }

        private void heapBubbleUp(int i) {
            while (i != 0 && heap[i] > heap[parent(i)]) {
                this.swap(i, parent(i));
                i = parent(i);
            }
        }

        private void heapify(int i) {
            int left = leftChild(i);
            int right = rightChild(i);

            int largest = i;
            if (left < n && heap[left] > heap[largest]) {
                largest = left;
            }
            if (right < n && heap[right] > heap[largest]) {
                largest = right;
            }

            if (largest != i) {
                swap(largest, i);
                heapify(largest);
            }
        }

        private int parent(int i) {
            return (i - 1) / 2;
        }

        private int leftChild(int i) {
            return 2 * i + 1;
        }

        private int rightChild(int i) {
            return 2 * i + 2;
        }

        private void insertLast(int val) {
            heap[n] = val;
            n++;
        }

        private void removeLast() {
            n--;
        }

        private void swap(int left, int right) {
            int temp = heap[left];
            heap[left] = heap[right];
            heap[right] = temp;
        }
    }

    public int maxResult(int[] nums, int k) {
        int[] maxScores = new int[nums.length];
        MaxHeap maxHeap = new MaxHeap(k + 1);

        for (int i = 0; i < nums.length; i++) {
            maxScores[i] = nums[i] + (maxHeap.isEmpty() ? 0 : maxHeap.getMax());

            maxHeap.add(maxScores[i]);

            if (i - k >= 0) {
                maxHeap.remove(maxScores[i - k]);
            }
        }

        return maxScores[nums.length - 1];
    }
}
