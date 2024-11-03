#include<iostream>
#include<fstream>
using namespace std;

int main() {
    int n;
    int h[100000];
    cin >> n;
    for (int i = 0; i < n; ++i) {
        cin >> h[i];
    }
    int hmax[100000];
    hmax[n - 1] = 0;
    for (int i = n - 2; i >= 0; --i) {
        hmax[i] = std::max(hmax[i + 1], h[i + 1]);
    }
    for (int i = 0; i < n; ++i) {
        if (h[i] > hmax[i]) {
            cout << 0 << " ";
        } else if (h[i] == hmax[i]) {
            cout << 1 << " ";
        } else {
            cout << hmax[i] + 1 - h[i] << " ";
        }
    }
    cout << endl;
}