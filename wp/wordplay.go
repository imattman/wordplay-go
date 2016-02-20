package wp

import (
	"fmt"
)

// Lexicon represents a list of valid words.
type Lexicon []string

// Rack represents a collection of characters that can be used to match against words of a Lexicon.
type Rack interface {
	// Chars returns all the characters in the Rack with possible duplicates.
	Chars() []rune

	// Chars returns all the unique characters in the Rack.
	CharsDistinct() []rune

	// CanMakeWord tests if a word can be formed using the Rack characters.
	CanMakeWord(w string) bool
}

// Match represents a word match and score.
type Match struct {
	Word  string
	Score int
}

func (m *Match) String() string {
	return fmt.Sprintf("%d\t%s", m.Score, m.Word)
}

// ByScore implements sort.Interface for []*Match based on word score
type ByScore []*Match

// Required by sort.Interface
func (m ByScore) Len() int {
	return len(m)
}

// Required by sort.Interface
func (m ByScore) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// Required by sort.Interface
func (m ByScore) Less(i, j int) bool {
	return m[i].Score < m[j].Score
}

var scrabScores = map[rune]int{
	'a': 1, 'b': 3, 'c': 3, 'd': 2, 'e': 1, 'f': 4,
	'g': 2, 'h': 4, 'i': 1, 'j': 8, 'k': 5, 'l': 1,
	'm': 3, 'n': 1, 'o': 1, 'p': 3, 'q': 10, 'r': 1,
	's': 1, 't': 1, 'u': 1, 'v': 4, 'w': 4, 'x': 8,
	'y': 4, 'z': 10,
}

func scoreByLetter(word string, letterScore map[rune]int) int {
	var total int
	for _, c := range word {
		total += letterScore[c]
	}

	return total
}
