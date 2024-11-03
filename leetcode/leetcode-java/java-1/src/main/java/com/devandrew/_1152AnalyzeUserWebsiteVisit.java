package com.devandrew;

import java.util.*;
import java.util.stream.Collectors;

public class _1152AnalyzeUserWebsiteVisit {
    class WebsiteVisit {
        private String website;
        private int timestamp;

        public WebsiteVisit(String website, int timestamp) {
            this.website = website;
            this.timestamp = timestamp;
        }
    }

    public List<String> mostVisitedPattern(String[] username, int[] timestamp, String[] website) {
        Map<String, List<WebsiteVisit>> visits = new HashMap<>();

        for (int i = 0; i < username.length; i++) {
            WebsiteVisit visit = new WebsiteVisit(website[i], timestamp[i]);

            if (visits.containsKey(username[i])) {
                visits.get(username[i]).add(visit);
            } else {
                List<WebsiteVisit> userVisits = new ArrayList<>();
                userVisits.add(visit);
                visits.put(username[i], userVisits);
            }
        }

        TreeMap<String, Integer> patternCounts = new TreeMap<>();

        for (Map.Entry<String, List<WebsiteVisit>> entry : visits.entrySet()) {
            Set<String> userPatterns = new HashSet<>();
            List<WebsiteVisit> userVisits = entry.getValue();
            userVisits.sort(Comparator.comparingInt(a -> a.timestamp));

            for (int i = 0; i < userVisits.size(); i++) {
                for (int j = i + 1; j < userVisits.size(); j++) {
                    for (int k = j + 1; k < userVisits.size(); k++) {
                        userPatterns.add(String.format("%s_%s_%s",
                                userVisits.get(i).website,
                                userVisits.get(j).website,
                                userVisits.get(k).website));
                    }
                }
            }

            for (String userPattern : userPatterns) {
                patternCounts.merge(userPattern, 1, Integer::sum);
            }
        }

        Map.Entry<String, Integer> firstEntry = patternCounts.firstEntry();
        String maxPattern = firstEntry.getKey();
        int maxCnt = firstEntry.getValue();

        for (Map.Entry<String, Integer> entry : patternCounts.entrySet()) {
            if (entry.getValue() > maxCnt) {
                maxPattern = entry.getKey();
                maxCnt = entry.getValue();
            }
        }

        return Arrays.stream(maxPattern.split("_"))
                .collect(Collectors.toList());
    }
}
