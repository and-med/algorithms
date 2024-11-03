package com.devandrew;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class _828CountUniqueStrings {
    public int uniqueLetterString(String s) {
        int[] prevSameIndex = new int[s.length()];
        Arrays.fill(prevSameIndex, -1);
        Map<Character, Integer> lastIndexOf = new HashMap<>();

        for (int i = 0; i < s.length(); i++) {
            if (lastIndexOf.containsKey(s.charAt(i))) {
                prevSameIndex[i] = lastIndexOf.get(s.charAt(i));
            }

            lastIndexOf.put(s.charAt(i), i);
        }

        lastIndexOf.clear();

        int[] nextSameIndex = new int[s.length()];
        Arrays.fill(nextSameIndex, s.length());

        for (int i = s.length() - 1; i >= 0; i--) {
            if (lastIndexOf.containsKey(s.charAt(i))) {
                nextSameIndex[i] = lastIndexOf.get(s.charAt(i));
            }

            lastIndexOf.put(s.charAt(i), i);
        }

        int cnt = 0;
        for (int i = 0; i < s.length(); i++) {
            cnt += (i - prevSameIndex[i]) * (nextSameIndex[i] - i);
        }

        return cnt;
    }
}
