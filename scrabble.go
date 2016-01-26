package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var letterScore = map[rune]int{
	'a': 1, 'b': 3, 'c': 3, 'd': 2, 'e': 1, 'f': 4,
	'g': 2, 'h': 4, 'i': 1, 'j': 8, 'k': 5, 'l': 1,
	'm': 3, 'n': 1, 'o': 1, 'p': 3, 'q': 10, 'r': 1,
	's': 1, 't': 1, 'u': 1, 'v': 4, 'w': 4, 'x': 8,
	'y': 4, 'z': 10,
}

func calcScore(word string) int {
	var total int
	for _, c := range word {
		total += letterScore[c]
	}

	return total
}

// Match represents a word match and its scrabble score
type match struct {
	word  string
	score int
}

func (m *match) String() string {
	return fmt.Sprintf("%d\t%s", m.score, m.word)
}

// byScore implements sort.Interface for []*match based on word score
type byScore []*match

func (m byScore) Len() int {
	return len(m)
}

func (m byScore) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m byScore) Less(i, j int) bool {
	return m[i].score < m[j].score
}

// loads and converts to lower case words from a text file
func loadWordList(wordFile string) ([]string, error) {
	file, err := os.Open(wordFile)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(bufio.NewReader(file))
	var words []string
	for scanner.Scan() {
		w := strings.ToLower(scanner.Text())
		words = append(words, w)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}
