package word

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// Tokenize scans the supplied reader tokenizing on whitespace boundaries and
// returns a slice of the tokenized values.
func Tokenize(r io.Reader, normalizers ...func(s string) string) ([]string, error) {
	var words []string
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		t := scanner.Text()
		for _, normFn := range normalizers {
			t = normFn(t)
		}
		words = append(words, t)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

// LoadFile reads a lexicon word list from the specified file, performing standard
// normalizing transformations on the text (e.g. lowercase).
func LoadFile(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Tokenize(f, strings.ToLower)
}
