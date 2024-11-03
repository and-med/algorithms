#include<iostream>
#include<algorithm>

void doTask(std::istream& in, std::ostream& out)
{
	int n;
	in >> n;
	int *s = new int[n];
	int **dp = new int*[3];

	for (int i = 0; i < 3; ++i)
	{
		dp[i] = new int[n];
	}
	for (int i = 0; i < n; ++i) in >> s[i];
	for (int i = 0; i < n; ++i) 
	{ 
		in >> dp[0][i];
		dp[1][i] = INT_MAX;
		dp[2][i] = INT_MAX;
	}

	for (int i = 0; i < n; ++i)
	{
		for (int j = 1; j <= 2; ++j)
		{
			for (int k = 0; k < i; ++k)
			{
				if (s[k] < s[i] && dp[j-1][k] != INT_MAX) dp[j][i] = std::min(dp[j][i], dp[j - 1][k] + dp[0][i]);
			}
		}
	}

	int min = INT_MAX;
	for (int i = 0; i < n; ++i) min = std::min(min, dp[2][i]);

	if (min == INT_MAX) out << -1 << std::endl;
	else out << min << std::endl;
}