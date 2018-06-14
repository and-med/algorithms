#include<iostream>
#include<algorithm>

long long getSum(int* a, int n)
{
	long long sum = 0;
	for (int i = 0; i < n; ++i)
	{
		sum += a[i];
	}
	return sum;
}

struct arrEntry
{
	long long sum;
	int i;
	int j;
};

void doTask()
{
	int k;
	std::cin >> k;
	arrEntry* arr = new arrEntry[2 * 100000];

	int index = 0;
	for (int i = 0; i < k; ++i)
	{
		int n;
		std::cin >> n;
		int* a = new int[n];
		for (int j = 0; j < n; j++)
		{
			std::cin >> a[j];
		}
		long long sum = getSum(a, n);
		for (int j = 0; j < n; j++)
		{
			arr[index].sum = sum - a[j];
			arr[index].i = i;
			arr[index].j = j;
			index++;
		}
		delete a;
	}

	std::sort(arr, arr + index, [](const arrEntry& left, const arrEntry& right) { return left.sum < right.sum || (left.sum == right.sum && left.i < right.i); });

	int isSatisfied = false;
	for (int i = 0; i < index - 1; ++i)
	{
		if (arr[i].sum == arr[i + 1].sum && arr[i].i != arr[i + 1].i)
		{
			std::cout << "YES" << std::endl;
			std::cout << arr[i].i + 1 << " " << arr[i].j + 1 << std::endl;
			std::cout << arr[i + 1].i + 1 << " " << arr[i + 1].j + 1 << std::endl;
			isSatisfied = true;
			break;
		}
	}
	if (!isSatisfied)
	{
		std::cout << "NO" << std::endl;
	}
}