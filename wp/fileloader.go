package wp

import (
	"bufio"
	"os"
	"strings"
)

// LexiconFromFile loads a Lexicon word list text file, converting words to lower case.
func LexiconFromFile(path string) (Lexicon, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(bufio.NewReader(file))
	var words []string
	for scanner.Scan() {
		w := strings.TrimSpace(strings.ToLower(scanner.Text()))
		if len(w) > 0 {
			words = append(words, w)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}
