/**
 * helpers.c
 *
 * Helper functions for Problem Set 3.
 */

#include <cs50.h>

#include "helpers.h"

/**
 * Returns true if value is in array of n values, else false.
 */
bool search(int value, int values[], int n)
{
    // If n is a negative number just return false, as per the spec.
    if (n < 1)
    {
        return false;
    }

    // Start is the start index
    int start = 0;
    // End is the end index
    int end = n - 1;

    while (start <= end)
    {
        int middle = (start + end) / 2;

        if (value == values[middle])
        {
            return true;
        }

        if (value > values[middle])
        {
            start = middle + 1;
        }
        else
        {
            end = middle - 1;
        }
    }

    return false;
}

/**
 * Sorts array of n values.
 */
void sort(int values[], int n)
{
    // We need to know the max number of the passed array so we can implement counting sort
    const int MAX = 65536;
    int counted[MAX];

    for (int i = 0; i < MAX; i++)
    {
        counted[i] = 0;
    }

    for (int i = 0; i < n; i++)
    {
        // We add the value, minus one because that's the relevant index. And ++ it.
        // I think this might cause trouble? Like, I think C memory can die this way.
        counted[values[i] - 1]++;
    }

    for (int i = 0, current = 0; i < MAX; i++)
    {
        if (counted[i] > 0)
        {
            // For each count element
            for (int j = 0; j < counted[i]; j++)
            {
                // The value is our current iteration - 1
                values[current] = i + 1;
                // We ++ current so we are on the next array index
                current++;
            }
        }
    }
}
