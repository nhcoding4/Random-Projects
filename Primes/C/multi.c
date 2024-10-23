#include <math.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>
#include <pthread.h>
#include <stdio.h>

// --------------------------------------------------------------------------------------------------------------------

// Structure definition.

typedef struct
{
    uint64_t start, end, result, capacity;
    uint64_t *found;
} Range;

// --------------------------------------------------------------------------------------------------------------------

// Forward definitions.

bool is_prime(uint64_t number);
void *calculate(void *range);
void execute_work(Range holding_array[], const uint64_t *workers);
void free_memory(Range holding_array[], const uint64_t *workers);
void print_results(Range holding_array[], const uint64_t *workers);
void split_work(Range holding_array[], const uint64_t *workers, const uint64_t *total_value);

// --------------------------------------------------------------------------------------------------------------------

int main()
{
    const uint64_t total_workers = 24;
    const uint64_t total_value = 250001;
    Range ranges[24] = {};

    split_work(ranges, &total_workers, &total_value);

    execute_work(ranges, &total_workers);

    print_results(ranges, &total_workers);

    free_memory(ranges, &total_workers);
}

// --------------------------------------------------------------------------------------------------------------------

// Brute force checks a if a number is prime. Inefficient by design.

bool is_prime(uint64_t number)
{
    for (uint64_t i = 2; i < number; i++)
    {
        if (number % i == 0)
        {
            return false;
        }
    }
    return true;
}

// --------------------------------------------------------------------------------------------------------------------

// Thread logic.

void *calculate(void *range)
{
    Range *data = (Range *)range;

    for (uint64_t i = data->start; i < data->end; i++)
    {
        if (is_prime(i))
        {
            data->found[data->result] = i;
            data->result++;
        }

        if (data->result == data->capacity - 1)
        {
            data->capacity += 1000;
            data->found = realloc(data->found, sizeof(uint64_t) * data->capacity);
        }
    }
    return NULL;
}

// --------------------------------------------------------------------------------------------------------------------

// Create threads and tells them to execute work.

void execute_work(Range holding_array[], const uint64_t *workers)
{
    pthread_t thread[24];

    for (uint64_t i = 0; i < *workers; i++)
    {
        int err;
        err = pthread_create(&thread[i], NULL, calculate, (void *)&holding_array[i]);
        if (err)
        {
            printf("Error creating thread\n");
            exit(1);
        }
    }

    for (uint64_t i = 0; i < *workers; i++)
    {
        pthread_join(thread[i], NULL);
    }
}

// --------------------------------------------------------------------------------------------------------------------

// Free allocated memory.

void free_memory(Range holding_array[], const uint64_t *workers)
{
    for (uint64_t i = 0; i < *workers; i++)
    {
        free(holding_array[i].found);
    }
}

// --------------------------------------------------------------------------------------------------------------------

// Sums the total found primes.

void print_results(Range holding_array[], const uint64_t *workers)
{
    uint64_t total_found = 0;

    for (uint64_t i = 0; i < *workers; i++)
    {
        total_found += holding_array[i].result;
    }

    printf("Found primes: %d\n", total_found);
}

// --------------------------------------------------------------------------------------------------------------------

// Create ranges for each thread to work on.

void split_work(Range holding_array[], const uint64_t *workers, const uint64_t *total_value)
{
    uint64_t current_base = 2;
    const uint64_t increase = round(*total_value / *workers);

    for (uint64_t i = 0; i < *workers; i++)
    {
        uint64_t end_value = current_base + increase;

        if (i == *workers - 1 && end_value != *total_value)
        {
            end_value = *total_value;
        }

        Range new_range = {
            .start = current_base,
            .end = end_value,
            .result = 0,
            .capacity = 1000,
            .found = (uint64_t *)malloc(sizeof(uint64_t) * new_range.capacity),
        };

        holding_array[i] = new_range;
        current_base += increase;
    }
}

// --------------------------------------------------------------------------------------------------------------------
