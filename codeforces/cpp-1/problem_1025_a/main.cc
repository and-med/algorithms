#include<iostream>
#include<fstream>
using namespace std;

int main() {
    int n;
    cin >> n;
    
    int existing[26];
    for (int i = 0; i < 26; ++i) {
        existing[i] = 0;
    }

    if (n == 1) {
        cout << "Yes" << endl;
        return 0;
    }
    char puppy;
    for (int i = 0; i < n; ++i) {
        cin >> puppy;
        if (existing[puppy-'a'] == 1) {
            cout << "Yes" << endl;
            return 0;
        }
        existing[puppy-'a'] = 1;
    }
    cout << "No" << endl;
}