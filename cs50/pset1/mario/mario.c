#include <stdio.h>
#include <cs50.h>

int main(void)
{
    int height;

    do
    {
        printf("Height: ");

        height = GetInt();

    } while (height < 0 || height > 23);

    for (int index = 1; index <= height; index++)
    {
        for (int i = 0; i < height - index; i++)
            printf(" ");

        for (int i = 0; i < index; i++)
            printf("#");

        // Newline and last # after our loop #s
        printf("#\n");
    }
}
