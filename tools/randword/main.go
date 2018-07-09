// A very simple word sampler.
// N total random words are sampled from the supplied file arguments or
// from STDIN if no additional arguments are given.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"time"
)

func main() {
	var (
		sampleSize int
	)
	flag.IntVar(&sampleSize, "n", 100, "Number of words to sample")
	flag.Parse()

	words, err := loadWords(flag.Args())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// shuffle all words
	rand.Seed(time.Now().UnixNano())
	for i := len(words) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		// fmt.Printf("swapping [%d] %q <-> [%d] %q\n", i, words[i], j, words[j])
		words[i], words[j] = words[j], words[i]
	}

	// sample upper bound is Min(sampleSize, len(words))
	n := len(words)
	if n > sampleSize {
		n = sampleSize
	}

	sample := words[:n]
	sort.Strings(sample)
	for _, w := range sample {
		fmt.Println(w)
	}
}

func loadWords(args []string) ([]string, error) {
	if len(args) < 1 {
		return tokenize(os.Stdin)
	}

	var words []string
	for _, path := range args {
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}

		ws, err := tokenize(f)
		f.Close() // be sure to close file before returning tokenize error
		if err != nil {
			return nil, err
		}

		words = append(words, ws...)
	}

	return words, nil
}

func tokenize(r io.Reader) ([]string, error) {
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
