#include<iostream>
#include<fstream>
#include<algorithm>
using namespace std;

int main() {
    int n;
    ifstream file("file.txt");
    file >> n;
    int a[100000];
    for (int i = 0; i < n; ++i) {
        file >> a[i];
    }
    std::sort(a, a + n);
    
    int ans = 0;
    for (int i = 0, j = n /2; i < n / 2 && j < n - 1; i++, j++) {
        if (a[i] < a[j] && a[i] < a[j + 1]) {
            ans++;
        }
    }
    cout << ans << endl;
    for (int k = 0, i = 0, j = n / 2; k < n; ++k) {
        if (j - i == n / 2) {
            cout << a[j] << " ";
            j++;
        } else {
            cout << a[i] << " ";
            i++;
        }
    }
    cout << endl;
}