package com.devandrew;

import java.util.ArrayList;
import java.util.List;

public class _5PalindromicTest {
    public String longestPalindrome(String s) {
        List<Integer> longestPalindrome = new ArrayList<>(s.length());

        for (int i = 0; i < s.length(); i++) {
            longestPalindrome.add(1);
        }

        for (int i = 1; i < s.length(); i++) {
            for (int j = 0; j < i; j++) {
                if (s.charAt(j) == s.charAt(i) && (longestPalindrome.get(j + 1) == (i - j - 1) || (i - j == 1))) {
                    longestPalindrome.set(j, (i - j + 1));
                }
            }
        }

        Integer max = 0;
        Integer idx = 0;
        for (int i = 0; i < s.length(); i++) {
            if (longestPalindrome.get(i) > max) {
                max = longestPalindrome.get(i);
                idx = i;
            }
        }

        return s.substring(idx, idx + max);
    }
}
