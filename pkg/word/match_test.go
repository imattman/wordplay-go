package word_test

import (
	"reflect"
	"testing"

	"github.com/imattman/wordplay-go/pkg/runes"
	"github.com/imattman/wordplay-go/pkg/word"
)

var matchingLexWords = []string{
	"aaa",
	"bbb",
	"apple",
	"orange",
	"atom",
	"vim",
	"wubba",
}

func TestFullScanMatcher(t *testing.T) {
	lex := word.NewLexicon(matchingLexWords)
	matcher := word.NewFullScanMatcher(lex)

	testCompleteWordMatcher(t, matcher)
}

func testCompleteWordMatcher(t *testing.T, matcher word.Matcher) {
	tests := []struct {
		rack     string
		expected []string
	}{
		{"", []string{}},
		{"aaa", []string{"aaa"}},
		{"aaaa", []string{"aaa"}},
		{"aaabbb", []string{"aaa", "bbb"}},
		{"vimxxwubbaxx", []string{"vim", "wubba"}},
	}

	for _, test := range tests {
		rack := runes.NewMultiset([]rune(test.rack))
		actual, err := matcher.Matches(rack)
		if err != nil {
			t.Errorf("Rack %q match error %s", rack, err)
		}

		expSet := toSet(toWords(test.expected))
		actSet := toSet(actual)

		if !reflect.DeepEqual(expSet, actSet) {
			t.Errorf("Rack(%q) bad match, expected %s, actual %s",
				rack,
				test.expected,
				actual)
		}
	}
}
