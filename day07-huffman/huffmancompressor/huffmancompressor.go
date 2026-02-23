package huffmancompressor

import "os"

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func BuildFrequencyTable(data []byte) map[rune]int {
	freq := make(map[rune]int)
	for _, r := range string(data) {
		freq[r]++
	}
	return freq
}
