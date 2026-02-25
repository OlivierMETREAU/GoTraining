package gocommentextractor

// Responsibility:
// - Parse .go files
// - Extract comments (package comments, type comments, func comments, inline comments)
// - Return a structured representation

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
}
