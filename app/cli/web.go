package cli

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/imattman/wordplay-go/app/lex"
	"github.com/imattman/wordplay-go/app/web"
)

var port int

func cmdWeb() cli.Command {
	return cli.Command{
		Name: "web",
		// Aliases: []string{"serve"},
		Usage:  "starts web application to serve requests",
		Action: actionWeb,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:        "port, p",
				Value:       8080,
				Usage:       "listening port",
				EnvVar:      "PORT",
				Destination: &port,
			},
		},
	}
}

func actionWeb(c *cli.Context) {
	log.Println("loading lexicon:", lexiconFile)
	lexicon, err := lex.LexiconFromFile(lexiconFile)
	if err != nil {
		log.Fatal(err)
	}

	//mxr := lex.NewSerialMatcher(lexicon, lex.NoopFilter)
	mxr := lex.NewConcurrentMatcher(lexicon, lex.PrePartitionByFirstChar(lexicon))
	pipeline, err := lex.NewPipeline(mxr)
	if err != nil {
		log.Panic(err)
	}

	web.Serve(port, pipeline)
}
