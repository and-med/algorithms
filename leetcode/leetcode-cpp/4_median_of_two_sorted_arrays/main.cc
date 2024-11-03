#include<iostream>
#include<vector>
#include<string>
#include<fstream>
#include<map>
#include<unordered_map>
using namespace std;

const int MIN_INFINITY = -10000000;
const int MAX_INFINITY = 10000000;

class Solution {
public:
    double findMedianSortedArrays(vector<int>& list1, vector<int>& list2) {
        vector<int>* p1 = &list1;
        vector<int>* p2 = &list2;
        if (p2->size() < p1->size()) {
            vector<int>* temp = p2;
            p2 = p1;
            p1 = temp;
        }
        vector<int>& a = *p1;
        vector<int>& b = *p2;

        int total = a.size() + b.size();
        int half = total / 2;
        int l = 0;
        int r = a.size();

        if (a.size() == 0 || b.size() == 0) {
            if (total % 2 == 0) {
                return (double)(b[half - 1] + b[half]) / (double)2;
            }
            return b[half];
        }

        if (a[0] >= b[b.size() - 1] && a.size() != b.size()) {
            if (total % 2 == 0) {
                return (double)(b[half - 1] + b[half]) / (double)2;
            }
            return b[half];
        }

        if (a[a.size() - 1] <= b[0] && a.size() != b.size()) {
            if (total % 2 == 0) {
                return (double)(b[half - a.size() - 1] + b[half - a.size()]) / (double) 2;
            }
            return b[half - a.size()];
        }

        int m = (l + r) / 2;
        int t = half - m - 2;

        while (true) {
            int a_m = m >= 0 ? a[m] : MIN_INFINITY;
            int a_m_1 = m + 1 < a.size() ? a[m + 1] : MAX_INFINITY;
            int b_t = t >= 0 ? b[t] : MIN_INFINITY;
            int b_t_1 = t + 1 < b.size() ? b[t + 1] : MAX_INFINITY;

            if (a_m <= b_t_1 && b_t <= a_m_1 || l == r) {
                if (total % 2 == 0) {
                    return (double)(max(a_m, b_t) + min(a_m_1, b_t_1)) / (double)2;
                }
                return min(a_m_1, b_t_1);
            }

            if (a_m > b_t_1) {
                r = m;
                m = (r == 0 ? -1 : ((l + r) / 2));
                t = half - m - 2;
            } else if (b_t > a_m_1) {
                l = m;
                m = (l + r) / 2;
                t = half - m - 2;
            }
        }

        return 0;
    }
};

vector<string> string_split(string s, char separator) {
    vector<string> res;
    string::iterator prev = s.begin();
    for (string::iterator it = s.begin(); it != s.end(); it++) {
        if (*it == separator) {
            string splitted(prev, it);
            if (splitted.length() > 0) {
                res.push_back(splitted);
            }
            prev = it + 1;
        }
    }

    string last(prev, s.end());
    if (last.length() > 0) {
        res.push_back(last);
    }
    return res;
}

vector<int> parse_array(string s) {
    string noBraces = s.substr(1, s.length() - 2);
    vector<int> res;
    vector<string> strings = string_split(noBraces, ',');
    transform(strings.begin(), strings.end(), std::back_inserter(res), [](string d) { return atoi(d.c_str()); });
    return res;
}

int main() {
    ifstream in("input.txt");
    string x, y;
    in >> x >> y;
    vector<int> arr1 = parse_array(x), arr2 = parse_array(y);
    Solution sol;
    double res = sol.findMedianSortedArrays(arr1, arr2);
    ofstream out("output.txt");
    out << res << endl;
}
