#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;

void buildTree(vector<pair<long long, int>>& st, vector<int>& a, int v, int l, int r)
{
	if (l == r)
	{
		st[v].first = a[l];
		st[v].second = a[l];
		return;
	}

	int m = (l + r) / 2;
	buildTree(st, a, v * 2, l, m);
	buildTree(st, a, v * 2 + 1, m + 1, r);
	st[v].first = st[v * 2].first + st[v * 2 + 1].first;
	st[v].second = max(st[v * 2].second, st[v * 2 + 1].second);
}

int sum(vector<pair<long long, int>>& st, int v, int tl, int tr, int l, int r)
{
	if (l > r) return 0;
	int m = (tl + tr) / 2;
	if (l == tl && r == tr) return st[v].first;
	return sum(st, v * 2, tl, m, l, min(m, r)) + sum(st, v * 2 + 1, m + 1, tr, max(l, m + 1), r);
}


int max(vector<pair<long long, int>>& st, int v, int tl, int tr, int l, int r)
{
	if (l > r) return 0;
	int m = (tl + tr) / 2;
	if (l == tl && r == tr) return st[v].second;
	return max(max(st, v * 2, tl, m, l, min(m, r)), max(st, v * 2 + 1, m + 1, tr, max(l, m + 1), r));
}

int down(vector<pair<long long, int>>& st, int v, int tl, int tr, long long sum)
{
	if (tr - tl == 0) return tl;
	int m = (tl + tr) / 2;
	if ((long long)st[v * 2].second >= sum) return down(st, v * 2, tl, m, sum);
	return down(st, v * 2 + 1, m, tr, sum);
}

int findFirstMax(vector<pair<long long, int>>& st, int v, int tl, int tr, int l, int r, long long sum, int n)
{
	if (l >= r) return n;
	if (tl == l && tr == r)
	{
		if ((long long)st[v].second < sum) return n;
		return down(st, v, tl, tr, sum);
	}

	int m = (tl + tr) / 2;
	int res = findFirstMax(st, v * 2, tl, m, l, min(m, r), sum, n);
	if (res != n) return res;
	return findFirstMax(st, v * 2 + 1, m, tr, max(m, l), r, sum, n);
}

void modify(vector<pair<long long, int>>& st, int v, int l, int r, int pos, int new_val)
{
	if (l == r)
	{
		st[v].first = new_val;
		st[v].second = new_val;
		return;
	}

	int m = (l + r) / 2;
	if (pos <= m) modify(st, v * 2, l, m, pos, new_val);
	else modify(st, v * 2 + 1, m + 1, r, pos, new_val);
	st[v].first = st[v * 2].first + st[v * 2 + 1].first;
	st[v].second = max(st[v * 2].second, st[v * 2 + 1].second);
}

int ask(vector<pair<long long, int>>& st, vector<int>& a, int n)
{
	int currPos = 1;
	if (a[0] == 0) return 1;
	else while (currPos != n)
	{
		long long s = sum(st, 1, 0, n - 1, 0, currPos - 1);
		if (s > 1e9) break;
		if (s == (long long)a[currPos]) return currPos + 1;
		int max_pos = findFirstMax(st, 1, 0, n, currPos + 1, n, s, n);
		currPos = max_pos;
	}
	return -1;
}

void doTask(istream& in, ostream& out)
{
	int n, q;
	in >> n >> q;
	vector<int> a(n);
	for (int i = 0; i < n; ++i) in >> a[i];

	vector<pair<long long, int>> st(4 * n + 1, make_pair(0, 0));

	buildTree(st, a, 1, 0, n - 1);

	for (int i = 0; i < q; ++i)
	{
		int pi, xi;
		in >> pi >> xi;
		a[pi - 1] = xi;
		modify(st, 1, 0, n - 1, pi - 1, xi);
		out << ask(st, a, n) << endl;
	}
}