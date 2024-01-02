package dictionary

import (
	"bufio"
	"fmt"
	"os"

	"bogo/trie"
)

func loadDictionary(dictFile string) ([]string, error) {
	file, err := os.Open(dictFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return words, nil
}

func LoadDictionary(dictFile string) (*trie.Trie, error) {
	d, err := loadDictionary(dictFile)
	if err != nil {
		fmt.Printf("Error loading dictionary: %s\n", err)
		return nil, err
	}

	t := trie.NewTrie()
	for _, w := range d {
		t.Insert(w)
	}
	return t, nil
}
