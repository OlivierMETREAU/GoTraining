package htmldocgenerator

// Responsibility:
// - Recursively walk the project
// - Spawn goroutines to extract comments using the previous package
// - Generate HTML pages (per file, per package, or a global index)
