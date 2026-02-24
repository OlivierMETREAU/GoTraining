package main

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"

	"example.com/day07-huffman/huffmancompressor"
)

func main() {
	fmt.Println("day07-huffman")

	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "testResources")
	path := filepath.Join(dir, "originalTextFile.txt")

	data, err := huffmancompressor.ReadFile(path)
	if err == nil {
		freq := huffmancompressor.BuildFrequencyTable(data)
		tree := huffmancompressor.BuildHuffmanTree(freq)
		//huffmancompressor.PrintTree(tree, "   ", false)
		codes := make(map[rune]string)
		huffmancompressor.GenerateCodes(tree, "", codes)
	}
}
