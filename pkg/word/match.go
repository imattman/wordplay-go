package word

import "github.com/imattman/wordplay-go/pkg/runes"

// Matcher defines the interface for finding Words that meet match criteria.
type Matcher interface {
	Matches(rack runes.Multiset) ([]Word, error)
}

type fullScanMatcher struct {
	lex *Lexicon
}

func (fm *fullScanMatcher) Matches(rack runes.Multiset) ([]Word, error) {
	matches := make([]Word, 0)
	for _, w := range fm.lex.AllWords() {
		if rack.ContainsAll(w.Multiset()) {
			matches = append(matches, w)
		}
	}

	return matches, nil
}

// NewFullScanMatcher contructs a basic Matcher that uses a full lexicon scan to match words.
func NewFullScanMatcher(lex *Lexicon) Matcher {
	return &fullScanMatcher{lex: lex}
}
