#include<iostream>
#include<algorithm>

void doTask()
{
	int* a = new int[100];
	int n = 0, k = 0;
	std::cin >> n >> k;
	for (int i = 0; i < n; ++i)
	{
		std::cin >> a[i];
	}


	int* different = new int[k];
	different[0] = 0;
	int countDifferent = 1;

	for (int i = 1; i < n && countDifferent < k; ++i)
	{
		bool isDifferent = true;
		for (int j = 0; j < countDifferent; ++j)
		{
			if (a[different[j]] == a[i])
			{
				isDifferent = false;
				break;
			}
		}

		if (isDifferent)
		{
			different[countDifferent++] = i;
		}
	}

	if (countDifferent == k)
	{
		std::cout << "YES" << std::endl;
		for (int i = 0; i < k; ++i)
		{
			std::cout << different[i] + 1 << " ";
		}
		std::cout << std::endl;
	}
	else
	{
		std::cout << "NO" << std::endl;
	}
}