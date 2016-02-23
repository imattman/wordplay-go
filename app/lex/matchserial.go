package lex

// type ScoringFunc func(word string) int
// type PartitioningFunc func(lex Lexicon, rack Rack) []Lexicon

// SerialMatcher is a LexiconMatcher that finds matches using minimal concurrency.
type SerialMatcher struct {
	lexicon Lexicon
}

// NewSerialMatcher constructs a SerialMatcher.
func NewSerialMatcher(lex Lexicon) *SerialMatcher {
	return &SerialMatcher{lex}
}

// Lexicon is the underlying lexicon of words available to be matched.
func (sm *SerialMatcher) Lexicon() Lexicon {
	return sm.lexicon
}

// Matches the given character rack against words in the underlying Lexicon,
// returning Matches on an output channel.
func (sm *SerialMatcher) Matches(r Rack) <-chan *Match {
	matches := make(chan *Match)

	go func() {
		for _, w := range sm.lexicon {
			if r.ContainsAll(w.Stats()) {
				matches <- &Match{Word: w}
			}
		}
		close(matches)
	}()

	return matches
}
