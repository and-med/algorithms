#include<iostream>
#include<fstream>
using namespace std;

int main() {
    int t, n;

    ifstream file("text.txt");
    cin >> t;
    for (int k = 0; k < t; ++k) {
        cin >> n;
        long long max = 0;
        long long sum = 0;
        long long a;
        for (int i = 0; i < n; ++i) {
            cin >> a;
            max = std::max(max, a);
            sum = sum + a;
        }

        long long roundedSum = sum - (sum % (n - 1)) + (sum % (n - 1) != 0 ? n - 1 : 0);
        long long koef = std::max(roundedSum/(n - 1), max);
        long long minimum = koef * (n - 1);

        cout << minimum - sum << endl;
    }
}