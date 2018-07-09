package word

import (
	"bufio"
	"io"
)

// TokenizeWords scans the supplied reader tokenizing on word boundaries and
// returns a slice of the tokenized values.
func TokenizeWords(r io.Reader) ([]string, error) {
	var words []string
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}
