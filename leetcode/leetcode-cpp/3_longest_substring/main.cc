#include<iostream>
#include<vector>
#include<string>
#include<fstream>
#include<map>
#include<unordered_map>
using namespace std;

class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        unordered_map<int, int> latestPos;
        int max_pos = 0;
        int start = 0;
        for (int i = 0; i < s.length(); i++) {
            char curr = s[i];

            auto it = latestPos.find(curr);
            
            if (it != latestPos.end() && it->second >= start) {
                start = latestPos[curr] + 1;
            } else {
                max_pos = max(max_pos, i - start + 1);
            }

            latestPos[curr] = i;
        }
        return max_pos;
    }
};

int main() {
    ifstream in("input.txt");
    string s;
    in >> s;
    Solution sol;
    int res = sol.lengthOfLongestSubstring(s);
    ofstream out("output.txt");
    out << res << endl;
}
