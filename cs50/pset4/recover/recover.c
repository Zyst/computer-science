// Recovers JPEG files from a file

#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
    // ensure proper usage
    if (argc != 2)
    {
        fprintf(stderr, "Usage: ./recover image\n");
        return 1;
    }

    // remember filenames
    char *infile = argv[1];

    // open input file
    FILE *inptr = fopen(infile, "r");
    if (inptr == NULL)
    {
        fprintf(stderr, "Could not open %s.\n", infile);
        return 2;
    }

    // The buffer we will use to read 512 bytes at a time
    unsigned char buffer[512];
    // Incremental variable we use for variable names
    int currentImage = 0;
    // Our format is 000.jpg, 002.jpg, 00n.jpg, so we need 8 items
    char fileName[8];
    // Have we ever found a JPEG? 0 if not, 1 if yes.
    int previouslyFoundJpeg = 0;
    // Our file
    FILE *outptr = NULL;

    while (1)
    {
        int readSize = fread(buffer, 1, 512, inptr);

        if (readSize != 512)
        {
            break;
        }

        // If the current buffer is the start of a JPEG file
        if (buffer[0] == 0xff &&
            buffer[1] == 0xd8 &&
            buffer[2] == 0xff &&
            (buffer[3] & 0xf0) == 0xe0)
        {
            if (previouslyFoundJpeg)
            {
                // If we have previously found a JPEG we increment
                currentImage = currentImage + 1;

                // Close the previous file
                fclose(outptr);
            }
            else
            {
                // We've found our first JPEG!
                previouslyFoundJpeg = 1;
            }

            // Create formatted filename, ie: 000.jpg, 001.jpg
            sprintf(fileName, "%03i.jpg", currentImage);

            // open new output file
            outptr = fopen(fileName, "w");

            if (outptr == NULL)
            {
                fclose(inptr);
                fprintf(stderr, "Could not create %s.\n", fileName);
                return 3;
            }
        }

        if (previouslyFoundJpeg)
        {
            fwrite(buffer, 1, 512, outptr);
        }
    }


    // close infile
    fclose(inptr);

    // success
    return 0;
}
