#include <iostream>
#include<vector>
using namespace std;

int main()
{
    // Calculates all fib numbers up to X limit.
    auto const limit{4000000};

    vector<int> fib_seq{1, 2};
    auto i{0};

    while (i <= limit)
    {
        i = fib_seq[fib_seq.size() - 1] + fib_seq[fib_seq.size() - 2];
        fib_seq.push_back(i);
    }

    // Get the total of all even fib numbers
    auto total{0};

    for (auto i = 0; i < fib_seq.size(); i++)
    {
        if (fib_seq[i] % 2 == 0)
        {
            total += fib_seq[i];
        }
    }

    cout << "The total of all even fib numbers under " << limit << ": " << total << endl;
}