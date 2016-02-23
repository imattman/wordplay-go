package cli

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/imattman/wordplay/app/lex"
)

var limit int

func cmdCliRack() cli.Command {
	return cli.Command{
		Name:  "rack",
		Usage: "find word matches for the supplied character Rack",
		// Description: "[rack chars]",
		Action: actionCliRack,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:        "num, n",
				Value:       10,
				Usage:       "limit match count",
				Destination: &limit,
			},
		},
	}
}

func actionCliRack(c *cli.Context) {
	rackChars := combineArgs(c.Args())
	if len(rackChars) < 1 {
		fmt.Fprintf(os.Stderr, "rack can not be empty\n")
		os.Exit(1)
	}

	lexicon, err := lex.LexiconFromFile(lexiconFile)
	if err != nil {
		log.Fatal(err)
	}

	mxr := lex.NewSerialMatcher(lexicon)
	pipe, err := lex.NewPipeline(mxr)
	rack := lex.NewRack(rackChars)

	debug("lexicon:\t%s (%d words)\n", lexiconFile, len(lexicon))
	debug("letter rack:\t%s\n", rack)

	matches, err := pipe.Process(rack)
	debug("matches:\t%d\n", len(matches))

	sort.Sort(sort.Reverse(lex.ByScore(matches)))

	if (len(matches) < limit) || (limit <= 0) {
		limit = len(matches)
	}

	for _, m := range matches[:limit] {
		fmt.Println(m)
	}
}

func combineArgs(args cli.Args) []rune {
	var rack string
	for _, arg := range args {
		rack += strings.TrimSpace(arg)
	}

	return []rune(rack)
}
