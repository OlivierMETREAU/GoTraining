package htmldocgenerator

import (
	"io/fs"
	"path/filepath"
	"strings"
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
