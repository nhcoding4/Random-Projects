#include <iostream>

int main()
{

    // Take input
    std::cout << "Enter a max number\n";
    int limit{};
    std::cin >> limit;

    // Fizzbuzzing
    for (int i = 0; i < limit; i++)
    {
        if (i % 5 == 0)
        {
            if (i % 3 == 0)
            {
                std::cout << "FizzBuzz";
            }
            else
            {
                std::cout << "Buzz";
            }
        }
        else if (i % 3 == 0)
        {
            std::cout << "Fizz";
        }
        else
        {
            std::cout << i;
        }
        std::cout << std::endl;
    }
}