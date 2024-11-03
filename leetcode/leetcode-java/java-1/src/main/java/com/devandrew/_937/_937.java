package com.devandrew._937;

import java.util.Arrays;
import java.util.Comparator;

public class _937 {
    public String[] reorderLogFiles(String[] logs) {
        Comparator<String> compare = (left, right) -> {
            Integer leftSpaceIdx = left.indexOf(' ');
            Integer rightSpaceIdx = right.indexOf(' ');

            boolean isLeftDigit = left.charAt(leftSpaceIdx + 1) >= '0'
                    && left.charAt(leftSpaceIdx + 1) <= '9';
            boolean isRightDigit = right.charAt(rightSpaceIdx + 1) >= '0'
                    && right.charAt(rightSpaceIdx + 1) <= '9';

            if (isLeftDigit != isRightDigit) {
                return isLeftDigit ? 1 : -1;
            }

            if (isLeftDigit) {
                return 0;
            }

            String leftLog = left.substring(leftSpaceIdx + 1);
            String rightLog = right.substring(rightSpaceIdx + 1);

            int res = leftLog.compareTo(rightLog);

            if (res == 0) {
                String leftId = left.substring(0, leftSpaceIdx);
                String rightId = right.substring(0, rightSpaceIdx);
                return leftId.compareTo(rightId);
            }

            return res;
        };

        Arrays.sort(logs, compare);
        return logs;
    }
}
