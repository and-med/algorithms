#include<iostream>
#include<algorithm>
#include<vector>

struct umbrella
{
	int position;
	int weight;

	umbrella() : position(0), weight(0) {}
	umbrella(int pos, int w) : position(pos), weight(w) {}
};

void doTask(std::istream& in, std::ostream& out)
{
	int a, n, m;
	in >> a >> n >> m;
	std::vector<bool> isRain(a + 1, false);
	std::vector<umbrella> umbrellas(m);
	std::vector<umbrella> minUmbrellas(a + 1, umbrella(-1, INT_MAX));

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
		in >> umbrellas[i].position >> umbrellas[i].weight;
	}

	for (int i = 0; i < m; ++i)
	{
		minUmbrellas[umbrellas[i].position].weight = std::min(minUmbrellas[umbrellas[i].position].weight, umbrellas[i].weight);
		minUmbrellas[umbrellas[i].position].position = umbrellas[i].position;
	}

	std::vector<std::vector<int>> minFatigues(a + 1, std::vector<int>(m + 1, INT_MAX));
	minFatigues[0][m] = 0;

	for (int i = 0; i < a; ++i)
	{
		for (int j = 0; j < m + 1; ++j)
		{
			if (minFatigues[i][j] == INT_MAX) continue;
			if (!isRain[i]) {
				minFatigues[i + 1][j] = std::min(minFatigues[i + 1][m], minFatigues[i][j]);
			}
			if (j < m)
			{
				minFatigues[i + 1][j] = std::min(minFatigues[i + 1][j], minFatigues[i][j] + umbrellas[j].weight);
			}
			if (minUmbrellas[i].position != -1)
			{
				minFatigues[i + 1][minUmbrellas[i].position] = std::min(minFatigues[i + 1][minUmbrellas[i].position], minFatigues[i][j] + minUmbrellas[i].weight);
			}
		}
	}

	int min = INT_MAX;
	for (int i = 0; i < m + 1; ++i)
	{
		min = std::min(min, minFatigues[a][i]);
	}

	if (min == INT_MAX)
	{
		out << -1 << std::endl;
	}
	else
	{
		out << min << std::endl;
	}
}