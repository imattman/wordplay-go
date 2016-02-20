package wp

type ScoringFunc func(word string) int
type PartitioningFunc func(lex Lexicon, rack Rack) []Lexicon

type Matcher interface {
	Matches(lex Lexicon, rack Rack) ([]*Match, error)
}

type MatcherOptions struct {
	Scorer      ScoringFunc
	Partitioner PartitioningFunc
}

type SerialMatcher struct {
	MatcherOptions
}

func NewSerialMatcher(fns ...func(mo *MatcherOptions)) *SerialMatcher {
	mo := MatcherOptions{
		Scorer: func(word string) int {
			return scoreByLetter(word, scrabScores)
		},
		Partitioner: func(lex Lexicon, r Rack) []Lexicon {
			return []Lexicon{lex}
		},
	}

	for _, f := range fns {
		f(&mo)
	}

	return &SerialMatcher{MatcherOptions: mo}
}

func (sm *SerialMatcher) Matches(lex Lexicon, rack Rack) ([]*Match, error) {
	var ms []*Match
	for _, word := range lex {
		if rack.CanMakeWord(word) {
			ms = append(ms, &Match{word, sm.Scorer(word)})
		}
	}

	return ms, nil
}

//type ConcurrentMatcher struct {}
