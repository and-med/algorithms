#include<iostream>
#include<vector>
#include<algorithm>
#include<set>
using namespace std;

long long gcd(long long u, long long v) {
	while (v != 0) {
		long long r = u % v;
		u = v;
		v = r;
	}
	return u;
}

long long lcm(long long a, long long b, long long gcd)
{
	long long temp = gcd;

	return temp ? (a / temp * b) : 0;
}

void doTask(istream& in, ostream& out)
{
	long long l, r, x, y;
	in >> l >> r >> x >> y;
	vector<long long> divisors;
	long long square_root = (long long)sqrt(y) + 1;
	for (long long i = 1; i <= r && i < square_root; ++i)
	{
		if (y % i == 0)
		{
			if (y / i == i)
			{
				if (i >= l && i <= r)
					divisors.push_back(i);
			}
			else
			{
				if (i >= l && i <= r)
					divisors.push_back(i);
				if (y / i >= l && y / i <= r)
					divisors.push_back(y / i);
			}
		}
	}

	vector<long long> newDivisors;
	for (long long i = 0; i < divisors.size(); ++i)
	{
		if (divisors[i] % x == 0) newDivisors.push_back(divisors[i]);
	}

	long long count = 0;
	for (long long i = 0; i < newDivisors.size(); ++i)
	{
		for (long long j = 0; j < newDivisors.size(); ++j)
		{
			long long currGcd = gcd(newDivisors[i], newDivisors[j]);
			if (currGcd == x && lcm(newDivisors[i], newDivisors[j], currGcd) == y)
			{
				count += 1;
			}
		}
	}

	out << count << endl;
}