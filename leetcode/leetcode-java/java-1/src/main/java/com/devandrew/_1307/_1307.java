package com.devandrew._1307;

import java.util.*;

public class _1307 {
    private int[][] words;
    private int[] result;
    private int[] map;
    private boolean[] used;
    private boolean[] cannot_be_zero;

    public boolean isSolvable(String[] words, String result) {
        Map<Character, Integer> codes = new HashMap<>();
        this.words = encode_words(words, codes);
        this.result = encode_word(result, codes);
        this.map = new int[10];
        this.used = new boolean[10];
        this.cannot_be_zero = build_cannot_be_zeroes();
        Arrays.fill(map, -1);
        int max_len = Arrays.stream(words).mapToInt(String::length).max().orElse(0);

        if (codes.size() > 10 || max_len > result.length()) {
            return false;
        }

        return isSolvable(0, 0, 0);
    }

    public boolean isSolvable(int row, int col, int carry) {
        if (col == result.length) {
            return carry == 0;
        }

        if (row == words.length) {
            for (int[] word : words) {
                carry += col < word.length ? map[word[col]] : 0;
            }
            int rem = carry % 10;

            if (rem == 0 && cannot_be_zero[result[col]]) {
                return false;
            }

            if (map[result[col]] == -1 && !used[rem]) {
                map[result[col]] = rem;
                used[rem] = true;
                boolean ans = isSolvable(0, col + 1, carry / 10);
                map[result[col]] = -1;
                used[rem] = false;
                return ans;
            }

            return map[result[col]] == rem && isSolvable(0, col + 1, carry / 10);
        }

        if (col < words[row].length && map[words[row][col]] == -1) {
            int start = cannot_be_zero[words[row][col]] ? 1 : 0;
            for (int i = start; i < 10; i++) {
                if (!used[i]) {
                    map[words[row][col]] = i;
                    used[i] = true;

                    if (isSolvable(row + 1, col, carry)) {
                        return true;
                    }

                    map[words[row][col]] = -1;
                    used[i] = false;
                }
            }
            return false;
        }

        return isSolvable(row + 1, col, carry);
    }

    private int[][] encode_words(String[] words, Map<Character, Integer> codes) {
        int[][] encoded_words = new int[words.length][];
        for (int i = 0; i < words.length; i++) {
            encoded_words[i] = encode_word(words[i], codes);
        }
        return encoded_words;
    }

    private int[] encode_word(String word, Map<Character, Integer> codes) {
        int[] result = new int[word.length()];

        for (int i = word.length() - 1; i >= 0; i--) {
            Integer code = codes.get(word.charAt(i));
            result[word.length() - i - 1] = code != null ? code : codes.size();
            if (code == null) {
                codes.put(word.charAt(i), codes.size());
            }
        }

        return result;
    }

    private boolean[] build_cannot_be_zeroes() {
        boolean[] ans = new boolean[10];

        for (int[] word: words) {
            if (word.length > 1) {
                ans[word[word.length - 1]] = true;
            }
        }

        if (result.length > 1) {
            ans[result[result.length - 1]] = true;
        }

        return ans;
    }

    private List<Character> codes;
    private String[] originalWords;
    private String originalResult;

    private String getCurrentEncodingStr(int[] map) {
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < words.length; i++) {
            for (int j = 0; j < words[i].length; j++) {
                Character c = originalWords[i].charAt(j);
                int code = codes.indexOf(c);

                if (map[code] == -1) {
                    sb.append(c);
                } else {
                    sb.append(map[code]);
                }
            }
            if (i != words.length - 1) {
                sb.append(" + ");
            }
        }
        sb.append(" = ");
        for (int i = 0; i < originalResult.length(); i++) {
            Character c = originalResult.charAt(i);
            int code = codes.indexOf(c);

            if (map[code] == -1) {
                sb.append(c);
            } else {
                sb.append(map[code]);
            }
        }
        return sb.toString();
    }
}
