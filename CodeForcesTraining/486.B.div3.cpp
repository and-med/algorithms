#include<iostream>
#include<algorithm>
#include<string>

bool isSubstring(const std::string& str, const std::string& value)
{
	for (int i = 0; i < str.length(); ++i)
	{
		bool isSubstring = true;
		int currentI = i;
		for (int j = 0; j < value.length(); ++j, ++i)
		{
			if (str[i] != value[j])
			{
				isSubstring = false;
				break;
			}
		}

		i = currentI;
		if (isSubstring)
		{
			return true;
		}
	}
	return false;
}

void doTask()
{
	std::string* b = new std::string[100];
	int n;
	std::cin >> n;
	for (int i = 0; i < n; ++i)
	{
		std::cin >> b[i];
	}

	std::sort(b, b + n,
		[](const std::string& left, const std::string& right)
	{
		return left.length() < right.length() || (left.length() == right.length() && left < right);
	});

	bool isSatisfying = true;
	for (int i = 0; i < n - 1; ++i)
	{
		if (!isSubstring(b[i + 1], b[i]))
		{
			isSatisfying = false;
			break;
		}
	}

	if (isSatisfying)
	{
		std::cout << "YES" << std::endl;
		for (int i = 0; i < n; ++i)
		{
			std::cout << b[i] << std::endl;
		}
	}
	else
	{
		std::cout << "NO" << std::endl;
	}
}