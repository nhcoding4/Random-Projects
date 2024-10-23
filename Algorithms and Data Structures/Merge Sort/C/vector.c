#include "def.h"

// --------------------------------------------------------------------------------------------------------------------

// Return a pointer to a new vector.

Vector *createVector()
{
    Vector *vector = (Vector *)malloc(sizeof(Vector) * 1);
    if (vector == NULL)
    {
        printf("createVector: error allocating memory\n");
        exit(1);
    }

    vector->size = 0;
    vector->data = NULL;

    return vector;
}

// --------------------------------------------------------------------------------------------------------------------

// Free memory.

void freeVector(Vector *vector)
{
    free(vector->data);
    free(vector);
}

// --------------------------------------------------------------------------------------------------------------------

// Prints all the values held in the vector.

void printVector(Vector *vector)
{
    for (int i = 0; i < vector->size; i++)
    {
        printf("%i ", vector->data[i]);
    }
    printf("\n");
}

// --------------------------------------------------------------------------------------------------------------------

// Adds an integer to the back of the vector.

void push(Vector *vector, int element)
{
    vector->size++;

    int *newData = (int *)malloc(sizeof(int) * vector->size + 1);
    if (newData == NULL)
    {
        printf("vector push: error allocating memory\n");
        exit(1);
    }

    if (vector->data != NULL)
    {
        for (int i = 0; i < vector->size; i++)
        {
            newData[i] = vector->data[i];
        }
    }

    newData[vector->size - 1] = element;

    free(vector->data);
    vector->data = newData;
}

// --------------------------------------------------------------------------------------------------------------------

// Split and sort vectors until base case is hit

Vector *mergeSort(Vector *vector)
{
    if (vector->size == 1)
    {
        return vector;
    }

    int leftSize = vector->size / 2;
    int rightSize = vector->size / 2;

    if (vector->size % 2 != 0)
    {
        leftSize += 1;
    }

    Vector *left = createVector();
    Vector *right = createVector();

    int index = 0;

    for (int i = 0; i < leftSize; i++)
    {
        push(left, vector->data[index]);
        index++;
    }

    for (int i = 0; i < rightSize; i++)
    {
        push(right, vector->data[index]);
        index++;
    }

    freeVector(vector);

    Vector *newLeft = mergeSort(left);
    Vector *newRight = mergeSort(right);

    return sort(newLeft, newRight);
}

// --------------------------------------------------------------------------------------------------------------------

// Sort 2 vectors of integers from low to high values.

Vector *sort(Vector *left, Vector *right)
{
    int i = 0;
    int j = 0;

    Vector *sortedData = createVector();

    while (i < left->size && j < right->size)
    {
        if (left->data[i] < right->data[j])
        {
            push(sortedData, left->data[i]);
            i++;
        }
        else
        {
            push(sortedData, right->data[j]);
            j++;
        }
    }

    while (i < left->size)
    {
        push(sortedData, left->data[i]);
        i++;
    }

    while (j < right->size)
    {
        push(sortedData, right->data[j]);
        j++;
    }

    freeVector(left);
    freeVector(right);

    return sortedData;
}

// --------------------------------------------------------------------------------------------------------------------
