package main

// Responsibility:
// - Parse flags (input folder, output folder, theme, server mode)
// - Call the generator
// - Optionally start an HTTP server

import (
	"fmt"
	"log"
	"os"

	"example.com/day10-doc-generator/gocommentextractor"
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

	// Print results for now (later you’ll generate HTML)
	// for _, r := range results {
	// 	printFileSummary(r)
	// }
}

func printFileSummary(fc gocommentextractor.FileComments) {
	if fc.Err != nil {
		fmt.Printf("Error in file %s: %v\n", fc.FilePath, fc.Err)
		return
	}

	fmt.Printf("\n=== %s ===\n", fc.FilePath)
	fmt.Printf("Package: %s\n", fc.Package)

	grouped := groupCommentsByContext(fc.Comments)

	for ctx, list := range grouped {
		fmt.Printf("\n  [%s]\n", ctx)

		limit := 3
		if len(list) < limit {
			limit = len(list)
		}

		for i := 0; i < limit; i++ {
			fmt.Printf("    - %q\n", trimComment(list[i]))
		}

		if len(list) > limit {
			fmt.Println("    - ...")
		}
	}
}
func groupCommentsByContext(comments []gocommentextractor.CommentBlock) map[string][]string {
	grouped := make(map[string][]string)
	for _, c := range comments {
		grouped[c.Context] = append(grouped[c.Context], c.Text)
	}
	return grouped
}
func trimComment(s string) string {
	const max = 80
	if len(s) <= max {
		return s
	}
	return s[:max] + "..."
}
