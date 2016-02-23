package lex

import (
	"sort"
	"sync"
)

// MatchPipeline is the assembly for matching, scoring, and finishing functionality.
type MatchPipeline struct {
	Matcher     LexiconMatcher
	ScoringFunc func(w Word) int
	// Finishers []
}

// NewPipeline constructs a MatchPipeline.
func NewPipeline(matcher LexiconMatcher, opt ...func(mp *MatchPipeline)) (*MatchPipeline, error) {
	mp := MatchPipeline{
		Matcher:     matcher,
		ScoringFunc: defaultScoringFunc,
	}

	// allow option functions to make modifications
	for _, f := range opt {
		f(&mp)
	}

	return &mp, nil
}

// Process applies the stages of the pipeline to a character rack argument.
func (mp *MatchPipeline) Process(r Rack) ([]*Match, error) {
	mchan := mp.Matcher.Matches(r)
	var matches []*Match
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for m := range mchan {
			m.Score = mp.ScoringFunc(m.Word)
			matches = append(matches, m)
		}
		wg.Done()
	}()

	wg.Wait()
	sort.Sort(sort.Reverse(ByScore(matches)))

	return matches, nil

	// if (len(matches) < limit) || (limit <= 0) {
	// 	limit = len(matches)
	// }
	//
	// for _, m := range matches[:limit] {
	// 	fmt.Println(m)
	// }
}
