package main

// Responsibility:
// - Parse flags (input folder, output folder, theme, server mode)
// - Call the generator
// - Optionally start an HTTP server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"example.com/day10-doc-generator/htmldocgenerator"
)

func main() {
	fmt.Println("day10-doc-generator")

	if len(os.Args) < 2 {
		log.Fatal("Usage: mydocgen <folder> OR mydocgen serve <folder>")
	}

	if os.Args[1] == "serve" {
		if len(os.Args) < 3 {
			log.Fatal("Usage: mydocgen serve <folder>")
		}
		runServer(os.Args[2])
		return
	}

	// Default: generate HTML + open browser
	input := os.Args[1]

	// Call your extraction pipeline with 4 workers
	results, err := htmldocgenerator.ExtractProjectInfo(input, 4)
	if err != nil {
		log.Fatalf("Error extracting project info: %v", err)
	}

	output := "docs.html"
	if err := htmldocgenerator.GenerateHTML(output, results, input); err != nil {
		log.Fatalf("HTML generation failed: %v", err)
	}

	fmt.Printf("Documentation generated: %s\n", output)

	// Open the HTML file automatically
	if err := openBrowser(output); err != nil {
		fmt.Printf("Could not open browser automatically: %v\n", err)
	}
}

func runServer(root string) {
	fmt.Println("Extracting project info…")

	results, err := htmldocgenerator.ExtractProjectInfo(root, 4)
	if err != nil {
		log.Fatalf("Extraction failed: %v", err)
	}

	fmt.Println("Starting server on http://localhost:8080")

	// Serve the HTML UI
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		htmldocgenerator.ServeHTML(w, results, root)
	})

	// Serve JSON API
	http.HandleFunc("/api/files", func(w http.ResponseWriter, r *http.Request) {
		htmldocgenerator.ServeJSON(w, results)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func openBrowser(path string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", path).Start()
	case "windows":
		return exec.Command("cmd", "/c", "start", path).Start()
	case "darwin":
		return exec.Command("open", path).Start()
	default:
		return fmt.Errorf("unsupported platform")
	}
}
