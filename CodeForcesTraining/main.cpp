#ifdef LOCAL
#include<fstream>
#endif // LOCAL
#include"492.D.div2.h"

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