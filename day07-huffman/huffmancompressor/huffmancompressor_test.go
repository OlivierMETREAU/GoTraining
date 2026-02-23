package huffmancompressor

import (
	"bytes"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..", "testResources")
	path := filepath.Join(dir, "originalTextFile.txt")

	data, err := ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile returned error: %v", err)
	}

	assert.True(t, bytes.HasPrefix(data, []byte("Le codage de Huffman")))
}

func TestBuildFrequencyTable(t *testing.T) {
	data := []byte("aabccc")
	freq := BuildFrequencyTable(data)

	expected := map[rune]int{'a': 2, 'b': 1, 'c': 3}
	assert.Equal(t, expected, freq)
}

func TestBuildHuffmanTree(t *testing.T) {
	freq := map[rune]int{'a': 5, 'b': 2, 'c': 1}
	tree := BuildHuffmanTree(freq)
	assert.Equal(t, 8, tree.Freq, "expected root freq 8, got %d", tree.Freq)
}

func TestGenerateCodes(t *testing.T) {
	freq := map[rune]int{'a': 5, 'b': 2, 'c': 1}
	tree := BuildHuffmanTree(freq)

	codes := make(map[rune]string)
	GenerateCodes(tree, "", codes)

	assert.Equal(t, 3, len(codes), "expected 3 codes, got %d", len(codes))
	assert.Equal(t, map[rune]string{'a': "1", 'b': "01", 'c': "00"}, codes)
}
