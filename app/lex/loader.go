package lex

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// LexiconFromFile loads a Lexicon word list from a text file.
func LexiconFromFile(path string) (Lexicon, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	return LexiconFromReader(file)
}

// LexiconFromReader builds a Lexicon word list from the supplied Reader,
// converting words to lower case.
func LexiconFromReader(r io.Reader) (Lexicon, error) {
	scanner := bufio.NewScanner(bufio.NewReader(r))
	var words []Word
	for scanner.Scan() {
		w := strings.TrimSpace(strings.ToLower(scanner.Text()))
		if len(w) > 0 {
			words = append(words, Word(w))
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}
