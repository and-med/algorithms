#include<iostream>
#include<fstream>
#include<vector>
#include<set>
using namespace std;

int main() {
    int n, m, k;
    cin >> n >> m >> k;
    vector< pair< pair<int,int>, int> > g(m, make_pair(make_pair(0, 0), 0));
    for (int i = 0; i < m; ++i) {
        cin >> g[i].first.first >> g[i].first.second >> g[i].second;
    }

    set<int> bakeries;
    int bakery;
    for (int i = 0; i < k; ++i) {
        cin >> bakery;
        bakeries.insert(bakery);
    }
    
    if (bakeries.empty()) {
        cout << -1 << endl;
        return 0;
    }

    int ru = INT_MAX;
    for (int i = 0; i < m; ++i) {
        if ((bakeries.find(g[i].first.first) != bakeries.end() && bakeries.find(g[i].first.second) == bakeries.end())
        || (bakeries.find(g[i].first.first) == bakeries.end() && bakeries.find(g[i].first.second) != bakeries.end())) {
            ru = std::min(ru, g[i].second);
        }
    }

    if (ru == INT_MAX) {
        cout << -1 << endl;
    } else {
        cout << ru << endl;
    }
}
