#include <stdio.h>
#include <stdlib.h>
#include <cs50.h>
#include <string.h>
#include <stdbool.h>

char vigenereShift(char, int);
void vigenereCipher(string, string);
bool isUpper(char);
bool isLower(char);

int main(int argc, string argv[])
{
    // If we don't have exactly two arguments
    if (argc != 2) {
        printf("Usage: ./vigenere k\n");
        return 1;
    }

    // Our first argument is the key we'll use for encryption
    string key = argv[1];

    // We receive a string from the user to encrypt
    printf("plaintext:  ");
    string encrypt = GetString();

    vigenereCipher(encrypt, key);

    return 0;
}

// Does the Vigenere cipher letter shift
char vigenereShift(char letter, char offset)
{

}

// Sequence that prints our caesar shifted values
void vigenereCipher(string encrypt, string key)
{
    for (int i = 0, n = strlen(encrypt); i < n; i++) {
        printf("%c", vigenereShift(encrypt[i], offset));
    }

    // Newline to finish our call
    printf("\n");
}

