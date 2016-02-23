package lex

import (
	"fmt"
)

// Match represents a word match and score.
type Match struct {
	Word  Word
	Score int
}

func (m *Match) String() string {
	return fmt.Sprintf("%d\t%s", m.Score, m.Word)
}

// LexiconMatcher defines the interface for finding word matches in a Lexicon.
type LexiconMatcher interface {
	// Lexicon is the underlying lexicon of words available to be matched.
	Lexicon() Lexicon

	// Matches uses the given character rack to find matching words in the
	// underlying Lexicon, returning Matches on an output channel.
	Matches(r Rack) <-chan *Match
}
