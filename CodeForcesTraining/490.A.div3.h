#include<iostream>
#include<vector>
using namespace std;

void doTask(istream& in, ostream& out)
{
	int n, k;
	in >> n >> k;
	vector<int> a(n);

	for (int i = 0; i < n; ++i) in >> a[i];

	int count = 0;

	for (int i = 1; i < n; ++i)
		if (a[i] <= k) count++;
		else break;
	for (int i = n - 1; i >= 0; --i)
		if (a[i] <= k) count++;
		else break;
}