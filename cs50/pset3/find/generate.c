/**
 * generate.c
 *
 * Generates pseudorandom numbers in [0,MAX), one per line.
 *
 * Usage: generate n [s]
 *
 * where n is number of pseudorandom numbers to print
 * and s is an optional seed
 */
 
#define _XOPEN_SOURCE

#include <cs50.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

// upper limit on range of integers that can be generated
#define LIMIT 65536

int main(int argc, string argv[])
{
    // If we don't have 2 or 3 arguments we display usage and exit
    if (argc != 2 && argc != 3)
    {
        printf("Usage: ./generate n [s]\n");
        return 1;
    }

    // We convert our first argument to a number, specifying the number of
    // random numbers we want
    int n = atoi(argv[1]);

    // If we have 3 arguments we seed rand48 with the second argument
    if (argc == 3)
    {
        srand48((long) atoi(argv[2]));
    }
    // We seed our srand with the current time in epoch format
    else
    {
        srand48((long) time(NULL));
    }

    // We print a number using the drand48 which gives us a number between [0.0, 1.0),
    // and multiply it by "LIMIT" to get a random int
    for (int i = 0; i < n; i++)
    {
        printf("%i\n", (int) (drand48() * LIMIT));
    }

    // success
    return 0;
}
