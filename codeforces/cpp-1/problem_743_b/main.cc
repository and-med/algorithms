#include<iostream>
#include<fstream>
using namespace std;

int main() {
    int n;
    long long k;
    cin >> n >> k;

    long long size = 1;
    for (int i = 0; i < n; ++i) {
        size *= 2;
    }

    int power = n;
    while (k != size / 2) {
        if (k > size / 2) {
            k -= size / 2;
        }
        size /= 2;
        power -= 1;
    }

    cout << power << endl;
}