#include<iostream>
#include<fstream>
using namespace std;

int main() {
    int t;
    cin >> t;
    int a[10000];
    for (int k = 0; k < t; k++) {
        int n;
        cin >> n;
        for (int i = 0; i < n; ++i) {
            cin >> a[i];
        }

        if (a[0] == 1) {
            cout << n + 1 << " ";
            for (int i = 0; i < n; ++i) {
                cout << i + 1 << " ";
            }
            cout << endl;
        } else if (a[n - 1] == 0) {
            for (int i = 0; i < n; ++i) {
                cout << i + 1 << " ";
            }
            cout << n + 1 << " ";
            cout << endl;
        } else {
            int id = 0;
            for (int i = 1; i < n; ++i) {
                if (a[i - 1] == 0 && a[i] == 1) {
                    id = i - 1;
                    break;
                }
            }

            for (int i = 0; i < n; ++i) {
                cout << i + 1 << " ";
                if (i == id) {
                    cout << n + 1 << " ";
                }
            }
            cout << endl;
        }
    }
}