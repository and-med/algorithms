#include<iostream>
#include<fstream>
#include<functional>
using namespace std;

int main() {
    int m, t;
    int a[2][100000];
    int res[100000];

    cin >> t;
    for (int k = 0; k < t; ++k) {
        cin >> m;
        
        for (int i = 0; i < 2; ++i) {
            for (int j = 0; j < m; ++j) {
                cin >> a[i][j];
            }
        }

        for (int i = 1; i < m; ++i) {
            a[1][i] = a[1][i-1] + a[1][i];
            a[0][m - 1 - i] = a[0][m - i] + a[0][m - 1 - i];
        }

        for (int i = 0; i < m; ++i) {
            res[i] = max(i < m - 1 ? a[0][i + 1] : 0, i > 0 ? a[1][i - 1] : 0);
        }

        cout << *min_element(res, res + m) << endl;
    }
}