package word_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/imattman/wordplay-go/pkg/word"
)

const cleanSource = `
cat
dog
fish
horse
monkey
zebra
`

const gnarlySource = `
cat
DOG

fish


  Horse
monkey  ZEBRA

`

func TestTokenize(t *testing.T) {
	sources := []string{
		cleanSource,
		gnarlySource,
	}

	for _, src := range sources {
		r := strings.NewReader(src)
		ws, err := word.Tokenize(r, strings.ToLower)
		if err != nil {
			t.Errorf("Unexpected error processing word source %v", err)
		}

		expected := []string{"cat", "dog", "fish", "horse", "monkey", "zebra"}
		if !reflect.DeepEqual(expected, ws) {
			t.Errorf("Tokenized stream %v\ndoesn't match expected words: %v", ws, expected)
		}
	}
}
