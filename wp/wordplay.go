package wp

import (
	"strings"
	"fmt"
)

// Match represents a word match and its scrabble score
type Match struct {
	Word  string
	Score int
}

func (m *Match) String() string {
	return fmt.Sprintf("%d\t%s", m.Score, m.Word)
}

// ByScore implements sort.Interface for []*match based on word score
type ByScore []*Match

func (m ByScore) Len() int {
	return len(m)
}

func (m ByScore) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m ByScore) Less(i, j int) bool {
	return m[i].Score < m[j].Score
}



func CanMakeWord(word, availChars string) bool {
	for _, c := range word {
		if i := strings.IndexRune(availChars, c); i > -1 {
			availChars = availChars[:i] + availChars[i+1:] // remove char from avail pool
		} else {
			return false
		}
	}

	return true
}



