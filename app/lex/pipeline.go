package lex

// MatchPipeline is the assembly for matching, scoring, and finishing functionality.
type MatchPipeline struct {
	Matcher    LexiconMatcher
	Processors []MatchProcessor
}

// NewPipeline constructs a MatchPipeline.
func NewPipeline(matcher LexiconMatcher, opt ...func(mp *MatchPipeline)) (*MatchPipeline, error) {
	pipe := MatchPipeline{
		Matcher: matcher,
		Processors: []MatchProcessor{
			ScrabbleScoreProcessor,
			SortByScoreProcessor,
		},
	}

	// allow option functions to make modifications
	for _, f := range opt {
		f(&pipe)
	}

	return &pipe, nil
}

// AddProcessor is a convenience method for appending a MatchProcessor to the
// existing chain of processors.
func (mp *MatchPipeline) AddProcessor(proc MatchProcessor) {
	mp.Processors = append(mp.Processors, proc)
}

// Process applies the stages of the pipeline to a character rack argument.
func (mp *MatchPipeline) Process(rack Rack) ([]*Match, error) {
	matches, err := mp.Matcher.Matches(rack)
	if err != nil {
		return matches, err
	}

	for _, proc := range mp.Processors {
		matches, err = proc.Process(matches)
		if err != nil {
			return matches, err
		}
	}

	return matches, nil

	// if (len(matches) < limit) || (limit <= 0) {
	// 	limit = len(matches)
	// }
	//
	// for _, m := range matches[:limit] {
	// 	fmt.Println(m)
	// }
}
