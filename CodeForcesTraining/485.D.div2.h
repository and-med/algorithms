#include<iostream>
#include<algorithm>
#include<vector>
#include<queue>
#include<numeric>
using namespace std;

void superSmartBfs(vector<vector<int>>& g, vector<int>& vertices, vector<int>& d, int n) 
{
	queue<int> openVertices;
	vector<bool> isClosed(n + 1);
	for (int val : vertices)
	{
		openVertices.push(val);
		isClosed[val] = true;
		d[val] = 0;
	}

	while (!openVertices.empty())
	{
		int r = openVertices.front();
		openVertices.pop();

		for (int v : g[r])
		{
			if (!isClosed[v]) 
			{
				isClosed[v] = true;
				openVertices.push(v);
				d[v] = d[r] + 1;
			}
		}
	}
}

void doTask(istream& in, ostream& out)
{
	int n, m, k, s;
	in >> n >> m >> k >> s;
	vector<vector<int>> g(n + 1, vector<int>());
	vector<vector<int>> a(k + 1, vector<int>());
	vector<vector<int>> d(k + 1, vector<int>(n + 1));

	for (int i = 0; i < n; ++i)
	{
		int ai;
		in >> ai;
		a[ai].push_back(i + 1);
	}

	for (int i = 0; i < m; ++i)
	{
		int u, v;
		in >> u >> v;
		g[u].push_back(v);
		g[v].push_back(u);
	}

	for (int i = 1; i <= k; ++i)
	{
		superSmartBfs(g, a[i], d[i], n);
	}

	for (int i = 1; i <= n; ++i)
	{
		vector<int> dist;
		for (int j = 1; j <= k; ++j)
		{
			dist.push_back(d[j][i]);
		}
		sort(dist.begin(), dist.end());
		out << accumulate(dist.begin(), dist.begin() + s, 0) << " ";
	}
	out << endl;
}