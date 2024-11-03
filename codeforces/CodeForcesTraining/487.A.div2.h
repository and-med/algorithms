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
		if (s[i] != s[i - 1] && s[i - 1] != s[i - 2] && s[i] != s[i - 2] && s[i] != '.' && s[i-1] != '.' && s[i-2] != '.')
		{
			exists = true;
			break;
		}
	}

	out << (exists ? "YES" : "NO") << endl;
}