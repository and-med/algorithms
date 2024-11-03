#include<iostream>
#include<fstream>
#include<vector>
using namespace std;

int main() {
    int t;
    cin >> t;
    for (int k = 0; k < t; k++) {
        int n;
        cin >> n;
        vector<long long> a(n);
        for (int i = 0; i < n; ++i) {
            cin >> a[i];
        }

        long long res = a[0] * a[1];
        for (int i = 1; i < n - 1; ++i) {
            res = std::max(res, a[i] * a[i + 1]);
        }
        cout << res << endl;
    }
}