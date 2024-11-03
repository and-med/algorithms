package com.devandrew;

public class _135Candy {
    public int count(int n) {
        return (n * (n + 1)) / 2;
    }

    public int candy(int[] ratings) {
        int answer = 0;
        int cntUp = 0;
        int cntDown = 0;
        int oldSlope = 0;

        for (int i = 1; i < ratings.length; i++) {
            int newSlope = Integer.compare(ratings[i], ratings[i - 1]);

            if ((oldSlope == 1 && newSlope == 0) || (oldSlope == -1 && newSlope >= 0)) {
                answer += count(cntUp) + count(cntDown) + Math.max(cntUp, cntDown);
                cntUp = 0;
                cntDown = 0;
            }

            if (newSlope == 1) {
                cntUp++;
            } else if (newSlope == -1) {
                cntDown++;
            } else {
                answer++;
            }

            oldSlope = newSlope;
        }

        answer += count(cntUp) + count(cntDown) + Math.max(cntUp, cntDown) + 1;

        return answer;
    }
}
