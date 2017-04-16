package cli

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

// see:
// http://nathanleclaire.com/blog/2014/08/31/why-codegangstas-cli-package-is-the-bomb-and-you-should-use-it/

const appName = "wordplay"

var app *cli.App
var lexiconFile string
var verbose bool

// Execute is the main entry for the program.
func Execute() {
	app = cli.NewApp()
	app.Name = appName
	app.Version = "0.1.0 (funkychicken)"
	app.Usage = "Look up valid words based on a 'Rack' of playable characters"
	// app.UsageText = "this is the UsageText"

	registerCommands()
	// app.Action = actionCliRack
	app.Action = cli.ShowAppHelp

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "file, f",
			Value:       "resources/sowpods.txt",
			Usage:       "word list lexicon file",
			EnvVar:      "WP_LEXICON_FILE",
			Destination: &lexiconFile,
		},
		cli.BoolFlag{
			Name:        "verbose, V",
			Usage:       "additional output",
			Destination: &verbose,
		},
	}

	app.RunAndExitOnError()
}

func registerCommands() {
	app.Commands = []cli.Command{
		cmdCliRack(),
		cmdWeb(),
	}
}

func debug(format string, args ...interface{}) error {
	if verbose {
		_, err := fmt.Fprintf(os.Stderr, format, args...)
		return err
	}
	return nil
}
