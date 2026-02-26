package htmldocgenerator

import (
	"encoding/json"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"example.com/day10-doc-generator/gocommentextractor"
)

//
// ============================================================
//  Package htmldocgenerator
//  ------------------------
//  Responsibilities:
//  - Recursively scan a Go project
//  - Extract comments using gocommentextractor (via worker pool)
//  - Prepare structured documentation data
//  - Generate HTML output or serve it through an HTTP server
// ============================================================
//

//
// =========================
//  Internal types
// =========================
//

// htmlFile represents a Go source file prepared for HTML rendering.
type htmlFile struct {
	ID          string
	RelPath     string
	Package     string
	GroupedDocs map[string][]gocommentextractor.CommentBlock
}

// htmlData is the top-level structure passed to the HTML template.
type htmlData struct {
	Files []htmlFile
}

// GroupedComments is a convenience alias for readability.
type GroupedComments map[string][]gocommentextractor.CommentBlock

//
// =========================
//  Go file discovery
// =========================
//

// FindGoFiles recursively scans a directory and returns all .go files,
// excluding *_test.go files.
func FindGoFiles(root string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		if strings.HasSuffix(path, "_test.go") {
			return nil
		}

		files = append(files, path)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

//
// =========================
//  Comment extraction (worker pool)
// =========================
//

// ExtractProjectInfo scans a folder, finds all Go files, and extracts
// comments concurrently using a worker pool.
func ExtractProjectInfo(root string, workerCount int) ([]gocommentextractor.FileComments, error) {
	files, err := FindGoFiles(root)
	if err != nil {
		return nil, err
	}
	return runWorkers(files, workerCount)
}

// runWorkers processes files concurrently and returns extracted comments.
func runWorkers(files []string, workerCount int) ([]gocommentextractor.FileComments, error) {
	jobs := make(chan string)
	results := make(chan gocommentextractor.FileComments)

	// Start worker goroutines
	for i := 0; i < workerCount; i++ {
		go func() {
			for path := range jobs {
				fc, err := gocommentextractor.GetCommentFromGoFile(path)
				if err != nil {
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

//
// =========================
//  HTML data preparation
// =========================
//

// groupByContext groups extracted comments by their semantic context
// (package, type, function, var, const, etc.).
func groupByContext(comments []gocommentextractor.CommentBlock) map[string][]gocommentextractor.CommentBlock {
	grouped := make(map[string][]gocommentextractor.CommentBlock)
	for _, c := range comments {
		grouped[c.Context] = append(grouped[c.Context], c)
	}
	return grouped
}

// prepareHTMLData transforms raw FileComments into the structure
// expected by the HTML template.
func prepareHTMLData(files []gocommentextractor.FileComments, root string) htmlData {
	data := htmlData{Files: make([]htmlFile, 0, len(files))}

	for _, fc := range files {
		rel, _ := filepath.Rel(root, fc.FilePath)
		id := "file_" + strings.ReplaceAll(rel, string(os.PathSeparator), "_")

		grouped := groupByContext(fc.Comments)

		data.Files = append(data.Files, htmlFile{
			ID:          id,
			RelPath:     rel,
			Package:     fc.Package,
			GroupedDocs: grouped,
		})
	}

	return data
}

//
// =========================
//  HTML generation (file mode)
// =========================
//

// GenerateHTML writes a complete HTML documentation file to disk.
func GenerateHTML(outputPath string, files []gocommentextractor.FileComments, root string) error {
	data := prepareHTMLData(files, root)

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl := template.Must(template.New("doc").Parse(htmlTemplate))
	return tmpl.Execute(f, data)
}

//
// =========================
//  HTTP server mode
// =========================
//

// ServeHTML renders the documentation HTML directly to an HTTP response.
func ServeHTML(w http.ResponseWriter, files []gocommentextractor.FileComments, root string) {
	data := prepareHTMLData(files, root)
	tmpl := template.Must(template.New("doc").Parse(htmlTemplate))
	_ = tmpl.Execute(w, data)
}

// ServeJSON exposes the raw extracted data as JSON.
// Useful for SPA frontends or debugging.
func ServeJSON(w http.ResponseWriter, files []gocommentextractor.FileComments) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(files)
}
