package word

import (
	"bufio"
	"io"
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
