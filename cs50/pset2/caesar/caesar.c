#include <stdio.h>
#include <stdlib.h>
#include <cs50.h>
#include <string.h>

char caesarShift(char, int);
void caesarCipher(string, int);

int main(int argc, string argv[])
{
    // If we don't have exactly two arguments
    if (argc != 2)
    {
        printf("Caesar should have exactly one command line argument.\n");
        return 1;
    }

    // Our first argument becomes the offset we use for encryption
    int offset = atoi(argv[1]);

    // We receive a string from the user to encrypt
    string encrypt = GetString();

    caesarCipher(encrypt, offset);

    return 0;
}

// Does the caesar cipher letter shift
char caesarShift(char letter, int offset)
{
    // If we have an uppercase or lowercase letter
    if (letter >= 'A' && letter <= 'Z')
    {
        return (((letter - 'A') + offset) % 26) + 'A';
    }
    else if (letter >= 'a' && letter <= 'z')
    {
        return (((letter - 'a') + offset) % 26) + 'a';
    }

    // If not alphanumeric we change nothing
    return letter;
}

// Sequence that prints our caesar shifted values
void caesarCipher(string encrypt, int offset)
{
    for (int i = 0, n = strlen(encrypt); i < n; i++)
    {
        printf("%c", caesarShift(encrypt[i], offset));
    }

    // Newline to finish our call
    printf("\n");
}
