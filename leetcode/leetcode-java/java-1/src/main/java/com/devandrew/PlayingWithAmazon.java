package com.devandrew;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class PlayingWithAmazon {
    /*
     * Complete the 'minimalHeaviestSetA' function below.
     *
     * The function is expected to return an INTEGER_ARRAY.
     * The function accepts INTEGER_ARRAY arr as parameter.
     */

    public static List<Integer> minimalHeaviestSetA(List<Integer> arr) {
        // Write your code here
        List<Integer> sorted = new ArrayList<>(arr);
        sorted.sort(Integer::compare);

        ArrayList<Integer> sums = new ArrayList<>();
        sums.add(sorted.get(0));

        for (int i = 1; i < sorted.size(); i++) {
            sums.add(sums.get(i - 1) + sorted.get(i));
        }

        ArrayList<Integer> ans = new ArrayList<>();

        int currSum = 0;
        for (int i = sorted.size() - 1; i >= 0; i--) {
            currSum += sorted.get(i);
            ans.add(sorted.get(i));

            if (currSum > sums.get(i - 1)) {
                return ans;
            }
        }

        ArrayList<Integer> newAns = new ArrayList<>();

        for (int i = ans.size() - 1; i >=0; i--) {
            newAns.add(ans.get(i));
        }

        return newAns;
    }
}
