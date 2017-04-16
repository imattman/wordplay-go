package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/imattman/wordplay-go/app/lex"
)

const rackParam = "rack"

type matchResult struct {
	Rack    []string     `json:"rack"`
	Matches []*lex.Match `json:"matches"`
}

type matchesHandler struct {
	pipe *lex.MatchPipeline
}

func (h matchesHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r := bone.GetValue(req, rackParam)
	rack := lex.NewRack([]rune(r))
	log.Print("rack:", rack)
	log.Print("rack (slice):", rack.StringSlice())

	ms, err := h.pipe.Process(rack)
	if err != nil {
		log.Println("********", err)
		return
	}
	log.Println("matches:", len(ms))

	result := matchResult{Rack: rack.StringSlice(), Matches: ms}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&result)
}
