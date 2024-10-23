#include <stdio.h>
#include <time.h>

int check(int number);

int main()
{
    clock_t timer;
    timer = clock();

    int number = 1;
    int found = 0;

    while (!found)
    {
        found = check(number);
        number++;
    }

    timer = clock() - timer;
    printf("Total time taken: %.5f seconds\n", ((float)timer) / CLOCKS_PER_SEC);
}

int check(int number)
{
    for (int i = 11; i <= 20; i++)
    {
        if (number % i != 0)
        {
            return 0;
        }
    }

    printf("Found number %i\n", number);
    return 1;
}