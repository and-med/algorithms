#include<iostream>
#include<fstream>
#include<vector>
#include<algorithm>
#include<functional>
using namespace std;

int main() {
    int t;
    cin >> t;

    for (int k = 0; k < t; ++k) {
        int n;
        cin >> n;
        vector<int> a(n);
        for (int i = 0; i < n; ++i) {
            cin >> a[i];
        }
        std::sort(a.begin(), a.end());

        vector<long long> sum(n + 1);
        sum[0] = 0;
        for (int i = 1; i <= n; ++i) {
            sum[i] = a[i - 1] + sum[i - 1];
        }

        double max_val = INT_MIN;
        for (int i = 0; i < n - 1; ++i) {
            double fa = (sum[i + 1] - sum[0]) / (double) (i + 1);
            double fb = (sum[n] - sum[i + 1]) / (double) (n - i - 1);
            max_val = std::max(max_val, fa + fb);
        }

        cout.precision(9);
        cout << std::fixed << max_val << endl;
    }
}