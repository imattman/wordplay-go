package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"html/template"
	"io"
	"log"
	"os"
	"time"

	"github.com/imattman/wordplay-go/pkg/word"
)

var goTmpl = `// generated {{.Timestamp}} from {{.Source}}
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
	out := os.Stdout
	if restultFile != "-" {
		outfile, err := os.Create(restultFile)
		if err != nil {
			log.Fatalln(err)
		}
		defer outfile.Close()
		out = outfile
	}

	data := struct {
		Source    string
		Words     []string
		Timestamp string
	}{
		wordFile,
		words,
		time.Now().Format(time.RFC3339),
	}

	var src bytes.Buffer
	err = t.Execute(&src, data)
	if err != nil {
		log.Fatalln(err)
	}

	formatted, err := format.Source(src.Bytes())
	if err != nil {
		log.Fatalln(err)
	}

	io.Copy(out, bytes.NewReader(formatted))
}

func usage() {
	app := os.Args[0]
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", app)
	flag.PrintDefaults()

	fmt.Fprintf(os.Stderr, `
  Example:

  %s -f resources/sowpods.txt -o pkg/word/wordlist.go
`, app)
}
