#include<iostream>
#include<vector>
using namespace std;

long long limit = 2e18;

void doTask(istream& in, ostream& out)
{
	int n, k;
	in >> n >> k;
	vector<int> a(n);

	for (int i = 0; i < n; ++i) in >> a[i];

	vector<int> next(n + 1);
	vector<long long> sum(n + 1);

	next[n] = n;
	for (int i = n - 1; i >= 0; --i) 
	{
		next[i] = next[i + 1];
		if (a[i] > 1) next[i] = i;
	}
	sum[0] = 0;
	for (int i = 1; i <= n; ++i) sum[i] = sum[i - 1] + a[i - 1];

	int ans = 0;
	for (int i = 0; i < n; ++i)
	{
		long long p = 1;
		for (int j = next[i]; j < n && a[j] <= limit / p; j = next[j + 1])
		{
			int s = sum[j + 1] - sum[i];
			if (a[j] <= limit / p) p = p * a[j];
			if (p % k == 0) {
				int nxt = next[j + 1];
				int countOnes = nxt - j - 1;
				int requiredSum = p / k;
				if (s <= requiredSum && s + countOnes >= requiredSum) {
					ans++;
				}
			}
		}
	}

	if (k == 1)
	{
		int countOnes = 0;
		for (int i = 0; i < n; ++i)
		{
			if (next[i] != i)
			{
				countOnes++;
			}
		}
		ans += countOnes;
	}
	out << ans << endl;
}