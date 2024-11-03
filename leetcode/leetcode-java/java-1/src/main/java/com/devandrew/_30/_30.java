package com.devandrew._30;

import java.util.*;

public class _30 {
    private class SlidingWindow {
        private final int wordLength;
        private final Map<String, Integer> frequencies;
        private final Map<String, Integer> slidingWindow;
        private final Set<String> visited;
        private int currWindowLength;
        private final int targetLength;
        private final int size;

        public SlidingWindow(String[] words) {
            this.frequencies = new HashMap<>();

            for (String word: words) {
                frequencies.put(word, frequencies.getOrDefault(word, 0) + 1);
            }

            this.slidingWindow = new HashMap<>(frequencies);
            this.visited = new HashSet<>();

            this.wordLength = words[0].length();
            this.currWindowLength = 0;
            this.size = words.length;
            this.targetLength = wordLength * size;
        }

        public void resetWindow() {
            this.currWindowLength = 0;
            for (String word: visited) {
                slidingWindow.put(word, frequencies.get(word));
            }
            visited.clear();
        }

        private void updateSlidingWindow(String word, int sign) {
            this.currWindowLength = currWindowLength + sign;
            int currCount = this.slidingWindow.getOrDefault(word, 0);
            this.visited.add(word);

            if (currCount - sign != 0) {
                this.slidingWindow.put(word, currCount - sign);
            } else {
                this.slidingWindow.remove(word);
            }
        }

        private String getWord(String s, int i) {
            return s.substring(i, i + wordLength);
        }

        public boolean addWord(String s, int idx) {
            boolean result = false;

            if (s.length() - idx < wordLength) {
                this.resetWindow();
                return false;
            }

            String word = getWord(s, idx);

            if (frequencies.containsKey(word)) {
                this.updateSlidingWindow(word, 1);

                if (this.slidingWindow.size() == 0) {
                    result = true;
                }
                if (this.currWindowLength == size) {
                    this.updateSlidingWindow(getWord(s, getWindowStart(idx)), -1);
                }
            } else {
                this.resetWindow();
            }

            return result;
        }

        public int getWindowStart(int lastWordIndex) {
            return lastWordIndex + wordLength - targetLength;
        }
    }

    public List<Integer> findSubstring(String s, String[] words) {
        SlidingWindow window = new SlidingWindow(words);

        List<Integer> res = new ArrayList<>();

        for (int i = 0; i < window.wordLength; i++) {
            window.resetWindow();
            for (int j = i; j < s.length(); j += window.wordLength) {
                if (window.addWord(s, j)) {
                    res.add(window.getWindowStart(j));
                }
            }
        }

        return res;
    }
}
