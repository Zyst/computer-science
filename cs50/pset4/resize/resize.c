// Copies a BMP file

#include <stdio.h>
#include <stdlib.h>

#include "bmp.h"

int main(int argc, char *argv[])
{
    // ensure proper usage
    if (argc != 4)
    {
        fprintf(stderr, "Usage: resize modifier_size infile outfile\n");
        return 1;
    }

    // Get the resize modifier we are doing
    float resize = strtof(argv[1], NULL);

    printf("%f\n", resize);

    // remember filenames
    char *infile = argv[2];
    char *outfile = argv[3];

    // open input file
    FILE *inptr = fopen(infile, "r");
    if (inptr == NULL)
    {
        fprintf(stderr, "Could not open %s.\n", infile);
        return 2;
    }

    // open output file
    FILE *outptr = fopen(outfile, "w");
    if (outptr == NULL)
    {
        fclose(inptr);
        fprintf(stderr, "Could not create %s.\n", outfile);
        return 3;
    }

    // read infile's BITMAPFILEHEADER
    BITMAPFILEHEADER bf;
    fread(&bf, sizeof(BITMAPFILEHEADER), 1, inptr);

    // read infile's BITMAPINFOHEADER
    BITMAPINFOHEADER bi;
    fread(&bi, sizeof(BITMAPINFOHEADER), 1, inptr);

    // ensure infile is (likely) a 24-bit uncompressed BMP 4.0
    if (bf.bfType != 0x4d42 || bf.bfOffBits != 54 || bi.biSize != 40 ||
        bi.biBitCount != 24 || bi.biCompression != 0)
    {
        fclose(outptr);
        fclose(inptr);
        fprintf(stderr, "Unsupported file format.\n");
        return 4;
    }

    int originalWidth = bi.biWidth;
    int originalHeight = bi.biHeight;
    int originalPadding = (4 - (bi.biWidth * sizeof(RGBTRIPLE)) % 4) % 4;

    // We edit the size of the file headers!
    bi.biWidth = bi.biWidth * resize;
    bi.biHeight = bi.biHeight * resize;

    // determine padding for scanlines
    int padding = (4 - (bi.biWidth * sizeof(RGBTRIPLE)) % 4) % 4;

    bi.biSizeImage = ((sizeof(RGBTRIPLE) * bi.biWidth) + padding) * abs(bi.biHeight);

    bf.bfSize = bi.biSizeImage + sizeof(BITMAPFILEHEADER) + sizeof(BITMAPINFOHEADER);

    // write outfile's BITMAPFILEHEADER
    fwrite(&bf, sizeof(BITMAPFILEHEADER), 1, outptr);

    // write outfile's BITMAPINFOHEADER
    fwrite(&bi, sizeof(BITMAPINFOHEADER), 1, outptr);

    RGBTRIPLE *pixelRow = malloc(sizeof(RGBTRIPLE) * originalWidth);

    // iterate over infile's scanlines
    for (int i = 0, biHeight = abs(originalHeight); i < biHeight; i++)
    {
        // iterate over pixels in scanline
        for (int j = 0; j < originalWidth; j++)
        {
            // temporary storage
            RGBTRIPLE triple;

            // read RGB triple from infile
            fread(&triple, sizeof(RGBTRIPLE), 1, inptr);


            pixelRow[j] = triple;
        }

        for (int z = 0; z < resize; z++)
        {
            for (int x = 0; x < originalWidth; x++)
            {
                for (int y = 0; y < resize; y++)
                {
                    // write RGB triple to outfile
                    fwrite(&pixelRow[x], sizeof(RGBTRIPLE), 1, outptr);
                }
            }

            // then add the new padding
            for (int k = 0; k < padding; k++)
            {
                fputc(0x00, outptr);
            }
        }

        // skip over padding from the old file, if any
        fseek(inptr, originalPadding, SEEK_CUR);
    }

    // close infile
    fclose(inptr);

    // close outfile
    fclose(outptr);

    free(pixelRow);

    // success
    return 0;
}
