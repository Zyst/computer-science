# Questions

## What's `stdint.h`?

`stdint.h` is a header file that defines a group of variables we are going to use. In this case:

```c
typedef uint8_t  BYTE;
typedef uint32_t DWORD;
typedef int32_t  LONG;
typedef uint16_t WORD;
```

## What's the point of using `uint8_t`, `uint32_t`, `int32_t`, and `uint16_t` in a program?

These definitions allow us to be consistent in cross platform. Compiler interpretations of what an int is can be inconsistent,
for instance. But when using something like int32_t we know exactly what we are going to get.

## How many bytes is a `BYTE`, a `DWORD`, a `LONG`, and a `WORD`, respectively?

BYTE:  8-bit unsigned integer.
DWORD: 32-bit unsigned integer.
LONG:  32-bit signed integer.
WORD:  16-bit unsigned integer.

## What (in ASCII, decimal, or hexadecimal) must the first two bytes of any BMP file be? Leading bytes used to identify file formats (with high probability) are generally called "magic numbers."

`0x4d42`

## What's the difference between `bfSize` and `biSize`?

First both are DWORDs, so they are 32 unsigned bits in length.

## What does it mean if `biHeight` is negative?

If biHeight is negative, the bitmap is a top-down DIB and its origin is the upper-left corner.

## What field in `BITMAPINFOHEADER` specifies the BMP's color depth (i.e., bits per pixel)?

`bitBitCount` `The number of bits-per-pixel. The biBitCount member of the BITMAPINFOHEADER structure determines the number of bits that define each pixel and the maximum number of colors in the bitmap. This member must be one of the following values.

## Why might `fopen` return `NULL` in lines 24 and 32 of `copy.c`?

The computer or program could be out of memory, so it would be unable to open it.

## Why is the third argument to `fread` always `1` in our code?

The function takes, as its thirt argument, the count of how many such items we should read.

In this case, we want to read "1" new item, of the specified size.

## What value does line 63 of `copy.c` assign to `padding` if `bi.biWidth` is `3`?

0

## What does `fseek` do?

Sets the file position to the position we pass in

## What is `SEEK_CUR`?

It's a constant which specifies a position. In this case, the current file position
