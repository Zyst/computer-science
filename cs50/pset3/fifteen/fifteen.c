/**
 * fifteen.c
 *
 * Implements Game of Fifteen (generalized to d x d).
 *
 * Usage: fifteen d
 *
 * whereby the board's dimensions are to be d x d,
 * where d must be in [DIM_MIN,DIM_MAX]
 *
 * Note that usleep is obsolete, but it offers more granularity than
 * sleep and is simpler to use than nanosleep; `man usleep` for more.
 */

#define _XOPEN_SOURCE 500

#include <cs50.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

// constants
#define DIM_MIN 3
#define DIM_MAX 9

// board
int board[DIM_MAX][DIM_MAX];
int win_state[DIM_MAX][DIM_MAX];

// dimensions
int d;

// prototypes
void clear(void);
void greet(void);
void init(void);
void draw(void);
bool move(int tile);
bool won(void);

int main(int argc, string argv[])
{
    // ensure proper usage
    if (argc != 2)
    {
        printf("Usage: fifteen d\n");
        return 1;
    }

    // ensure valid dimensions
    d = atoi(argv[1]);
    if (d < DIM_MIN || d > DIM_MAX)
    {
        printf("Board must be between %i x %i and %i x %i, inclusive.\n",
               DIM_MIN, DIM_MIN, DIM_MAX, DIM_MAX);
        return 2;
    }

    // open log
    FILE *file = fopen("log.txt", "w");
    if (file == NULL)
    {
        return 3;
    }

    // greet user with instructions
    greet();

    // initialize the board
    init();

    // accept moves until game is won
    while (true)
    {
        // clear the screen
        clear();

        // draw the current state of the board
        draw();

        // log the current state of the board (for testing)
        for (int i = 0; i < d; i++)
        {
            for (int j = 0; j < d; j++)
            {
                fprintf(file, "%i", board[i][j]);
                if (j < d - 1)
                {
                    fprintf(file, "|");
                }
            }
            fprintf(file, "\n");
        }
        fflush(file);

        // check for win
        if (won())
        {
            printf("ftw!\n");
            break;
        }

        // prompt for move
        printf("Tile to move: ");
        int tile = get_int();

        // quit if user inputs 0 (for testing)
        if (tile == 0)
        {
            break;
        }

        // log move (for testing)
        fprintf(file, "%i\n", tile);
        fflush(file);

        // move if possible, else report illegality
        if (!move(tile))
        {
            printf("\nIllegal move.\n");
            usleep(500000);
        }

        // sleep thread for animation's sake
        usleep(500000);
    }

    // close log
    fclose(file);

    // success
    return 0;
}

/**
 * Clears screen using ANSI escape sequences.
 */
void clear(void)
{
    printf("\033[2J");
    printf("\033[%d;%dH", 0, 0);
}

/**
 * Greets player.
 */
void greet(void)
{
    clear();
    printf("WELCOME TO GAME OF FIFTEEN\n");
    usleep(2000000);
}

/**
 * Initializes the game's board with tiles numbered 1 through d*d - 1
 * (i.e., fills 2D array with values but does not actually print them).
 *
 * It also initializes the win state array, which we'll use to facilitate
 * won comparisons, without iterating through it every time.
 */
void init(void)
{
    // First we get the number of tiles we should use so to say
    int current_number = d * d - 1;

    // We iterate through our 2D array
    for (int i = 0; i < d; i++)
    {
        for (int j = 0; j < d; j++)
        {
            // If our current number(s) haven't run out
            if (current_number > 0)
            {
                // We assign the current number to the relevant board position
                board[i][j] = current_number;

                // If this is a pair board we will do a special case override
                if (d % 2 == 0 && (current_number == 2 || current_number == 1))
                {
                    // We flip 2 and 1
                    if (current_number == 2)
                    {
                        board[i][j] = 1;
                    }
                    else
                    {
                        board[i][j] = 2;
                    }
                }

                // We reduce the current number by one
                current_number--;
            }
            else
            {
                // We use -1 to denotate an empty space
                board[i][j] = -1;
            }
        }
    }

    int incrementing_number = 1;

    /**
     * We are going to iterate again to set our initial win condition.
     * Although we could do this in the above iteration, I decided to bring
     * it to it's own loop. In the worst case scenario we added 81 iterations
     * but I think the increased readability makes it a worthwhile tradeoff.
     */
    for (int i = 0; i < d; i++)
    {
        for (int j = 0; j < d; j++)
        {
            win_state[i][j] = incrementing_number;

            if (incrementing_number == d * d)
            {
                // If we reached our maximum number, our last tile is -1 instead
                win_state[i][j] = -1;
            }

            incrementing_number++;
        }
    }
}

/**
 * Prints the board in its current state.
 */
void draw(void)
{
    for (int i = 0; i < d; i++)
    {
        for (int j = 0; j < d; j++)
        {
            if (board[i][j] != -1)
            {
                printf(" %2i", board[i][j]);
            }
            else
            {
                printf("  _");
            }
        }
        printf("\n");
    }
}

/**
 * If tile borders empty space, moves tile and returns true, else
 * returns false.
 */
bool move(int tile)
{
    // We check if the passed tile is valid at all, if not we exit early
    if (tile >= d * d)
    {
        return false;
    }
}

/**
 * Returns true if game is won (i.e., board is in winning configuration),
 * else false.
 */
bool won(void)
{
    // We won unless specified otherwise
    bool win_check = true;

    for (int i = 0; i < d; i++)
    {
        for (int j = 0; j < d; j++)
        {
            // If the current tile is different from the win state
            // we haven't won yet
            if (board[i][j] != win_state[i][j])
            {
                win_check = false;
            }
        }
    }

    return win_check;
}
