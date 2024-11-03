#include<iostream>
#include<fstream>
using namespace std;

int main() {
    int t;
    cin >> t;
    int ans[300][300];

    for (int i = 0; i < t; ++i) {
        int n, k;
        cin >> n >> k;
        if (k % n == 0) {
            cout << 0 << endl;
        } else {
            cout << 2 << endl;
        }

        for (int j = 0; j < n; ++j) {
            for (int z = 0; z < n; ++z) {
                ans[j][z] = 0;
            }
        }

        int p = 0;
        int q = 0;
        while (k > 0) {
            k--;
            ans[p][q] = 1;
            p++;
            q++;
            q = q%n;
            if (p == n) {
                p = 0;
                q++;
                q = q%n;
            }
        }

        for (int j = 0; j < n; ++j) {
            for (int z = 0; z < n; ++z) {
                cout << ans[j][z];
            }
            cout << endl;
        }
    }
}