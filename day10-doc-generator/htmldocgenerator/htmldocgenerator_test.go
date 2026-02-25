package htmldocgenerator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindGoFiles_UsingLocalBoltRepo(t *testing.T) {
	// Resolve the path to ../../bolt relative to this test file
	baseDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("cannot get working directory: %v", err)
	}

	boltPath := filepath.Join(baseDir, "..", "..", "..", "bolt")

	// Ensure the directory exists
	stat, err := os.Stat(boltPath)
	if err != nil {
		t.Fatalf("bolt directory not found at %s: %v", boltPath, err)
	}
	if !stat.IsDir() {
		t.Fatalf("expected a directory at %s", boltPath)
	}

	// Run the function
	files, err := FindGoFiles(boltPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(files) == 0 {
		t.Fatalf("expected to find Go files in %s, found none", boltPath)
	}

	// Basic sanity checks
	for _, f := range files {
		if filepath.Ext(f) != ".go" {
			t.Fatalf("non-go file returned: %s", f)
		}
		if filepath.Base(f) == "_test.go" {
			t.Fatalf("test file should not be included: %s", f)
		}
	}
}
