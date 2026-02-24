package huffmancompressor

import (
	"bytes"
	"container/heap"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"
)

// Node represents a node in the Huffman tree.
type Node struct {
	R     rune
	Freq  int
	Left  *Node
	Right *Node
}

// HuffmanCompressor implements the Compressor interface.
type HuffmanCompressor struct {
	tree  *Node
	codes map[rune]string
}

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

// Type MinNodeHeap
// - stores nodes
// - is a heap
// - returns the minimum node
type MinNodeHeap []*Node

func (h MinNodeHeap) Len() int           { return len(h) }
func (h MinNodeHeap) Less(i, j int) bool { return h[i].Freq < h[j].Freq }
func (h MinNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinNodeHeap) Push(x any) {
	*h = append(*h, x.(*Node))
}

func (h *MinNodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// BuildHuffmanTree constructs a Huffman tree from a frequency table.
// It uses a min‑heap (priority queue) so that each Pop() always returns
// the node with the smallest frequency — exactly what the Huffman
// algorithm requires at each step.
func BuildHuffmanTree(freq map[rune]int) *Node {
	// Create an empty priority queue (min‑heap).
	// The heap is ordered by Node.Freq because PriorityQueue.Less() compares frequencies.
	pq := &MinNodeHeap{}
	heap.Init(pq)

	// Insert one leaf node per character into the heap.
	for r, f := range freq {
		heap.Push(pq, &Node{R: r, Freq: f})
	}

	// Huffman's algorithm:
	// Repeatedly extract the two nodes with the smallest frequency,
	// merge them into a new parent node, and push that parent back.
	for pq.Len() > 1 {
		// Pop the two smallest nodes.
		left := heap.Pop(pq).(*Node)  // smallest frequency
		right := heap.Pop(pq).(*Node) // second smallest

		// Create a new internal node whose frequency is the sum of the two children.
		parent := &Node{
			R:     0, // internal node → no rune
			Freq:  left.Freq + right.Freq,
			Left:  left,
			Right: right,
		}

		// Push the merged node back into the heap.
		// The heap will reorder itself so the smallest node is again at index 0.
		heap.Push(pq, parent)
	}

	// When only one node remains, it is the root of the Huffman tree.
	return heap.Pop(pq).(*Node)
}

func GenerateCodes(n *Node, prefix string, codes map[rune]string) {
	if n == nil {
		return
	}
	if n.Left == nil && n.Right == nil {
		codes[n.R] = prefix
		return
	}
	GenerateCodes(n.Left, prefix+"0", codes)
	GenerateCodes(n.Right, prefix+"1", codes)
}

func PrintTree(n *Node, indent string, last bool) {
	if n == nil {
		return
	}
	fmt.Print(indent)
	if last {
		fmt.Print("└─")
		indent += "  "
	} else {
		fmt.Print("├─")
		indent += "│ "
	}
	if n.R == 0 {
		fmt.Printf("(%d)\n", n.Freq)
	} else {
		fmt.Printf("%q (%d)\n", n.R, n.Freq)
	}
	PrintTree(n.Left, indent, false)
	PrintTree(n.Right, indent, true)
}

func Encode(data []byte, codes map[rune]string) string {
	var b strings.Builder
	for _, r := range string(data) {
		b.WriteString(codes[r])
	}
	return b.String()
}

func Decode(encoded string, root *Node) string {
	var out strings.Builder
	node := root

	for _, bit := range encoded {
		if bit == '0' {
			node = node.Left
		} else {
			node = node.Right
		}

		if node.Left == nil && node.Right == nil {
			out.WriteRune(node.R)
			node = root
		}
	}

	return out.String()
}

func NewHuffmanCompressor() *HuffmanCompressor {
	return &HuffmanCompressor{}
}

func (h *HuffmanCompressor) Compress(input []byte) ([]byte, error) {
	// Build frequency table and tree
	freq := BuildFrequencyTable(input)
	h.tree = BuildHuffmanTree(freq)
	h.codes = make(map[rune]string)
	GenerateCodes(h.tree, "", h.codes)

	// Encode input into bitstring
	encoded := Encode(input, h.codes)

	// Serialize tree
	buf := new(bytes.Buffer)
	serializeTree(h.tree, buf)

	// Separator (optional but useful)
	// You can remove this if you want, but it's handy for debugging.
	buf.WriteByte(0xFF)

	// Write encoded data
	buf.WriteString(encoded)

	return buf.Bytes(), nil
}

func (h *HuffmanCompressor) Decompress(input []byte) ([]byte, error) {
	buf := bytes.NewReader(input)

	// Deserialize tree
	h.tree = deserializeTree(buf)

	// Read separator (0xFF)
	sep, _ := buf.ReadByte()
	if sep != 0xFF {
		return nil, fmt.Errorf("invalid compressed file format")
	}

	// Remaining bytes = encoded bitstring
	encodedBytes, _ := io.ReadAll(buf)
	encoded := string(encodedBytes)

	// Decode using the reconstructed tree
	decoded := Decode(encoded, h.tree)
	return []byte(decoded), nil
}

func (h *HuffmanCompressor) Name() string {
	return "Huffman"
}

// Leaf marker = 1, Internal node = 0
func serializeTree(n *Node, buf *bytes.Buffer) {
	if n.Left == nil && n.Right == nil {
		buf.WriteByte(1) // leaf
		binary.Write(buf, binary.LittleEndian, int32(n.R))
		return
	}
	buf.WriteByte(0) // internal
	serializeTree(n.Left, buf)
	serializeTree(n.Right, buf)
}

func deserializeTree(buf *bytes.Reader) *Node {
	marker, _ := buf.ReadByte()
	if marker == 1 {
		var r int32
		binary.Read(buf, binary.LittleEndian, &r)
		return &Node{R: rune(r)}
	}
	left := deserializeTree(buf)
	right := deserializeTree(buf)
	return &Node{Left: left, Right: right}
}
