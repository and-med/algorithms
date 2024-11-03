#include<iostream>
#include<algorithm>
#include<vector>
using namespace std;

struct umbrella
{
	int position;
	int weight;

	umbrella() : position(0), weight(0) {}
	umbrella(int pos, int w) : position(pos), weight(w) {}
};

void doTask(istream& in, ostream& out)
{
	int a, n, m;
	in >> a >> n >> m;

	vector<umbrella> umbrellas(m);
	vector<pair<int, int>> minUmbrellas(a + 1, make_pair(INT_MAX, -1));
	vector<bool> isRain(a);

	for (int i = 0; i < n; ++i)
	{
		int l, r;
		in >> l >> r;
		for (int j = l; j < r; ++j)
		{
			isRain[j] = true;
		}
	}

	for (int i = 0; i < m; ++i)
	{
		int x, p;
		in >> x >> p;
		umbrellas[i].position = x;
		umbrellas[i].weight = p;

		minUmbrellas[umbrellas[i].position].second = umbrellas[i].weight < minUmbrellas[umbrellas[i].position].first ? i : minUmbrellas[umbrellas[i].position].second;
		minUmbrellas[umbrellas[i].position].first = min(minUmbrellas[umbrellas[i].position].first, umbrellas[i].weight);
	}

	vector<vector<int>> d(a + 1, vector<int>(m + 1, INT_MAX));
	d[0][m] = 0;

	for (int i = 0; i < a; ++i)
	{
		for (int j = 0; j <= m; ++j)
		{
			if (d[i][j] == INT_MAX) continue;
			if (!isRain[i]) d[i + 1][m] = min(d[i + 1][m], d[i][j]);
			if (j < m) d[i + 1][j] = min(d[i + 1][j], d[i][j] + umbrellas[j].weight);
			if(minUmbrellas[i].second != -1) d[i + 1][minUmbrellas[i].second] = min(d[i + 1][minUmbrellas[i].second], d[i][j] + minUmbrellas[i].first);
		}
	}

	int minFatigue = INT_MAX;
	for (int i = 0; i <= m; ++i)
	{
		minFatigue = min(minFatigue, d[a][i]);
	}
	if (minFatigue == INT_MAX) out << -1 << endl;
	else out << minFatigue << endl;
}