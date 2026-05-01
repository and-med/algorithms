#include<iostream>
#include<vector>
#include<map>
using namespace std;

int cost_at(int cell) {
    if (cell == 0) {
        return 0;
    } else if (cell == 1) {
        return 1;
    } else if (cell == 2) {
        return 1;
    }

    throw "impossible";
}

struct Path {
    int score;
    int cost;
    bool initialized;

    Path(): score(0), cost(0), initialized(false) {}
    Path(int score, int cost): score(score), cost(cost), initialized(true) {}
    Path(const Path& prev, int score) {
        this->score = prev.score + score;
        this->cost = prev.cost + cost_at(score);
        this->initialized = true;
    }
};

int INF = 1<<30;

class Solution {
private:
    vector<vector<vector<Path>>> _cache;

public:
    Path bestPathLowerThan(vector<vector<int>>& grid, int i, int j, int k) {
        if (k < 0) {
            return Path(-1, INF);
        }
        if (i < 0 && j == 0) {
            return Path(0, 0);
        }
        if (j < 0 && i == 0) {
            return Path(0, 0);
        }
        if (i < 0) {
            return Path(-1, INF);
        }
        if (j < 0) {
            return Path(-1, INF);
        }

        Path cache_entry = this->_cache[i][j][k];
        if (cache_entry.initialized) {
            return cache_entry;
        }

        int currentGridCost = cost_at(grid[i][j]);
        int currentGridScore = grid[i][j];
        Path left = bestPathLowerThan(grid, i, j - 1, k - currentGridCost);
        Path top = bestPathLowerThan(grid, i - 1, j, k - currentGridCost);

        // cout << "left best at (" << i << "," << j << ")" << " with threshold of " << k << " is cost=" << left.cost << ", score=" << left.score << endl;
        // cout << "top best at (" << i << "," << j << ")" << " with threshold of " << k << " is cost=" << top.cost << ", score=" << top.score << endl;

        Path result;

        if (top.cost + currentGridCost > k && left.cost + currentGridCost > k) {
            result = Path(-1, INF);
        } else if (top.cost + currentGridCost > k) {
            // top too expensive - go left
            result = Path(left, currentGridScore);
        } else if (left.cost + currentGridCost > k) {
            // left too expensive - go top
            result = Path(top, currentGridScore);
        } else {
            // both are not too expensive - choose best score?
            if (top.score + currentGridScore > left.score + currentGridScore) {
                // top score better
                result = Path(top, currentGridScore);
            } else {
                // left score better
                result = Path(left, currentGridScore);
            }
        }

        this->_cache[i][j][k] = result;
        return result;
    }

    int maxPathScore(vector<vector<int>>& grid, int k) {        
        int n = grid.size();
        int m = grid[0].size();
        this->_cache = vector<vector<vector<Path>>>(n, vector<vector<Path>>(m, vector<Path>(1001)));
        return bestPathLowerThan(grid, n - 1, m - 1, k).score;
    }
};

int main() {

}