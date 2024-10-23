#include <iostream>
#include <vector>

using namespace std;

int main()
{
    // Get all multiples of 3 and 5 below a certain limit
    auto const limit{1000};
    vector<int> numbers{};

    for (auto i = 0; i < limit; i++)
    {
        if (i % 3 == 0 || i % 5 == 0)
        {
            numbers.push_back(i);
        }
    }

    // Find the total of all the multiples.
    auto total{0};
    for (auto i = 0; i < numbers.size(); i++)
    {
        total += numbers[i];
    }

    cout << "The total of all multiples of 3 and 5 below " << limit << ": " << total << endl;
}