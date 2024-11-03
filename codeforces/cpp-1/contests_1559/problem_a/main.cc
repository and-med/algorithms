#include<iostream>
#include<fstream>
#include<algorithm>
#include<functional>
using namespace std;

int main() {
    int t;
    cin >> t;
    int a[100];

    for (int k = 0; k < t; ++k) {
        int n;
        cin >> n;
        for (int i = 0; i < n; ++i) {
            cin >> a[i];
        }
        int res = a[0];
        for (int i = 1; i < n; ++i) {
            res &= a[i];
        }
        cout << res << endl;
    }
}