#include<iostream>
#include<fstream>
#include<set>
using namespace std;

struct PairsComparator {
    bool operator()(const std::pair<int, int>& a, const std::pair<int, int>& b) {
        return a.first < b.first;
    } 
};

int main() {
    int t;
    ifstream file("text.txt");
    cin >> t;

    int h[100000];
    int id[100000];
    for (int k = 0; k < t; ++k) {
        int n, m, x;
        cin >> n >> m >> x;

        for (int i = 0; i < n; ++i) {
            cin >> h[i];
        }
        multiset< std::pair<int, int>, PairsComparator > towers;
        for (int i = 0; i < n; ++i) {
            if (towers.size() < m) {
                towers.insert(make_pair(h[i], i + 1));
                id[i] = i + 1;
            } else {
                multiset< std::pair<int, int>, PairsComparator >::iterator it = towers.begin();
                std::pair<int, int> tower = *it;
                towers.erase(it);
                towers.insert(make_pair(tower.first + h[i], tower.second));
                id[i] = tower.second;
            }
        }

        cout << "YES" << endl;
        for (int i = 0; i < n; ++i) {
            cout << id[i] << " ";
        }
        cout << endl;
    }
}