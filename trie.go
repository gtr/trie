package trie

import "errors"

// Trie represents a trie object.
type Trie struct {
	Root *Node
}

// NewTrie returns a pointer to an empty trie.
func NewTrie() *Trie {
	return &Trie{
		Root: NewNode(),
	}
}

// InsertWord inserts a new word into the trie.
func (t *Trie) InsertWord(word string) {
	runner := t.Root

	for _, currChar := range word {
		if _, ok := runner.Children[currChar]; !ok {
			runner.Children[currChar] = NewNode()
		}
		runner = runner.Children[currChar]
	}

	runner.IsWord = true
}

// InsertWords inserts multiple words into the trie.
func (t *Trie) InsertWords(words []string) {
	for _, word := range words {
		t.InsertWord(word)
	}
}

// FindWord returns a bool if a word exists in the trie.
func (t *Trie) FindWord(word string) bool {
	return t.existsPrefix(word, true)
}

// GetAllWords returns a slice of strings containing all the words in the
// entire trie.
func (t *Trie) GetAllWords() []string {
	runner := t.Root
	return runner.GetAllSubWords("")
}

// existsPrefix checks if a prefix exists in the trie. If validWord is set to
// true, then it also checks if that prefix is a valid word.
func (t *Trie) existsPrefix(prefix string, validWord bool) bool {
	runner := t.Root

	for _, currChar := range prefix {
		if _, ok := runner.Children[currChar]; !ok {
			return false
		}
		runner = runner.Children[currChar]
	}

	if validWord && !runner.IsWord {
		return false
	}
	return true
}

// getNodeAtPrefix returns a pointer to the last node satisfying the given
// prefix. Returns error if the prefix does not exist.
func (t *Trie) getNodeAtPrefix(prefix string) (*Node, error) {
	runner := t.Root

	for _, currChar := range prefix {
		if _, ok := runner.Children[currChar]; !ok {
			return nil, errors.New("prefix not found")
		}
		runner = runner.Children[currChar]
	}

	return runner, nil
}

// AutoComplete returns a slice of strings containing all the possible words
// that can autocomplete the given prefix.
func (t *Trie) AutoComplete(prefix string) ([]string, error) {
	if t.existsPrefix(prefix, false) {
		node, err := t.getNodeAtPrefix(prefix)
		if err != nil {
			return nil, err
		}
		return node.GetAllSubWords(prefix), nil
	}

	return nil, errors.New("prefix not found")
}
