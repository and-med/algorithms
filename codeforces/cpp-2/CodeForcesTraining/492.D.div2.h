#include<iostream>
#include<vector>
using namespace std;

void doTask(istream& in, ostream& out)
{
	int n;
	in >> n;
	vector<int> a(2*n);

	for (int i = 0; i < 2*n; ++i) in >> a[i];

	int count = 0;
	for (int i = 0; i < 2*n; i += 2)
	{
		int ai = a[i];
		int j = i + 1;
		for (; j < 2*n; j++) if (a[j] == a[i]) break;
		if (j - i == 1) continue;
		for (; j > i + 1; j--)
		{
			swap(a[j], a[j - 1]);
			count++;
		}
	}

	out << count << endl;
}