#include<iostream>
#include<fstream>
#include<vector>
using namespace std;

int bfs(const vector <vector<int> >& g, const vector< vector<int> >& v, int* color, int* gcolor, const int& cc, const int& x) {
    if (color[x] != 0) {
        return 0;
    }
    color[x] = cc;
    int ans = 1;
    for (auto group : g[x]) {
        if (gcolor[group] != 0) {
            continue;
        }
        gcolor[group] = cc;
        for (auto vertice : v[group]) {
            ans += bfs(g, v, color, gcolor, cc, vertice);
        }
    }
    return ans;
}

int main() {
    int n, m;
    cin >> n >> m;

    vector< vector<int> > g(n + 1, vector<int>());
    vector< vector<int> > v(m + 1, vector<int>());
    int k;
    int vert;

    for (int i = 0; i < m; ++i) {
        cin >> k;
        for (int j = 0; j < k; ++j) {
            cin >> vert;
            v[i + 1].push_back(vert);
            g[vert].push_back(i + 1);
        }
    }

    int color[5 * 100000 + 1];
    int gcolor[5 * 100000 + 1];
    int size[5 * 100000 + 1];
    for (int i = 0; i <= std::max(n, m); ++i) {
        color[i] = 0;
        gcolor[i] = 0;
        size[i] = 0;
    }
    int currc = 1;
    for (int i = 1; i <= n; ++i) {
        if (color[i] == 0) {
            size[currc] = bfs(g, v, color, gcolor, currc, i);
            currc++;
        }

        cout << size[color[i]] << " ";
    }
    cout << endl;
}