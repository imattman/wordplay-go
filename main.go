package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var lexiconFile string
var verbose bool

func init() {
	flag.Usage = func() {
		usage()
	}

	flag.StringVar(&lexiconFile, "f", "sowpods.txt", "word list lexicon file")
	flag.BoolVar(&verbose, "v", false, "verbose output")

	flag.Parse()
	if len(flag.Args()) < 1 {
		usage()
		os.Exit(1)
	}
}

func main() {
	wordList, err := loadWordList(lexiconFile)
	if err != nil {
		log.Fatal(err)
	}

	rack := flag.Arg(0)
	debug("letter rack:\t%s\n", rack)

	var matches []*match
	for _, word := range wordList {
		if canMakeWord(word, rack) {
			matches = append(matches, &match{word, calcScore(word)})
		}
	}

	sort.Sort(sort.Reverse(byScore(matches)))
	for _, m := range matches {
		fmt.Println(m)
	}
}

func canMakeWord(word, availChars string) bool {
	for _, c := range word {
		if i := strings.IndexRune(availChars, c); i > -1 {
			availChars = availChars[:i] + availChars[i+1:] // remove char from avail pool
		} else {
			return false
		}
	}

	return true
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
