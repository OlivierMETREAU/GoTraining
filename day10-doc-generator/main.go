package main

// Responsibility:
// - Parse flags (input folder, output folder, theme, server mode)
// - Call the generator
// - Optionally start an HTTP server

import (
	"fmt"
	"log"
	"os"

	"example.com/day10-doc-generator/htmldocgenerator"
)

func main() {
	fmt.Println("day10-doc-generator")

	if len(os.Args) < 2 {
		log.Fatal("Usage: mydocgen <input-folder>")
	}

	input := os.Args[1]

	// Call your extraction pipeline with 4 workers
	results, err := htmldocgenerator.ExtractProjectInfo(input, 4)
	if err != nil {
		log.Fatalf("Error extracting project info: %v", err)
	}

	if err := htmldocgenerator.GenerateHTML("docs.html", results, input); err != nil {
		log.Fatalf("HTML generation failed: %v", err)
	}
}
