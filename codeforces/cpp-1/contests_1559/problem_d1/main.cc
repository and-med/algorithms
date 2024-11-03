#include<iostream>
#include<fstream>
#include<vector>
using namespace std;

int gett(int* forest, int x) {
    if (x == forest[x]) {
        return x;
    }
    return forest[x] = gett(forest, forest[x]);
}

int main() {
    int n, m1, m2;
    cin >> n >> m1 >> m2;
    int mocha[1000];
    int diana[1000];

    for (int i = 0; i < n; i++) {
        mocha[i] = i;
        diana[i] = i;
    }

    int u, v;
    for (int i = 0; i < m1; i++) {
        cin >> u >> v;
        int tu = gett(mocha, u - 1);
        int tv = gett(mocha, v - 1);
        mocha[tu] = tv;
    }

    for (int i = 0; i < m2; i++) {
        cin >> u >> v;
        int tu = gett(diana, u - 1);
        int tv = gett(diana, v - 1);
        diana[tu] = tv;
    }

    vector< pair<int, int> > ans;
    for (int i = 0; i < n; ++i) {
        for (int j = i + 1; j < n; ++j) {
            if (gett(mocha, i) != gett(mocha, j) && gett(diana, i) != gett(diana, j)) {
                ans.push_back(make_pair(i, j));
                mocha[gett(mocha, i)] = gett(mocha, j);
                diana[gett(diana, i)] = gett(diana, j);
            }
        }
    }

    cout << ans.size() << endl;
    for (int i = 0; i < ans.size(); ++i) {
        cout << ans[i].first + 1 << " " << ans[i].second + 1 << endl;
    }
}