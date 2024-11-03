#include<iostream>
#include<fstream>
#include<vector>
#include<algorithm>
using namespace std;

int main() {
    int t;
    cin >> t;
    for (int k = 0; k < t; ++k) {
        int n;
        cin >> n;
        vector<int> a(n, 0);
        vector< vector<int> > cnt(2, vector<int>(100000 + 1, 0));
        for (int i = 0; i < n; ++i) {
            cin >> a[i];
            cnt[i % 2][a[i]]++;
        }

        std::sort(a.begin(), a.end());

        bool possible = true;
        for (int i = 0; i < n; ++i) {
            if (--cnt[i % 2][a[i]] < 0) {
                possible = false;
                break;
            }
        }

        cout << (possible ? "YES" : "NO") << endl;
    }
}