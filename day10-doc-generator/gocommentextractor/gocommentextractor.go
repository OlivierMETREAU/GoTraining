package gocommentextractor

// Package gocommentextractor
//
// Responsibilities:
//   - Parse .go source files using the Go AST
//   - Extract documentation comments (package, type, function, const, var, import)
//   - Return a structured representation suitable for documentation generation

import (
	"go/ast"
	"go/parser"
	"go/token"
)

//
// ============================================================
//  Data structures
// ============================================================
//

// CommentBlock represents a single extracted comment block.
// It includes the text, the line range, the semantic context,
// and an optional sub-context (e.g., function name).
type CommentBlock struct {
	Text       string // Raw comment text
	LineStart  int    // First line of the comment block
	LineEnd    int    // Last line of the comment block
	Context    string // "package", "type", "function", "var", "const", "import", ...
	SubContext string // Additional detail, e.g. function name ("Open")
}

// FileComments represents all extracted comments for a single Go file.
type FileComments struct {
	FilePath string         // Absolute or relative path to the file
	Package  string         // Package name declared in the file
	Comments []CommentBlock // All extracted comment blocks
	Err      error          // Optional error encountered during parsing
}

//
// ============================================================
//  Public API
// ============================================================
//

// GetCommentFromGoFile parses a Go source file and extracts all
// documentation comments (package-level and declaration-level).
func GetCommentFromGoFile(path string) (FileComments, error) {
	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return FileComments{}, err
	}

	fc := FileComments{
		FilePath: path,
		Package:  file.Name.Name,
		Comments: make([]CommentBlock, 0),
	}

	extractPackageComments(file, fset, &fc)
	extractDeclarationComments(file, fset, &fc)

	return fc, nil
}

//
// ============================================================
//  Extraction helpers
// ============================================================
//

// extractPackageComments extracts the comment block immediately above
// the `package xyz` declaration, if present.
func extractPackageComments(file *ast.File, fset *token.FileSet, fc *FileComments) {
	if file.Doc == nil {
		return
	}

	start := fset.Position(file.Doc.Pos()).Line
	end := fset.Position(file.Doc.End()).Line

	fc.Comments = append(fc.Comments, CommentBlock{
		Text:      file.Doc.Text(),
		LineStart: start,
		LineEnd:   end,
		Context:   "package",
	})
}

// extractDeclarationComments extracts comments attached to:
//   - type declarations
//   - const blocks
//   - var blocks
//   - import blocks
//   - function declarations
//
// It walks the AST and collects all top-level documentation comments.
func extractDeclarationComments(file *ast.File, fset *token.FileSet, fc *FileComments) {
	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {

		// General declarations: type, const, var, import
		case *ast.GenDecl:
			if node.Doc == nil {
				return true
			}

			ctx := ""
			switch node.Tok {
			case token.TYPE:
				ctx = "type"
			case token.CONST:
				ctx = "const"
			case token.VAR:
				ctx = "var"
			case token.IMPORT:
				ctx = "import"
			}

			start := fset.Position(node.Doc.Pos()).Line
			end := fset.Position(node.Doc.End()).Line

			fc.Comments = append(fc.Comments, CommentBlock{
				Text:      node.Doc.Text(),
				LineStart: start,
				LineEnd:   end,
				Context:   ctx,
			})

		// Function declarations
		case *ast.FuncDecl:
			if node.Doc == nil {
				return true
			}

			start := fset.Position(node.Doc.Pos()).Line
			end := fset.Position(node.Doc.End()).Line

			fc.Comments = append(fc.Comments, CommentBlock{
				Text:       node.Doc.Text(),
				LineStart:  start,
				LineEnd:    end,
				Context:    "function",
				SubContext: node.Name.Name,
			})
		}

		return true
	})
}
