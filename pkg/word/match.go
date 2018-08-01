package word

import "github.com/imattman/wordplay-go/pkg/runes"

// Matcher defines the interface for finding Words that meet match criteria.
type Matcher interface {
	Matches(rack runes.Multiset) ([]Word, error)
}

type fullScanMatcher struct {
	lex *Lexicon
}

// NewFullScanMatcher contructs a basic Matcher that uses a full lexicon scan to match words.
func NewFullScanMatcher(lex *Lexicon) Matcher {
	return &fullScanMatcher{lex: lex}
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

type partitionedMatcher struct {
	wordsByPrefix map[rune][]Word
}

// NewPartitionedMatcher contructs a Matcher that partitions the underlying lexicon for more
// efficient matching scans.
func NewPartitionedMatcher(lex *Lexicon) Matcher {
	partitions := map[rune][]Word{}
	for _, word := range lex.AllWords() {
		if len(word) < 1 {
			continue
		}
		firstChar := []rune(string(word))[0]
		ws, ok := partitions[firstChar]
		if !ok {
			ws = []Word{}
		}
		ws = append(ws, word)
		partitions[firstChar] = ws
	}

	return &partitionedMatcher{wordsByPrefix: partitions}
}

func (pm *partitionedMatcher) Matches(rack runes.Multiset) ([]Word, error) {
	matches := make([]Word, 0)
	rack.Slice()
	for _, r := range rack.Unique() {
		for _, w := range pm.wordsByPrefix[r] {
			if rack.ContainsAll(w.Multiset()) {
				matches = append(matches, w)
			}
		}
	}

	return matches, nil
}
