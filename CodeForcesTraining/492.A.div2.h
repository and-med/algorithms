#include<iostream>
using namespace std;

void doTask(istream& in, ostream& out)
{
	int n;
	in >> n;

	int count = n / 100;
	n = n % 100;
	count += n / 20;
	n = n % 20;
	count += n / 10;
	n = n % 10;
	count += n / 5;
	n = n % 5;
	count += n / 1;

	out << count << endl;
}