#include<iostream>
#include<algorithm>
#include<set>

void doTask(std::istream& in, std::ostream& out)
{
	int n;
	in >> n;
	std::set<long long> x;

	for (int i = 0; i < n; ++i)
	{
		int a;
		in >> a;
		x.insert(a);
	}

	for (std::set<long long>::iterator it = x.begin(); it != x.end(); ++it)
	{
		long long powerOf2 = 1;
		for (int power = 0; power < 32; ++power, powerOf2*=2)
		{
			if (x.find(*it - powerOf2) != x.end() && x.find(*it + powerOf2) != x.end())
			{
				out << 3 << std::endl << *it - powerOf2 << " " << *it << " " << *it + powerOf2 << std::endl;
				return;
			}
		}
	}

	for (std::set<long long>::iterator it = x.begin(); it != x.end(); ++it)
	{
		long long powerOf2 = 1;
		for (int power = 0; power < 32; ++power, powerOf2 *= 2)
		{
			if (x.find(*it + powerOf2) != x.end())
			{
				out << 2 << std::endl << *it << " " << *it + powerOf2 << std::endl;
				return;
			}
		}
	}

	out << 1 << std::endl << *x.begin() << std::endl;
}