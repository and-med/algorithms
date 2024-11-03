#include<iostream>
#include<fstream>
#include<string>
#include<vector>
using namespace std;

int main() {
    int t;
    cin >> t;

    for (int k = 0; k < t; k++) {
        int n;
        cin >> n;
        string s;
        cin >> s;
        
        int cnt = 0;
        for (int i = 0; i < n; ++i) {
            cnt += s[i] != '?' ? 1 : 0;
        }
        if (cnt == 0) s[0] = 'R';
        for (int i = 1; i < n; ++i) {
            if (s[i] == '?' && s[i - 1] != '?') {
                s[i] = (s[i - 1] == 'R' ? 'B' : 'R');
            }
        }
        for (int i = n - 2; i >= 0; i--) {
            if (s[i] == '?' && s[i + 1] != '?') {
                s[i] = (s[i + 1] == 'R' ? 'B' : 'R');
            }
        }
        cout << s << endl;
    }
}