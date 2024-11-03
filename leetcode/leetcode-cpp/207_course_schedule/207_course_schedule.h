#ifndef ALGORITHMS_TRAINING_CPP_207_COURSE_SCHEDULE_H
#define ALGORITHMS_TRAINING_CPP_207_COURSE_SCHEDULE_H

#include<vector>
#include<map>
#include<set>
#include<list>
using namespace std;

class Solution {
public:
    bool has_loops_dfs(map<int, list<int>>& map, set<int> visited, int k) {
        if (visited.find(k) != visited.end()) {
            return true;
        }

        visited.insert(k);
        if (map.find(k) != map.end()) {
            const list<int>& elements = map[k];

            for (auto el : elements) {
                if (has_loops_dfs(map, visited, el)) {
                    return true;
                }
            }
        }

        visited.erase(k);
        map.erase(k);
        return false;
    }

    bool canFinish(int numCourses, vector<vector<int>>& prerequisites) {
        map<int, list<int>> courseMap;

        for (auto prerequisite : prerequisites) {
            if (courseMap.find(prerequisite[0]) != courseMap.end()) {
                courseMap[prerequisite[0]].push_back(prerequisite[1]);
            } else {
                list<int> v;
                v.push_back(prerequisite[1]);
                courseMap[prerequisite[0]] = v;
            }
        }

        set<int> visited;

        for (int i = 0; i < numCourses; i++) {
            if (has_loops_dfs(courseMap, visited, i)) {
                return false;
            }
        }

        return true;
    }
};

#endif //ALGORITHMS_TRAINING_CPP_207_COURSE_SCHEDULE_H

