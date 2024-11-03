#include<iostream>
#include<fstream>
using namespace std;

int main() {
    int n;
    cin >> n;
    int a;
    int prev;
    long long res = 0;
    int cons = 1;
    for (int i = 0; i < n; ++i) {
        cin >> a;
        if (a == prev) {
            cons++;
        } else {
            cons = 1;
        }
        res += cons;
        prev = a;
    }

    cout << res << endl;
}