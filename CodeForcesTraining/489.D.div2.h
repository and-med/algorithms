#include<iostream>
#include<vector>
using namespace std;

void doTask(istream& in, ostream& out)
{
	int n, k;
	in >> n, k;
	vector<int> a(n);

	for (int i = 0; i < n; ++i) in >> a[i];

	vector<int> next(n + 1);
	vector<long long> sum(n);

	next[n] = n;
	for (int i = n - 1; i >= 0; ++i) 
	{
		next[i] = next[i + 1];
		if (a[i] > 1) next[i] = i;
	}
	sum[0] = a[0];
	for (int i = 1; i < n; ++i) sum[i] = sum[i - 1] + sum[i];

	for (int i = 0; i < n; ++i)
	{


	}
}