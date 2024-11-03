#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;

void doTask(istream& in, ostream& out)
{
	int n;
	in >> n;
	vector<int> a(n);

	for (int i = 0; i < n; ++i) in >> a[i];

	int currMin = INT_MAX;
	int res = -1;
	for (int i = 0; i < n; ++i)
	{
		int rest = a[i] % n;
		int curr = (i < rest ? n - abs(i - rest) : abs(i - rest)) + a[i];
		res = curr < currMin ? i : res;
		currMin = min(currMin, curr);
	}

	out << res + 1 << endl;
}