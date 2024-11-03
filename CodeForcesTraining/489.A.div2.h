#include<iostream>
#include<vector>
#include<algorithm>
#include<set>
using namespace std;

void doTask(istream& in, ostream& out)
{
	int n;
	in >> n;
	vector<int> a(n);
	for (int i = 0; i < n; ++i) in >> a[i];
	sort(a.begin(), a.end());
	set<int> unique;
	for (int i = 0; i < n; ++i)
	{
		if (a[i] != 0) unique.insert(a[i]);
	}
	out << unique.size() << endl;
}