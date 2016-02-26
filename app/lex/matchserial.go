package lex

// SerialMatcher is a LexiconMatcher that finds word matches using minimal concurrency.
type SerialMatcher struct {
	Lex  Lexicon
	Filt PartitioningFilter
}

// NewSerialMatcher constructs a SerialMatcher.
func NewSerialMatcher(lex Lexicon, filter PartitioningFilter) *SerialMatcher {
	return &SerialMatcher{
		Lex:  lex,
		Filt: filter,
	}
}

// Lexicon is the underlying lexicon of words available to be matched.
func (sm *SerialMatcher) Lexicon() Lexicon {
	return sm.Lex
}

// Matches the given character rack against words in the underlying Lexicon.
func (sm *SerialMatcher) Matches(rack Rack) ([]*Match, error) {
	var matches []*Match

	lexs, err := sm.Filt.Filter(sm.Lex, rack)
	if err != nil {
		return matches, err
	}

	for _, l := range lexs {
		for _, w := range l {
			if rack.ContainsAll(w.Stats()) {
				matches = append(matches, &Match{Word: w})
			}
		}
	}

	return matches, nil
}
