package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"sort"

	"bogo/dictionary"
	"bogo/trie"
)

var (
	boardFile = flag.String("board", "board.csv", "a csv file in the format of letter,letter,letter,letter\nletter,...'")
	dictFile  = flag.String("dict", "dictionary.txt", "a txt file with one word per line")
)

type Board struct {
	g [4][4]string
}

func loadBoard() (*Board, error) {
	b := &Board{}
	file, err := os.Open(*boardFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}
	for j, line := range lines {
		for i, c := range line {
			b.g[i][j] = c
		}
	}

	return b, nil
}

func solve(b *Board, t *trie.Trie) []string {
	results := []string{}
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			results = append(results, solveFromBase(b.g[y][x], x, y, []coord{}, b, t)...)
		}
	}

	deduper := make(map[string]bool)
	for _, r := range results {
		deduper[r] = true
	}
	results = []string{}
	for k, _ := range deduper {
		results = append(results, k)
	}
	sort.Strings(results)
	return results
}

type coord struct {
	X int
	Y int
}

func adjacent(x int, y int, used []coord) []coord {
	c := []coord{}
	for i := x - 1; i < x+2; i++ {
		for j := y - 1; j < y+2; j++ {
			if i < 0 || i > 3 || j < 0 || j > 3 {
				continue
			}
			if i == x && j == y {
				continue
			}
			for _, uc := range used {
				if uc.X == i && uc.Y == j {
					goto skip
				}
			}
			c = append(c, coord{i, j})
		skip:
		}
	}
	return c
}

func solveFromBase(base string, x int, y int, used []coord, b *Board, t *trie.Trie) []string {
	results := []string{}
	prefix, match := t.Search(base)
	if prefix {
		if match && len(base) >= 3 {
			results = append(results, base)
		}
		for _, coords := range adjacent(x, y, used) {
			results = append(results, solveFromBase(base+b.g[coords.Y][coords.X], coords.X, coords.Y, append(used, coord{x, y}), b, t)...)
		}
	}
	return results
}

func main() {
	flag.Parse()
	b, err := loadBoard()
	if err != nil {
		fmt.Printf("Error loading board: %s\n", err)
		return
	}

	t, err := dictionary.LoadDictionary(*dictFile)
	if err != nil {
		fmt.Printf("Error loading dictionary: %s\n", err)
		return
	}

	res := solve(b, t)
	fmt.Printf("Genius results:\n%+v\n", res)
	fmt.Printf("%d\n", len(res))
}
