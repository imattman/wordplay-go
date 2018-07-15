package word_test

import (
	"testing"

	"github.com/imattman/wordplay-go/pkg/word"
)

var noScore = map[rune]int{}
var vowelsOnly = map[rune]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}
var abc = map[rune]int{'a': 1, 'b': 1, 'c': 1}
var scrab = map[rune]int{
	'a': 1, 'b': 3, 'c': 3, 'd': 2, 'e': 1, 'f': 4,
	'g': 2, 'h': 4, 'i': 1, 'j': 8, 'k': 5, 'l': 1,
	'm': 3, 'n': 1, 'o': 1, 'p': 3, 'q': 10, 'r': 1,
	's': 1, 't': 1, 'u': 1, 'v': 4, 'w': 4, 'x': 8,
	'y': 4, 'z': 10,
}

func TestScorer(t *testing.T) {
	tests := []struct {
		word     string
		rscores  map[rune]int
		expected int
	}{
		{"", noScore, 0},
		{"", vowelsOnly, 0},
		{"", abc, 0},
		{"", scrab, 0},

		{"apple orange", noScore, 0},
		{"apple orange", vowelsOnly, 5},
		{"apple orange", abc, 2},
		{"apple orange", scrab, 16},

		{"aeiou", noScore, 0},
		{"aeiou", vowelsOnly, 5},
		{"aeiou", abc, 1},
		{"aeiou", scrab, 5},

		{"caboose", noScore, 0},
		{"caboose", vowelsOnly, 4},
		{"caboose", abc, 3},
		{"caboose", scrab, 11},
	}

	for _, test := range tests {
		scoreFn := word.MakeScorer(test.rscores)
		w := word.Word(test.word)
		actual := scoreFn.Score(w.Multiset())
		if test.expected != actual {
			t.Errorf("scoring %q expected: %d, actual: %d",
				test.word,
				test.expected,
				actual)
		}
	}
}
