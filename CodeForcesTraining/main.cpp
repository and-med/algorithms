#ifdef LOCAL
#include<fstream>
#endif // LOCAL
#include<iostream>
#include<string>
using namespace std;

void doTask(istream& in, ostream& out)
{
	string s;
	in >> s;

	bool exists = false;
	for (int i = 2; i < s.length(); ++i)
	{
		if (s[i] != s[i - 1] && s[i - 1] != s[i - 2] && s[i] != s[i - 2] && s[i] != '.' && s[i - 1] != '.' && s[i - 2] != '.')
		{
			exists = true;
			break;
		}
	}

	out << (exists ? "YES" : "NO") << endl;
}

int main()
{
#ifdef LOCAL
	std::ifstream in("intput.txt");
	std::ofstream out("output.txt");
	std::cout << "LOCAL ENVIRONMNET: using files" << std::endl;
	doTask(in, out);
#else
	doTask(std::cin, std::cout);
#endif // LOCAL
}