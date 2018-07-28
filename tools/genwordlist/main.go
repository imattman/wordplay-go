package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/imattman/wordplay-go/pkg/word"
)

var goTmpl = `// generated from {{.Source}}
package word

// DefaultWordList is the default lexicon if an external source is not loaded.
var DefaultWordList = []string{
	{{range .Words}}"{{.}}",
	{{end}}  }
`

func main() {
	var (
		wordFile    string
		restultFile string
	)
	flag.StringVar(&wordFile, "f", "resources/sowpods.txt", "Word list file")
	flag.StringVar(&restultFile, "o", "pkg/word/wordlist.go", "Generated Go source file")
	flag.Usage = usage
	flag.Parse()

	words, err := word.LoadFile(wordFile)
	if err != nil {
		log.Fatalln(err)
	}

	t := template.Must(template.New("source").Parse(goTmpl))

	// write to STDOUT if '-' is specified, otherwise standard file result
	result := os.Stdout
	if restultFile != "-" {
		out, err := os.Create(restultFile)
		if err != nil {
			log.Fatalln(err)
		}
		defer out.Close()
		result = out
	}

	data := struct {
		Source string
		Words  []string
	}{
		wordFile,
		words,
	}

	err = t.Execute(result, data)
	if err != nil {
		log.Fatalln(err)
	}
}

func usage() {
	app := os.Args[0]
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", app)
	flag.PrintDefaults()

	fmt.Fprintf(os.Stderr, `
  Example:

  %s -f resources/sowpods.txt -o pkg/word/wordlist.go && \
      gofmt -w pkg/word/wordlist.go
`, app)
}
