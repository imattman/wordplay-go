package word

import (
	"github.com/imattman/wordplay-go/pkg/runes"
)

// Scorer defines the interface for scoring Words and groups of runes.
type Scorer interface {
	Score(rs runes.Multiset) int
}

// DefaultScorer uses the default letter score values to calculate a score.
var DefaultScorer = MakeScorer(defaultRuneScores())

// MakeScorer generates a Scorer using the supplied map of runes with their
// respective scores.
func MakeScorer(runeScore map[rune]int) Scorer {
	f := func(rs runes.Multiset) int {
		var total int
		for r, cnt := range rs.Counts() {
			total += cnt * runeScore[r]
		}

		return total
	}

	return scoringFunc(f)
}

type scoringFunc func(rs runes.Multiset) int

func (f scoringFunc) Score(rs runes.Multiset) int {
	return f(rs)
}

// letter scores pulled from a popular word game
func defaultRuneScores() map[rune]int {
	scores := map[int]string{
		0:  " ",
		1:  "aeilnorstu",
		2:  "dg",
		3:  "bcmp",
		4:  "fhvwy",
		5:  "k",
		8:  "jx",
		10: "qz",
	}

	rscore := map[rune]int{}
	for score, runes := range scores {
		for _, r := range []rune(runes) {
			rscore[r] = score
		}
	}

	return rscore
}
