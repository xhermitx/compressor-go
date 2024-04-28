package huffman

import (
	"container/heap"
)

// INTERFACE FOR THE TREE
type HuffmanTree interface{
	Freq() int
}

// LEAF NODE WITH CHARACTER-FREQ MAP
type HuffmanLeaf struct{
	character rune
	freq int
}

// INTERNAL NODES WITH ONLY THE SUM OF CHILD FREQ
type HuffmanNode struct{
	freq int
	left_child, right_child HuffmanTree
}

// FUNCTIONS TO RETURN THE FREQUENCY OF EACH NODE

func (self HuffmanLeaf) Freq() int{
	return self.freq
}

func (self HuffmanNode) Freq() int{
	return self.freq
}


// CREATING A HEAP TO TRACK THE MINIMUM ELEMENTS
type TreeHeap []HuffmanTree

// IMPLEMENTING THE HEAP FUNCTIONS
func (h TreeHeap)Len() int{
	return len(h)
}

func (h TreeHeap)Swap(i,j int){
	h[i],h[j] = h[j],h[i]
}

func (h TreeHeap)Less(i,j int) bool {
	return h[i].Freq() < h[j].Freq()
}

// PUSH AN ELEMENT INTO THE HEAP
func (h *TreeHeap)Push(element interface{}) {
	*h = append(*h, element.(HuffmanTree))
}

// POP RETURNS THE MINIMUM ELEMENT FROM THE HEAP
func (h *TreeHeap)Pop()(element interface{}){
	element = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return 
}

func BuildTree(charMap map[rune]int) HuffmanTree{
	
	var trees TreeHeap
	for char,freq := range charMap{
		trees = append(trees, HuffmanLeaf{char,freq})	
	}

	heap.Init(&trees)

	for trees.Len()>1{

		// EXTRACT THE FREQ OF 2 MINIMUM NODES/TREES
		tree1 := heap.Pop(&trees).(HuffmanTree)
		tree2 := heap.Pop(&trees).(HuffmanTree)

		heap.Push(&trees, HuffmanNode{tree1.Freq() + tree2.Freq(), tree1, tree2})
	}

	return heap.Pop(&trees).(HuffmanTree)
}

func GenerateCodes(tree HuffmanTree, prefix []byte, encoder map[rune]string) map[rune]string{

	switch i := tree.(type){
		// LEAF ONLY CONTAINS THE VALUE AND CHARACTER, SO WE RETURN THEM
		// WITH THE PREFIX CODE
	case HuffmanLeaf:
		encoder[i.character] = string(prefix)

	case HuffmanNode:
		// ASSIGN '0' WHILE TRAVERSING THE LEFT SUB TREE
		prefix = append(prefix, '0')
		GenerateCodes(i.left_child, prefix, encoder) // RECURSIVE CALL ON THE LEFT SUB TREE
		prefix = prefix[:len(prefix)-1] // REMOVE THE LAST EXTRA BYTE

		// ASSIGN '1' WHILE TRAVERSING THE LEFT SUB TREE
		prefix = append(prefix, '1')
		GenerateCodes(i.right_child, prefix, encoder) // RECURSIVE CALL ON THE RIGHT SUB TREE
		prefix = prefix[:len(prefix)-1] // REMOVE THE LAST EXTRA BYTE
	}
	return encoder
}