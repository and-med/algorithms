#include<iostream>
#include<string>
#include<algorithm>

bool is25Divisor(std::string n)
{
	if (n.length() < 2) return false;

	std::string lastTwo = n.substr(n.length() - 2, 2);
	return lastTwo == "00" || lastTwo == "25" || lastTwo == "50" || lastTwo == "75";
}

void doTask(std::istream& in, std::ostream& out)
{
	std::string n;
	in >> n;

	int ans = INT_MAX;
	std::string original = n;
	for (int i = 0; i < n.length(); ++i)
	{
		for (int j = 0; j < n.length(); ++j)
		{
			if (i == j)
			{
				continue;
			}

			n = original;
			int p1 = i;
			int p2 = j;
			int swaps = 0;

			if (p2 < p1)
			{
				p1--;
			}

			while (p2 != n.length() - 1)
			{
				std::swap(n[p2], n[p2 + 1]);
				p2++;
				swaps++;
			}

			while (p1 != n.length() - 2)
			{
				std::swap(n[p1], n[p1 + 1]);
				p1++;
				swaps++;
			}

			int notZero = 0;
			while (n[notZero] == '0') notZero++;

			while (notZero != 0)
			{
				std::swap(n[notZero], n[notZero - 1]);
				notZero--;
				swaps++;
			}

			if (is25Divisor(n))
			{
				ans = std::min(ans, swaps);
			}
		}
	}

	if (ans == INT_MAX)
	{
		out << -1 << std::endl;
	}
	else
	{
		out << ans << std::endl;
	}
}

int main()
{
	doTask(std::cin, std::cout);
}