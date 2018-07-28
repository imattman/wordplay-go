package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/imattman/wordplay-go/pkg/runes"
	"github.com/imattman/wordplay-go/pkg/word"
)

var (
	emptyLexicon = word.NewLexicon(nil)
)

func main() {
	var (
		lexiconFile string
		resultLimit int
	)
	flag.StringVar(&lexiconFile, "f", "", "word lexicon file")
	flag.IntVar(&resultLimit, "n", 10, "match result limit")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fatalf("at least one letter must be supplied for letter rack")
	}

	var wordlist = word.DefaultWordList
	if lexiconFile != "" {
		var err error
		wordlist, err = word.LoadFile(lexiconFile)
		if err != nil {
			fatalf("Error loading lexicon: %v", err)
		}
	}

	lex := word.NewLexicon(wordlist)
	chars := []rune(flag.Arg(0))
	rack := runes.NewMultiset(chars)
	matcher := word.NewFullScanMatcher(lex)
	words, err := matcher.Matches(rack)
	if err != nil {
		fatalf("Match error: %v", err)
	}

	scorer := word.DefaultScorer
	scoredWords := make([]word.ScoredWord, 0, len(words))
	for _, w := range words {
		s := word.ScoredWord{W: w, S: scorer.Score(w.Multiset())}
		scoredWords = append(scoredWords, s)
	}

	sort.Sort(word.ScoredWords(scoredWords))
	limit := min(resultLimit, len(scoredWords))
	scoredWords = scoredWords[:limit]

	for _, sw := range scoredWords {
		fmt.Printf("%d\t%s\n", sw.S, sw.W)
	}
}

func fatalf(msg string, args ...interface{}) {
	if !strings.HasSuffix(msg, "\n") {
		msg += "\n"
	}
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}

func loadLexicon(path string) (*word.Lexicon, error) {
	f, err := os.Open(path)
	if err != nil {
		return emptyLexicon, err
	}
	defer f.Close()

	words, err := word.Tokenize(f, strings.ToLower)
	if err != nil {
		return emptyLexicon, err
	}

	return word.NewLexicon(words), nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
