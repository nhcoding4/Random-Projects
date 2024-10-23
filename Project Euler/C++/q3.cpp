#include <iostream>
#include <vector>
#include <bits/stdc++.h>

using namespace std;

int main()
{
    vector<int> factors{};
    auto prime{2};
    auto n{600851475143};

    while (n > 1)
    {
        while (n % prime == 0)
        {
            factors.push_back(prime);
            n /= prime;
        }
        prime++;

        if ((prime * prime) > n)
        {
            if (n > 1)
            {
                factors.push_back(n);
                break;
            }
        }
    }

    sort(factors.begin(), factors.end());

    cout << "The largest prime factor: " << factors[factors.size() - 1] << endl;
}