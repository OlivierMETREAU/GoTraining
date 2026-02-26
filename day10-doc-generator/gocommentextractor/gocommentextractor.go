package gocommentextractor

// Responsibility:
// - Parse .go files
// - Extract comments (package comments, type comments, func comments, inline comments)
// - Return a structured representation

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// CommentBlock structure containing the linStart, lineEnd, text and context of each comment block
type CommentBlock struct {
	LineStart int
	LineEnd   int
	Text      string
	Context   string // "package", "type", "func", "var", "inline"
}

// FileComments strcuture containing the filePth, the package name and the slice of comments
type FileComments struct {
	FilePath string
	Package  string
	Comments []CommentBlock
	Err      error
}

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

func extractDeclarationComments(file *ast.File, fset *token.FileSet, fc *FileComments) {
	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {

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

		case *ast.FuncDecl:
			if node.Doc == nil {
				return true
			}

			ctx := "func " + node.Name.Name

			start := fset.Position(node.Doc.Pos()).Line
			end := fset.Position(node.Doc.End()).Line

			fc.Comments = append(fc.Comments, CommentBlock{
				Text:      node.Doc.Text(),
				LineStart: start,
				LineEnd:   end,
				Context:   ctx,
			})
		}

		return true
	})
}
