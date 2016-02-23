package cli

import (
	"fmt"
	"sort"

	"github.com/codegangsta/cli"
	"github.com/imattman/wordplay/app/lex"
)

func cmdFoo() cli.Command {
	return cli.Command{
		Name:   "foo",
		Usage:  "dev testing cli options",
		Action: testFunctionality,
	}
}

func testFunctionality(c *cli.Context) {
	s := "foobarbaz"
	w := lex.Word(s)
	fmt.Printf("word: %s\t%v\n", w, unique(w.Stats()))

	r := lex.NewRack([]rune(s))
	fmt.Printf("rack: %s\t%v\n", r, unique(r))
}

func unique(rs lex.RuneStats) string {
	u := rs.Unique()
	sort.Sort(lex.RuneSlice(u))

	return string(u)
}
