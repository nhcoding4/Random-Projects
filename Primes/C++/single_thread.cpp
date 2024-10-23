// 5.9 seconds

#include <iostream>
#include <vector>

using namespace std;

bool is_prime(int n)
{
    for (auto i = 2; i < n; i++)
    {
        if (n % i == 0)
        {
            return false;
        }
    }
    return true;
}

int main()
{
    vector<int> primes{};

    for (auto i = 2; i < 250001; i++)
    {
        if (is_prime(i))
        {
            primes.push_back(i);
        }
    }

    cout << "Found " << primes.size() << " prime numbers\n";
}