#include<iostream>
#include<algorithm>

void doTask(std::istream& in, std::ostream& out)
{
	int n;
	in >> n;
	int *s = new int[n];
	int *c = new int[n];
	int *leastSum2 = new int[n];

	for (int i = 0; i < n; ++i) in >> s[i];
	for (int i = 0; i < n; ++i) in >> c[i];

	leastSum2[0] = INT_MAX;
	for (int i = 1; i < n; ++i)
	{
		int min = INT_MAX;
		for (int j = 0; j < i; j++)
		{
			if (s[j] < s[i]) min = std::min(min, c[i] + c[j]);
		}
		leastSum2[i] = min;
	}

	int *leastSum3 = new int[n];
	leastSum3[0] = INT_MAX;
	leastSum3[1] = INT_MAX;

	for (int i = 2; i < n; ++i)
	{
		int min = INT_MAX;
		for (int j = 1; j < i; j++)
		{
			if (s[j] < s[i] && leastSum2[j] != INT_MAX) min = std::min(min, c[i] + leastSum2[j]);
		}
		leastSum3[i] = min;
	}

	int min = INT_MAX;
	for (int i = 0; i < n; ++i) min = std::min(min, leastSum3[i]);

	if (min == INT_MAX) out << -1 << std::endl;
	else out << min << std::endl;
}