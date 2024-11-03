#include <iostream>
#include <vector>
#include <string>
#include <fstream>
#include "207_course_schedule.h"

using namespace std;

vector<string> string_split(const string& s, const string& separator) {
    vector<string> res;

    unsigned prev = 0;
    unsigned next = s.find(separator);

    while (next != -1) {
        string piece(s.begin() + prev, s.begin() + next);
        res.push_back(piece);
        prev = next + separator.length();
        next = s.find(separator, prev);
    }

    string last(s.begin() + prev, s.end());
    if (last.length() > 0) {
        res.push_back(last);
    }
    return res;
}

string erase_curly_braces(const string& s) {
    unsigned start = s[0] == '[' ? 1 : 0;
    unsigned length = s[s.length() - 1] == ']' ? s.length() - start - 1 : s.length() - start;
    return s.substr(start, length);
}

vector<vector<int>> read_arr(const string& s) {
    vector<string> nested = string_split(erase_curly_braces(s), "],[");
    vector<vector<int>> res;
    for (auto i : nested) {
        vector<int> temp;
        for (auto j : string_split(erase_curly_braces(i), ",")) {
            temp.push_back(atoi(j.c_str()));
        }
        res.push_back(temp);
    }
    return res;
}

int main() {
    string task_name = "207";
    ifstream in("./inputs/" + task_name + ".txt");
    int numCourses;
    in >> numCourses;

    string arr;
    in >> arr;
    vector<vector<int>> prerequisites = read_arr(arr);

    Solution sol;
    bool canFinish = sol.canFinish(numCourses, prerequisites);
    ofstream out("outputs/" + task_name + ".txt");
    out << canFinish << endl;
    return 0;
}
