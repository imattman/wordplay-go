package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/imattman/wordplay/wp"
)

var lexiconFile string
var limit int
var verbose bool

func init() {
	flag.Usage = func() {
		usage()
	}

	flag.StringVar(&lexiconFile, "f", "sowpods.txt", "word list lexicon file")
	flag.IntVar(&limit, "n", 10, "limit match count")
	flag.BoolVar(&verbose, "v", false, "verbose output")
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		usage()
		os.Exit(1)
	}

	debug("lexicon:\t%s\n", lexiconFile)

	wordList, err := wp.LoadWordList(lexiconFile)
	if err != nil {
		log.Fatal(err)
	}

	rack := flag.Arg(0)
	debug("letter rack:\t%s\n", rack)

	var matches []*wp.Match
	for _, word := range wordList {
		if wp.CanMakeWord(word, rack) {
			matches = append(matches, &wp.Match{word, wp.CalcScore(word)})
		}
	}
	debug("matches:\t%d\n", len(matches))

	sort.Sort(sort.Reverse(wp.ByScore(matches)))

	if (len(matches) < limit) || (limit <= 0) {
		limit = len(matches)
	}

	for _, m := range matches[:limit] {
		fmt.Println(m)
	}
}

func usage() {
	msg := "Usage:  %s [options] <letter rack>\nOptions:\n"
	fmt.Fprintf(os.Stderr, msg, os.Args[0])
	flag.PrintDefaults()
}

func debug(format string, args ...interface{}) error {
	if verbose {
		_, err := fmt.Fprintf(os.Stderr, format, args...)
		return err
	}
	return nil
}
