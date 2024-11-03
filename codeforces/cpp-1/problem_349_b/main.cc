#include<iostream>
#include<fstream>
#include<vector>
using namespace std;

int main() {
    int v;
    int a[9];

    cin >> v;
    for (int i = 0; i < 9; ++i) {
        cin >> a[i];
    }

    int minDigit = 0;
    for (int i = 1; i < 9; ++i) {
        if (a[i] <= a[minDigit]) {
            minDigit = i;
        }
    }

    int len = v / a[minDigit];
    int weight = a[minDigit] * len;

    if (len == 0) {
        cout << -1 << endl;
        return 0;
    }

    vector<int> number = vector<int>(len, minDigit + 1);
    for (int i = 0; i < number.size() && weight < v; ++i) {
        for (int j = 8; j > minDigit; --j) {
            if (weight - a[minDigit] + a[j] <= v) {
                number[i] = j + 1;
                weight = weight - a[minDigit] + a[j];
                break;
            }
        }
    }

    for (int i = 0; i < number.size(); ++i) {
        cout << number[i];
    }
    cout << endl;
}