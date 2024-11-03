#include<iostream>
#include<fstream>
#include<vector>
#include<set>
using namespace std;

int main() {
    int n, m;
    cin >> n >> m;
    
    vector<int> fr(n, 0);
    vector<int> weak_fr(n, 0);
    int survivals = n;
    for (int i = 0; i < m; ++i) {
        int u, v;
        cin >> u >> v;
        int surv_u = weak_fr[u - 1] == fr[u - 1];
        int surv_v = weak_fr[v - 1] == fr[v - 1];
        fr[u - 1]++;
        fr[v - 1]++;
        weak_fr[u - 1] += u > v;
        weak_fr[v - 1] += v > u;
        survivals -= surv_u == true && weak_fr[u - 1] != fr[u - 1];
        survivals -= surv_v == true && weak_fr[v - 1] != fr[v - 1];
    }

    int q;
    cin >> q;
    for (int k = 0; k < q; ++k) {
        int op;
        cin >> op;
        int u, v;
        if (op == 1 || op == 2) {
            cin >> u >> v;
        }

        if (op == 1) {
            int surv_u = weak_fr[u - 1] == fr[u - 1];
            int surv_v = weak_fr[v - 1] == fr[v - 1];
            fr[u - 1]++;
            fr[v - 1]++;
            weak_fr[u - 1] += u > v;
            weak_fr[v - 1] += v > u;
            survivals -= surv_u == true && weak_fr[u - 1] != fr[u - 1];
            survivals -= surv_v == true && weak_fr[v - 1] != fr[v - 1];
        } else if (op == 2) {
            int surv_u = weak_fr[u - 1] == fr[u - 1];
            int surv_v = weak_fr[v - 1] == fr[v - 1];
            fr[u - 1]--;
            fr[v - 1]--;
            weak_fr[u - 1] -= u > v;
            weak_fr[v - 1] -= v > u;
            survivals += surv_u == false && weak_fr[u - 1] == fr[u - 1];
            survivals += surv_v == false && weak_fr[v - 1] == fr[v - 1];
        } else {
            cout << survivals << endl;
        }
    }
}