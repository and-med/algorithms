package com.devandrew;

import java.util.HashSet;
import java.util.Set;

public class _2272SubstringWithLargestVariance {
    public int largestVariance(String s) {
        Set<Character> unique = new HashSet<>();

        for (char c : s.toCharArray()) {
            unique.add(c);
        }

        int ans = 0;

        for (Character a : unique) {
            for (Character b : unique) {
                if (a != b) {
                    int curr_var = 0;
                    boolean has_b = false;
                    boolean first_b = false;

                    for (Character c : s.toCharArray()) {
                        if (a == c) {
                            curr_var++;
                        }
                        if (b == c) {
                            has_b = true;
                            boolean is_second_b = first_b && curr_var >= 0;
                            if (is_second_b) {
                                // Do not decrease counter here. Example "abbbab", at position i = 4 sum should be 2 not 1
                                first_b = false;
                            } else if (--curr_var < 0) {
                                first_b = true;
                                curr_var = -1;
                            }
                        }

                        ans = Math.max(ans, has_b ? curr_var : 0);
                    }
                }
            }
        }

        return ans;
    }
}
