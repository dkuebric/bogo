package trie

import (
	"fmt"
	"testing"
)

func TestTrieBasic(t *testing.T) {
	words := []string{
		"A",
		"AA",
		"AB",
		"B",
		"BB",
		"BAT",
		"CAT",
		"CATTY",
		"CATS",
	}

	trie := NewTrie()

	for _, w := range words {
		fmt.Printf("Inserting %s\n", w)
		trie.Insert(w)
		if !trie.Search(w) {
			t.Fatalf("Failed to find just-inserted element %s", w)
		}
	}
	for _, w := range words {
		if !trie.Search(w) {
			t.Fatalf("Failed to find element %s", w)
		}
	}

	missing := []string{
		"D",
		"BATTY",
	}
	for _, w := range missing {
		if trie.Search(w) {
			t.Fatalf("Found non-existent element %s", w)
		}
	}
}
