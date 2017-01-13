#include <stdio.h>
#include <cs50.h>
#include <math.h>

int giveChange(int money, int coins)
{
    if (money >= 25)
    {
        coins += money / 25;
        money = money % 25;
    }
    if (money >= 10)
    {
        coins += money / 10;
        money = money % 10;
    }
    if (money >= 5)
    {
        coins += money / 5;
        money = money % 5;
    }
    if (money >= 1)
    {
        coins += money / 1;
        money = money % 1; // Should always be 0
    }
    return coins;
}

int main(void)
{
    float money;

    printf("O hai! ");

    do
    {
        printf("How much change is owed?\n");

        money = GetFloat();

    } while (money <= 0);

    int moneyInt = round(money * 100.0);
    int coins = 0;

    printf("%d\n", giveChange(moneyInt, coins));
}
