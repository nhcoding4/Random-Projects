#include <stdio.h>
#include <stdlib.h>

// --------------------------------------------------------------------------------------------------------------------

typedef struct
{
    int *data;
    int size;
} Vector;

// --------------------------------------------------------------------------------------------------------------------

Vector *createVector();
Vector *mergeSort(Vector *vector);
Vector *sort(Vector *left, Vector *right);

void freeVector(Vector *vector);
void push(Vector *vector, int element);
void printVector(Vector *vector);