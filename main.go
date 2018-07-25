package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/imattman/wordplay-go/pkg/runes"
)

func main() {
	v := "1234"
	if len(os.Args) > 1 {
		v = os.Args[1]
	}

	combos := []string{}
	for _, rs := range runes.Combinations([]rune(v)) {
		combos = append(combos, string(rs))
	}
	sort.Strings(combos)

	for _, s := range combos {
		fmt.Println(s)
	}
}
