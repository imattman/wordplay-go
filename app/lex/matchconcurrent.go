package lex

import "sync"

// ConcurrentMatcher is a LexiconMatcher that finds word matches using additional goroutines.
type ConcurrentMatcher struct {
	Lex  Lexicon
	Filt PartitioningFilter
}

// NewConcurrentMatcher constructs a ConcurrentMatcher.
func NewConcurrentMatcher(lex Lexicon, filter PartitioningFilter) *ConcurrentMatcher {
	filter = filter
	return &ConcurrentMatcher{
		Lex:  lex,
		Filt: filter,
		// Filt: noopFilter,
	}
}

// Lexicon is the underlying lexicon of words available to be matched.
func (cm *ConcurrentMatcher) Lexicon() Lexicon {
	return cm.Lex
}

// Matches the given character rack against words in the underlying Lexicon.
func (cm *ConcurrentMatcher) Matches(rack Rack) ([]*Match, error) {
	var matches []*Match

	lexs, err := cm.Filt.Filter(cm.Lex, rack)
	if err != nil {
		return matches, err
	}

	mchan := make(chan *Match)
	var wg sync.WaitGroup

	for _, l := range lexs {
		wg.Add(1)
		go func(lex Lexicon, rack Rack) {
			for _, w := range lex {
				if rack.ContainsAll(w.Stats()) {
					mchan <- &Match{Word: w}
				}
			}
			wg.Done()
		}(l, rack)
	}

	go func() {
		wg.Wait()
		close(mchan)
	}()

	for m := range mchan {
		matches = append(matches, m)
	}

	return matches, nil
}
