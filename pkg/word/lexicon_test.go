package word_test

import (
	"reflect"
	"testing"

	"github.com/imattman/wordplay-go/pkg/word"
)

var sourceWords = []string{
	"cat",
	"dog",
	"apple",
	"banana",
	"pineapple",
	"orange",
	"plumb",
	"hammer",
	"nail",
	"saw",
}

func TestLexicon_ContainsWord(t *testing.T) {
	lex := word.NewLexicon(sourceWords)
	for _, s := range sourceWords {
		word, ok := lex.Lookup(s)
		if !ok {
			t.Errorf("Expected word %q not in Lexicon %v", word, lex)
		}
	}

	unknown := []string{
		"programmer",
		"wacky",
		"airplane",
		"river",
	}
	for _, s := range unknown {
		word, ok := lex.Lookup(s)
		if ok {
			t.Errorf("Unexpected word %q found in Lexicon %v", word, lex)
		}
	}

}

func TestLexicon_AllWords(t *testing.T) {
	lex := word.NewLexicon(sourceWords)

	// order isn't guaranteed so convert to sets for comparison
	expected := toSet(toWords(sourceWords))
	actual := lex.AllWords()
	if !reflect.DeepEqual(expected, toSet(actual)) {
		t.Errorf("Expected list (size %d) does not match Lexicon %v (size %d)",
			len(expected),
			lex,
			len(actual),
		)
	}
}

func toWords(xs []string) []word.Word {
	ws := make([]word.Word, 0, len(xs))
	for _, w := range xs {
		ws = append(ws, word.Word(w))
	}

	return ws
}

func toSet(ws []word.Word) map[word.Word]struct{} {
	set := map[word.Word]struct{}{}
	for _, val := range ws {
		set[val] = struct{}{}
	}

	return set
}
