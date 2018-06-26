#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;

vector<vector<int>> moves(3, vector<int>());

void moveCar(vector<vector<int>>& cars, int crow, int ccol, int row, int col)
{
	int k = cars[crow][ccol];
	cars[crow][ccol] = 0;
	cars[row][col] = k;
	moves[0].push_back(k);
	moves[1].push_back(row + 1);
	moves[2].push_back(col + 1);
}

int moveCars(vector<vector<int>>& cars, int n)
{
	int carsMoved = 0;
	for (int i = 1; i <= 2; ++i)
	{
		for (int j = 0; j < n; ++j)
		{
			int p = 3 * (i - 1);
			if (cars[i][j] == cars[p][j] && cars[p][j] != 0)
			{
				moveCar(cars, i, j, p, j);
				carsMoved++;
			}
		}
	}
	return carsMoved;
}

pair<int, int> prevPosition(int i, int j, int n)
{
	int movei = 0, movej = 0;
	if (j >= 0 && j < n - 1 && i == 1)	movei = 0, movej = 1;
	if (j > 0 && j < n		&& i == 2)	movei = 0, movej = -1;
	if (j == n - 1 && i == 1)	movei = 1, movej = 0;
	if (j == 0 && i == 2)	movei = -1, movej = 0;

	return make_pair(i + movei, j + movej);
}

void roundTrip(vector<vector<int>>& cars, int n, int k)
{
	pair<int, int> next = make_pair(1, 0);
	pair<int, int> curr = prevPosition(next.first, next.second, n);
	pair<int, int> prev = prevPosition(curr.first, curr.second, n);
	while (k != 0)
	{
		if (cars[next.first][next.second] == 0 && cars[curr.first][curr.second] != 0)
		{
			moveCar(cars, curr.first, curr.second, next.first, next.second);
			k--;
		}
		next.first = curr.first; next.second = curr.second;
		curr.first = prev.first; curr.second = prev.second;
		prev = prevPosition(curr.first, curr.second, n);
	}
}

void doTask(istream& in, ostream& out)
{
	int n, k;
	in >> n >> k;
	vector<vector<int>> cars(4, vector<int>(n));

	for (int i = 0; i < 4; ++i) for (int j = 0; j < n; ++j) in >> cars[i][j];

	int moved = moveCars(cars, n);
	k -= moved;
	if (moved == 0 && k == 2 * n)
	{
		out << -1 << endl;
		return;
	}

	while (k > 0)
	{
		roundTrip(cars, n, k);
		moved = moveCars(cars, n);
		k -= moved;
	}

	out << moves[0].size() << endl;
	for (int i = 0; i < moves[0].size(); ++i)
	{
		out << moves[0][i] << " " << moves[1][i] << " " << moves[2][i] << endl;
	}
}