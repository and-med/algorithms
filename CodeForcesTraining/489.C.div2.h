#include<iostream>
#include<vector>
#include<algorithm>
#include<set>
using namespace std;

long long binPowMod(long long n, long long pow, long long mod)
{
	if (pow == 0) return 1;
	if (pow == 1) return n % mod;
	long long calc = binPowMod(n, pow / 2, mod);
	long long res = (calc * calc) % mod;
	return pow % 2 ? (res * n) % mod : res;
}

void doTask(istream& in, ostream& out)
{
	long long mod = 1e9 + 7;
	long long x, k;
	in >> x >> k;
	if (x == 0) out << 0 << endl;
	else {
		long long res = (binPowMod(2, k + 1, mod) * (x % mod)) % mod - binPowMod(2, k, mod) + 1;
		res = (res + mod) % mod;
		out << res << endl;
	}
}