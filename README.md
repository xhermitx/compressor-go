# Huffman Compression Tool Documentation

## Overview

The Huffman Compression Tool is a command-line utility written in Go that enables users to compress and decompress text files using the Huffman coding algorithm. This algorithm is a popular method of lossless data compression that uses variable-length codes to represent characters based on their frequencies.

## Installation

To install the Huffman Compression Tool, you must have Go installed on your machine. Follow these steps:

1. Download or clone the tool's repository to your local machine.
2. Navigate to the directory where the tool's source code is located.
3. Compile the code using the Go compiler:

`go build -o main`

This command will create an executable file named `main`.

## Usage

The Huffman Compression Tool accepts command-line arguments to specify the operation mode (compression or decompression), the input filename, and the output filename.

### Compressing a File

To compress a file, use the `-e` (encode) flag followed by the input filename and the desired output filename for the compressed file.

`./main -e "inputfilename.txt" "outputfilename.huf"`

Here, `"inputfilename.txt"` is the name of the text file you want to compress, and `"outputfilename.huf"` is the name of the file where the compressed data will be stored.

### Decompressing a File

To decompress a previously compressed file, use the `-o` (decode) flag followed by the input filename (the compressed file) and the desired output filename for the decompressed text.

`./main -o "inputfilename.huf" "outputfilename.txt"`

Here, `"inputfilename.huf"` is the name of the compressed file you want to decompress, and `"outputfilename.txt"` is the name of the file where the decompressed text will be written.

## Notes

- Ensure that the input file exists and is readable before attempting to compress or decompress it.
- When compressing a file, the tool also generates a header containing the Huffman tree used for encoding. This header is crucial for decompression and should not be modified.
- The output file will be overwritten if it already exists.
- The tool currently supports only text files and may not work correctly with binary files or files containing non-text data.
- To achieve effective compression, the input text file should be sufficiently large and contain some redundancy in characters.

## Examples

Compressing `document.txt` and saving the compressed data to `document.huf`:

`./main -e "document.txt" "document.huf"`

Decompressing `document.huf` and restoring the original text to `document.txt`:

`./main -o "document.huf" "document.txt"`

## INSPIRATION

This tool is inspired from [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-huffman).
