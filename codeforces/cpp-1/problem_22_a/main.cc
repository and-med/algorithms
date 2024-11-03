#include<iostream>
#include<fstream>
#include<cstring>
using namespace std;

int main() {
    int a[201];
    std::memset(a, 0, 201 * 4);
    int n;
    cin >> n;
    int c;
    for (int i = 0; i < n; ++i) {
        cin >> c;
        a[c + 100] = a[c + 100] + 1;
    }

    int cnt = 0;
    for (int i = 0; i < 201; ++i) {
        if (a[i] > 0 && cnt == 1) {
            cout << i - 100 << endl;
            return 0;
        } else if (a[i] > 0) {
            cnt++;
        }
    }

    cout << "NO" << endl;
}