package com.devandrew._279;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class _279 {
    private final List<Integer> squares = new ArrayList<>();
    private int n;
    private int currMin;

    public void minNumSquares(int i, int sum, int size) {
        if (i == squares.size() - 1) {
            currMin = Math.min(currMin, size + (n - sum));
            return;
        } else if (sum == n) {
            currMin = Math.min(currMin, size);
            return;
        }

        int square = squares.get(i);
        int times = (n - sum) / square;

        if (size + times >= currMin) {
            return;
        }

        while (times >= 0) {
            minNumSquares(i + 1, sum + times * square, size + times);
            times -= 1;
        }
    }

    private void generateSquares() {
        int i = 1;
        int square = 1;
        while (square <= n) {
            squares.add(square);
            i = i + 1;
            square = i * i;
        }
        Collections.reverse(squares);
    }

    public int numSquares(int n) {
        this.n = n;
        this.currMin = Integer.MAX_VALUE;
        this.generateSquares();

        minNumSquares(0, 0,0);

        return currMin;
    }
}
