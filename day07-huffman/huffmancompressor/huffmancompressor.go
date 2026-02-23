package huffmancompressor

import (
	"container/heap"
	"os"
)

type Node struct {
	R     rune
	Freq  int
	Left  *Node
	Right *Node
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
