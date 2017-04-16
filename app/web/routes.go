package web

import (
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/go-zoo/bone"
	"github.com/imattman/wordplay-go/app/lex"
)

// Serve starts the HTTP server.
func Serve(port int, pipe *lex.MatchPipeline) {
	apiMux := bone.New()
	matchHandler := matchesHandler{pipe: pipe}
	apiMux.Get("/matches/:rack", matchHandler)

	mux := bone.New()
	mux.SubRoute("/api", apiMux)

	// Handle take http.Handler
	// mux.Handle("/", http.HandlerFunc(RootHandler))

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":" + strconv.Itoa(port))
}
