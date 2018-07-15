package word

import (
	"errors"

	"github.com/imattman/wordplay-go/pkg/runes"
)

const (
	// place holder value returned when an entry is not found in a Lexicon.
	notFound = Word("")
)

var (
	// ErrNotFound is returned when a word is not found in a Lexicon.
	ErrNotFound = errors.New("word not found")
)

// Word represents a valid word contained in a Lexicon.
type Word string

// Multiset returns the rune multiset representation of a Word.
func (w Word) Multiset() runes.Multiset {
	return runes.NewMultiset([]rune(string(w)))
}

// A MatchingLexicon provides both Lexicon and Matcher functionality.
type MatchingLexicon struct {
	Lexicon
	Matcher
}

// Lexicon represents a collection of valid words.
type Lexicon struct {
	wordList []Word
	lookup   map[Word]struct{}
}

// AllWords returns all words contained in the Lexicon.
func (l *Lexicon) AllWords() []Word {
	return l.wordList
}

// Lookup tests if the given string is a valid word in the Lexicon.
func (l *Lexicon) Lookup(s string) (Word, bool) {
	w := Word(s)
	if _, ok := l.lookup[w]; !ok {
		return notFound, false
	}

	return w, true
}

// NewLexicon contructs a Lexicon of words from a slice of strings.
func NewLexicon(words []string) *Lexicon {
	wordList := make([]Word, 0, len(words))
	lookup := make(map[Word]struct{})

	for _, s := range words {
		w := Word(s)
		wordList = append(wordList, w)
		lookup[w] = struct{}{}
	}

	return &Lexicon{
		wordList: wordList,
		lookup:   lookup,
	}
}
