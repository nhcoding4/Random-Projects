#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

// --------------------------------------------------------------------------------------------------------------------

char *get_string(FILE *stream);
char *reverse(char *base_string);

// --------------------------------------------------------------------------------------------------------------------

int main()
{
    printf("Enter some words to reverse the order of:\n");

    char *user_input = get_string(stdin);
    char *reversed_string = reverse(user_input);
    printf("%s reversed is:\n%s\n", user_input, reversed_string);

    free(user_input);
    free(reversed_string);
}

// --------------------------------------------------------------------------------------------------------------------

// Takes a string and reverses the order of that string.
char *reverse(char *base_string)
{
    int length = strlen(base_string);
    char *reversed_string = malloc(sizeof(char) * length + 1);
    int index = 0;

    for (int i = length - 1; i >= 0; i--)
    {
        reversed_string[index] = base_string[i];
        index++;
    }

    reversed_string[length] = 0;

    return reversed_string;
}

// --------------------------------------------------------------------------------------------------------------------

// Takes input from a stream.
char *get_string(FILE *stream)
{
    int bytes = 0;
    int capacity = 50;
    char *buf = malloc(capacity);
    char c;

    while ((c = fgetc(stream)) != EOF && c != '\r' && c != '\n')
    {
        bytes++;

        if (bytes + 1 >= capacity)
        {
            capacity *= 2;
            buf = realloc(buf, capacity);
            if (buf == NULL)
            {
                return NULL;
            }
        }

        buf[bytes - 1] = c;
    }

    if (c == '\r')
    {
        c = fgetc(stream);
        if (c != '\n')
        {
            ungetc(c, stream);
        }
    }

    if (bytes == 0)
    {
        if (c == EOF)
        {
            free(buf);
            return NULL;
        }

        else
        {
            buf = malloc(1);
        }
    }

    buf[bytes] = 0;

    return buf;
}

// --------------------------------------------------------------------------------------------------------------------
