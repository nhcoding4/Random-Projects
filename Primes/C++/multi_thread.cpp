// 0.76s.

#include <iostream>
#include <future>
#include <vector>

using namespace std;

// --------------------------------------------------------------------------------------------------------------------

typedef struct
{
    int workers;
    int limit;
    vector<future<vector<int>>> task_pool;
} Pool;

// --------------------------------------------------------------------------------------------------------------------

// Checks if a number is prime or not.
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

// --------------------------------------------------------------------------------------------------------------------

// Finds primes for a range of numbers and saves them to a vector.
vector<int> find_primes(int min, int max)
{
    vector<int> primes_found{};
    for (auto i = min; i < max; i++)
    {
        if (is_prime(i))
        {
            primes_found.push_back(i);
        }
    }
    return primes_found;
}

// --------------------------------------------------------------------------------------------------------------------

int main()
{
    // Set up a pool that holds the tasks to be performed and the amount of workers to distribute the task between.
    auto tasks{Pool{}};
    tasks.workers = 24;
    tasks.limit = 250001;

    // Set the range of numbers for each worker to work on.
    auto min{2};
    auto increase{int(tasks.limit / tasks.workers)};

    // Create tasks and change range of numbers.
    for (auto i = 0; i < tasks.workers; i++)
    {
        if (min > tasks.limit)
        {
            min = tasks.limit;
        }
        tasks.task_pool.push_back(future<vector<int>>{async(&find_primes, min, min + increase)});
        min += increase;
    }

    // Get and sum the results.
    auto results{0};
    for (auto i = 0; i < tasks.task_pool.size(); i++)
    {
        auto result{tasks.task_pool[i].get()};
        results += result.size();
    }

    cout << "Found: " << results << " primes." << endl;
}

// --------------------------------------------------------------------------------------------------------------------
