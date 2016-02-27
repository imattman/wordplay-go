package lex

import (
	"fmt"
)

// Match represents a word match and score.
type Match struct {
	Word  Word `json:"word"`
	Score int  `json:"score"`
}

func (m *Match) String() string {
	return fmt.Sprintf("%d\t%s", m.Score, m.Word)
}

// LexiconMatcher defines the interface for finding word matches in a Lexicon.
type LexiconMatcher interface {
	// Lexicon is the underlying lexicon of words available to be matched.
	Lexicon() Lexicon

	// Matches uses the given character rack to find matching words in the
	// underlying Lexicon.
	Matches(r Rack) ([]*Match, error)
}

// NewLimitProcessor creates a MatchProcessor that caps the number of word matches.
func NewLimitProcessor(limit int) MatchProcessor {
	fn := func(ms []*Match) ([]*Match, error) {
		if (limit > 0) && (limit < len(ms)) {
			ms = ms[:limit]
		}

		return ms, nil
	}

	return MatchProcessorFunc(fn)
}

// MatchProcessor defines the interface for sorting, truncating, and generally
// post-processing a list of word matches.
type MatchProcessor interface {
	Process(ms []*Match) ([]*Match, error)
}

// MatchProcessorFunc is a type adapter to allow the use of an ordinary function
// as a MatchProcessor.
type MatchProcessorFunc func(ms []*Match) ([]*Match, error)

// Process calls fn(ms)
func (fn MatchProcessorFunc) Process(ms []*Match) ([]*Match, error) {
	return fn(ms)
}
