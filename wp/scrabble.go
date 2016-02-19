package wp

import (
	"bufio"
	"os"
	"strings"
)

var letterScore = map[rune]int{
	'a': 1, 'b': 3, 'c': 3, 'd': 2, 'e': 1, 'f': 4,
	'g': 2, 'h': 4, 'i': 1, 'j': 8, 'k': 5, 'l': 1,
	'm': 3, 'n': 1, 'o': 1, 'p': 3, 'q': 10, 'r': 1,
	's': 1, 't': 1, 'u': 1, 'v': 4, 'w': 4, 'x': 8,
	'y': 4, 'z': 10,
}

func CalcScore(word string) int {
	var total int
	for _, c := range word {
		total += letterScore[c]
	}

	return total
}

// LoadWordList loads words from a text file and converts to lower case
func LoadWordList(wordFile string) ([]string, error) {
	file, err := os.Open(wordFile)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(bufio.NewReader(file))
	var words []string
	for scanner.Scan() {
		w := strings.ToLower(scanner.Text())
		words = append(words, w)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}
