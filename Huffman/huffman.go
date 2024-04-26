package huffman

import (
	"container/heap"
	"fmt"
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
	leftChild, rightChild HuffmanTree
}

// FUNCTIONS TO RETURN THE FREQUENCY OF EACH NODE

func (self HuffmanLeaf) Freq() int{
	return self.freq
}

func (self HuffmanNode) Freq() int{
	return self.freq
}


// CREATING A HEAP TO TRACK THE MINIMUM ELEMENTS
type treeHeap []HuffmanTree

// IMPLEMENTING THE HEAP FUNCTIONS
func (h treeHeap)Len() int{
	return len(h)
}

func (h treeHeap)Swap(i,j int){
	h[i],h[j] = h[j],h[i]
}

func (h treeHeap)Less(i,j int) bool {
	return h[i].Freq() < h[j].Freq()
}

// PUSH AN ELEMENT INTO THE HEAP
func (h *treeHeap)Push(element interface{}) {
	*h = append(*h, element.(HuffmanTree))
}

// POP RETURNS THE MINIMUM ELEMENT FROM THE HEAP
func (h *treeHeap)Pop()(element interface{}){
	element = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return 
}

func BuildTree(charMap map[rune]int) HuffmanTree{
	
	var trees treeHeap
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

func PrintCodes(tree HuffmanTree, prefix []byte){
	switch i := tree.(type){
		// LEAF ONLY CONTAINS THE VALUE AND CHARACTER, SO WE PRINT THEM OUT
		// WITH THE PREFIX CODE
	case HuffmanLeaf:
		fmt.Printf("%d\t%d\t%s\n",i.freq,i.character,string(prefix))

	case HuffmanNode:
		// ASSIGN '0' WHILE TRAVERSING THE LEFT SUB TREE
		prefix = append(prefix, '0')
		PrintCodes(i.leftChild, prefix) // CALL THE PRINT FUNCTION ON THE LEFT SUB TREE
		prefix = prefix[:len(prefix)-1] // REMOVE THE LAST EXTRA BYTE

		// ASSIGN '1' WHILE TRAVERSING THE LEFT SUB TREE
		prefix = append(prefix, '1')
		PrintCodes(i.rightChild, prefix) // CALL THE PRINT FUNCTION ON THE LEFT SUB TREE
		prefix = prefix[:len(prefix)-1] // REMOVE THE LAST EXTRA BYTE
	}
}