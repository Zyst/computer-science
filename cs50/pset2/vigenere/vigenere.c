#include <stdio.h>
#include <stdlib.h>
#include <cs50.h>
#include <string.h>
#include <stdbool.h>

char vigenereShift(char, char);
void vigenereCipher(string, string);
bool isUpper(char);
bool isLower(char);

int main(int argc, string argv[])
{
  // If we don't have exactly two arguments
  if (argc != 2)
  {
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

bool isUpper(char letter)
{
  return letter >= 'A' && letter <= 'Z';
}

bool isLower(char letter)
{
  return letter >= 'a' && letter <= 'z';
}

// Does the Vigenere cipher letter shift
char vigenereShift(char letter, char offset)
{
  if (isUpper(letter))
  {

    // If our offset is lowercase we turn it into uppercase
    // because Vigenere keys are case insensitive
    if (isLower(offset))
    {
      offset -= 32;
    }

    // We return the letter turned into an alphabetical index (0-25 A-Z) plus
    // the offset turned into an alphabetical index, then we do a modulo
    // 26 operation, and finally turn the letter into the ASCII range again
    return ((letter - 'A' + offset - 'A') % 26) + 'A';
  }
  else if (isLower(letter))
  {
    // If our offset is uppercase we turn it into lowercase
    // because Vigenere keys are case insensitive
    if (isUpper(offset))
    {
      offset += 32;
    }

    // We return the letter turned into an alphabetical index (0-25 a-z) plus
    // the offset turned into an alphabetical index, then we do a modulo
    // 26 operation, and finally turn the letter into the ASCII range again
    return ((letter - 'a' + offset - 'a') % 26) + 'a';
  }
  // If we got to this point we return the character unchanged
  return letter;
}

// Sequence that prints our caesar shifted values
void vigenereCipher(string encrypt, string key)
{
  printf("ciphertext: ");

  for (int i = 0, n = strlen(encrypt); i < n; i++)
  {
    printf("%c", vigenereShift(encrypt[i], key[0]));
  }

  // Newline to finish our call
  printf("\n");
}
