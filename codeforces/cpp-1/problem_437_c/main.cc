#include<iostream>
#include<fstream>
#include<vector>
#include<set>
#include<functional>
#include<algorithm>
using namespace std;

int main() {
    int n, m;
    cin >> n >> m;
    int v[1000];
    for (int i = 0; i < n; ++i) {
        cin >> v[i];
    }

    int res = 0;
    int x, y;
    for (int i = 0; i < m; ++i) {
        cin >> x >> y;
        res += std::min(v[x - 1], v[y - 1]);
    }

    cout << res << endl;
}