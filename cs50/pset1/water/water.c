#include <stdio.h>
#include <cs50.h>

int minutesToWaterBottles(int min) {
    return min * 192 / 16;
}

int main(void) {
    printf("minutes: ");
    int minutes = GetInt();

    printf("bottles: %d\n", minutesToWaterBottles(minutes));
}
