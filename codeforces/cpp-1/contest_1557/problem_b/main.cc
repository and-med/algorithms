#include<iostream>
#include<fstream>
#include<vector>
#include<algorithm>
using namespace std;

int main() {
    int t;
    ifstream file("file.txt");
    file >> t;

    for (int j = 0; j < t; j++) {
        int n, k;
        file >> n >> k;
        vector< pair<int, int> > v(n);
		for (int i = 0; i < n; i++) {
			file >> v[i].first;
			v[i].second = i;
		}
		sort(v.begin(), v.end());
		int ans = 1;
		for (int i = 1; i < n; i++)
			if (v[i - 1].second + 1 != v[i].second)
				ans++;

        if (ans <= k) {
            cout << "YES" << endl;
        } else {
            cout << "NO" << endl;
        }
    }
}