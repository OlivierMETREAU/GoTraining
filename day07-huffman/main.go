package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"example.com/day07-huffman/huffmancompressor"
)

func main() {
	fmt.Println("day07-huffman")

	// Define CLI flags
	mode := flag.String("mode", "", "Operation mode: compress or decompress")
	input := flag.String("in", "", "Input file path")
	output := flag.String("out", "", "Output file path")

	flag.Parse()

	// Validate parameters
	if *mode == "" || *input == "" || *output == "" {
		fmt.Println("Usage:")
		fmt.Println("  huff -mode compress   -in input.txt   -out output.huff")
		fmt.Println("  huff -mode decompress -in output.huff -out decoded.txt")
		os.Exit(1)
	}

	// Read input file
	data, err := os.ReadFile(*input)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	compressor := huffmancompressor.NewHuffmanCompressor()

	var result []byte

	switch *mode {
	case "compress":
		result, err = compressor.Compress(data)
	case "decompress":
		result, err = compressor.Decompress(data)
	default:
		err = errors.New("invalid mode: must be 'compress' or 'decompress'")
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Write output file
	if err := os.WriteFile(*output, result, 0644); err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Success: %s → %s (%s)\n", *input, *output, *mode)
}

// examples :
// > go run main.go -mode compress -in .\testResources\originalTextFile.txt -out .\testResources\compressed.huff
// > go run main.go -mode decompress -in .\testResources\compressed.huff -out .\testResources\decompressed.txt
