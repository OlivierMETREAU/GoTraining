package htmldocgenerator

import (
	"io/fs"
	"path/filepath"
	"strings"

	"example.com/day10-doc-generator/gocommentextractor"
)

// Responsibility:
// - Recursively walk the project
// - Spawn goroutines to extract comments using the previous package
// - Generate HTML pages (per file, per package, or a global index)

// FindGoFiles recursively scans a directory and returns all .go files
// except *_test.go files.
func FindGoFiles(root string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Only .go files
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		// Exclude test files
		if strings.HasSuffix(path, "_test.go") {
			return nil
		}

		// This is a valid go source code file to be analyzed
		files = append(files, path)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

// ExtractProjectInfo scans a folder, finds all Go files, and extracts
// package names using a worker pool calling gocommentextractor.
func ExtractProjectInfo(root string, workerCount int) ([]gocommentextractor.FileComments, error) {
	files, err := FindGoFiles(root)
	if err != nil {
		return nil, err
	}

	return runWorkers(files, workerCount)
}

func runWorkers(files []string, workerCount int) ([]gocommentextractor.FileComments, error) {
	jobs := make(chan string)
	results := make(chan gocommentextractor.FileComments)

	// Start workers
	for i := 0; i < workerCount; i++ {
		go func() {
			for path := range jobs {
				fc, err := gocommentextractor.GetCommentFromGoFile(path)
				if err != nil {
					// Return an error inside the struct so the pipeline keeps flowing
					results <- gocommentextractor.FileComments{
						FilePath: path,
						Err:      err,
					}
					continue
				}
				results <- fc
			}
		}()
	}

	// Feed jobs
	go func() {
		for _, f := range files {
			jobs <- f
		}
		close(jobs)
	}()

	// Collect results
	output := make([]gocommentextractor.FileComments, 0, len(files))
	for range files {
		output = append(output, <-results)
	}

	return output, nil
}
