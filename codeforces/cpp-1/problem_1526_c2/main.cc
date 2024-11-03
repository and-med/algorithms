#include<iostream>
#include<fstream>
#include<set>
using namespace std;

int main() {
    int n;
    cin >> n;
    int a[200000];
    for (int i = 0; i < n; ++i) {
        cin >> a[i];
    }

    multiset<int> neg_p;
    long long health = 0;
    int res = 0;
    for (int i = 0; i < n; ++i) {
        if (a[i] >= 0) {
            health += a[i];
            res += 1;
        } else {
            if (health + a[i] >= 0) {
                health += a[i];
                neg_p.insert(a[i]);
                res += 1;
            } else if (!neg_p.empty()) {
                multiset<int>::iterator minIt = neg_p.begin();
                int minVal = *minIt;
                if (minVal < a[i]) {
                    health -= minVal;
                    health += a[i];

                    neg_p.erase(minIt);
                    neg_p.insert(a[i]);
                }
            }
        }
    }

    cout << res << endl;
}