#include<iostream>
#include<vector>
#include<fstream>
#include<functional>
#include<queue>
#include<stack>
using namespace std;

int code(char a) {
    return a - 'a';
}

int main() {
    int g[26][26];
    for (int i = 0; i < 26; ++i) {
        for (int j = 0; j < 26; ++j) {
            g[i][j] = 0;
        }
    }

    ifstream file("test.txt");
    int n;
    file >> n;
    vector<string> words(n);
    for (int i = 0; i < n; ++i) {
        file >> words[i];
    }

    for (int i = 1; i < n; ++i) {
        int len = min(words[i-1].size(), words[i].size());
        int j = 0;
        while (j < len) {
            if (words[i-1][j] != words[i][j]) {
                g[code(words[i-1][j])][code(words[i][j])] = 1;
                break;
            }
            j++;
        }
        if (j == len && words[i-1].size() > words[i].size()) {
            cout << "Impossible" << endl;
            return 0;
        }
    }

    for (int k = 0; k < 26; ++k) {
        for (int i = 0; i < 26; ++i) {
            for (int j = 0; j < 26; ++j) {
                g[i][j] = g[i][j] | (g[i][k] & g[k][j]);
            }
        }
    }

    for (int k = 0; k < 26; ++k) {
        if (g[k][k] == 1) {
            cout << "Impossible" << endl;
            return 0;
        }
    }

    int visited[26];
    for (int i = 0; i < 26; ++i) {
        visited[i] = 0;
    }

    for (int k = 0; k < 26; ++k) {
        if (visited[k] == 1) {
            continue;
        }
        stack<int> vertices;
        vertices.push(k);

        while (!vertices.empty()) {
            int vertex = vertices.top();
            if (visited[vertex] == 1) {
                vertices.pop();
                continue;
            }
            bool anyLess = false;
            for (int i = 0; i < 26; ++i) {
                if (g[i][vertex] == 1 && visited[i] == 0) {
                    anyLess = true;
                    vertices.push(i);
                }
            }
            if (!anyLess) {
                cout << (char)('a' + vertex);
                visited[vertex] = 1;
                vertices.pop();
            }
        }
    }

    cout << endl;
}