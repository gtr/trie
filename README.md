a quick and effective trie library implementation for go

## install

use the `go get` command:

```bash
go get github.com/gtr/trie
```

## usage

import the library into your go main file:
```go
import "github.com/gtr/trie"
```

example usage:
```go
package main

import (
	"fmt"
	"log"

	"github.com/gtr/trie"
)

func main() {
	t := trie.NewTrie()
	t.InsertWord("hello")
	t.InsertWords([]string{
		"hi",
		"hoop",
		"hook",
		"breakfast",
		"brunch",
		"brush",
		"bank",
	})

	auto, err := t.AutoComplete("br")
	if err != nil {
		log.Fatalf("AutoComplete: %s", err)
	}

	for _, word := range auto {
		fmt.Println(word)
	}

	all := t.GetAllWords()
	fmt.Println("----")
	for _, word := range all {
		fmt.Println(word)
	}
}

```

output:
```
breakfast
brush
brunch
----
hoop
hook
hello
hi
breakfast
brunch
brush
bank

```

## public api

### trie
```go
type Trie struct {
    Root *Node
}
```
Trie represents a trie object.

```go
func NewTrie() *Trie 
```
NewTrie returns a pointer to an empty trie.

```go
func (t *Trie) InsertWord(word string)
```
InsertWord inserts a new word into the trie.

```go
func (t *Trie) InsertWords(words []string)
```
InsertWords inserts multiple words into the trie.

```go
func (t *Trie) FindWord(word string) bool 
```
FindWord returns a bool if a word exists in the trie.

```go
func (t *Trie) GetAllWords() []string 
```
GetAllWords returns a slice of strings containing all the words in the entire trie.

```go
func (t *Trie) AutoComplete(prefix string) ([]string, error)
```
AutoComplete returns a slice of strings containing all the possible words that can autocomplete the given prefix.

### node

```go
type Node struct {
    IsWord      bool
    Children    map[rune]*Node
}
```
Node represents a Node in the trie.

```go
func NewNode() *Node
```
NewNode returns a pointer to an empty node.

```go
func (n *Node) GetAllSubWords(curr string) []string
```
GetAllSubWords returns a slice of strings containing all the words in the subtrie contained in the current node n.