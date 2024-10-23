#include "./def.h"

// --------------------------------------------------------------------------------------------------------------------

int main()
{
    Vector *vector = createVector();

    for (int i = 20; i > 0; i--)
    {
        push(vector, i);
    }
    printVector(vector);

    Vector *sortedVector = mergeSort(vector);
    printVector(sortedVector);
    freeVector(sortedVector);

    Vector *vector2 = createVector();

    for (int i = 19; i > 0; i--)
    {
        push(vector2, i);
    }
    printVector(vector2);

    Vector *sortedVector2 = mergeSort(vector2);
    printVector(sortedVector2);
    freeVector(sortedVector2);

    exit(0);
}

// --------------------------------------------------------------------------------------------------------------------
