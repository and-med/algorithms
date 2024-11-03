package com.devandrew;

public class _528RandomPickWithIndex {
    private final int[] sums;
    private int sum;

    public _528RandomPickWithIndex(int[] w) {
        this.sums = new int[w.length];
        this.sum = 0;
        for (int i = 0; i < w.length; i++) {
            this.sum += w[i];
            sums[i] = this.sum;
        }
    }

    public int getFirstIndexOfGreaterThan(int l, int r, int x) {
        if (l == r) {
            return l;
        }

        int m = (l + r) / 2;

        if (x < sums[m]) {
            return getFirstIndexOfGreaterThan(l, m, x);
        } else {
            return getFirstIndexOfGreaterThan(m + 1, r, x);
        }
    }

    public int pickIndex() {
        int random = (int)Math.floor(Math.random() * sum);
        return getFirstIndexOfGreaterThan(0, sums.length - 1, random);
    }
}
