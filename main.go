package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const lexicon = "sowpods.txt"

func init() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
}

func main() {
	rack := os.Args[1]
	fmt.Printf("rack: %q\n", rack)

	wordList, err := loadWordList(lexicon)
	if err != nil {
		log.Fatal(err)
	}

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
	msg := "Usage:  %s [options] <letter rack>\n"
	fmt.Fprintf(os.Stderr, msg, os.Args[0])
}
