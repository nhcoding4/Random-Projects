#include <stdio.h>
#include <time.h>
#include <stdlib.h>

int main()
{
    // Randomly selects words from the array and creates a sentence out of them.

    char *words[5][3] = {
        {"Apples", "Oranges", "Pears"},
        {"is", "was", "could be"},
        {"looking", "moving", "speaking"},
        {"at", "to", "towards"},
        {"the floor.", "the door.", "the wall."},
    };

    srand(time(0));

    for (int i = 0; i < 5; i++)
    {
        printf("%s ", words[i][rand() % (3)]);
    }
    printf("\n");
}
