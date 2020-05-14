package trie

// Node represents a Node in the trie.
type Node struct {
	IsWord   bool
	Children map[rune]*Node
}

// NewNode returns a pointer to an empty node.
func NewNode() *Node {
	return &Node{
		IsWord:   false,
		Children: make(map[rune]*Node),
	}
}

// GetAllSubWords returns a slice of strings containing all the words in the
// subtrie contained in the current node n.
func (n *Node) GetAllSubWords(curr string) []string {
	var listOfWords []string

	for char, child := range n.Children {
		result := child.GetAllSubWords(curr + string(char))
		listOfWords = append(listOfWords, result...)
	}

	if n.IsWord {
		listOfWords = append(listOfWords, curr)
	}

	return listOfWords
}
