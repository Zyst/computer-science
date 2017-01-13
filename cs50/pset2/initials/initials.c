#include <stdio.h>
#include <cs50.h>
#include <string.h>
#include <ctype.h>

void getInitials(string);

int main(void)
{
    string name = GetString();

    getInitials(name);
}

// This is a sequence that does the meat of our work
// I originally wanted this to return a string, but
// I have no idea of how you functionally compose arrays
// in this kind of language
void getInitials(string name)
{

    // If our first character is a letter
    if (isalpha(name[0]))
    {
        // We print our first letter
        printf("%c", toupper(name[0]));
    }

    // We iterate through our array, up to n - 1 since our check is for i + 1
    for (int i = 1, n = strlen(name); i < n - 1; i++)
    {
        // If we have a space and our next letter is alphabetical
        if (isspace(name[i]) && isalpha(name[i + 1]))
        {
            printf("%c", toupper(name[i + 1]));
        }
    }

    // Program is complete, we output a newline
    printf("\n");
}
