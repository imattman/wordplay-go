package lex

import (
	"sort"
	"strconv"
	"strings"
)

// Lexicon represents a list of valid words.
type Lexicon []Word

// RuneStats defines stat behaviors for a group of runes.
type RuneStats interface {
	// Unique returns the set of unique runes in the group.
	Unique() []rune

	// Count returns the count of a particular rune in the group.
	Count(r rune) int
}

// ComparableRuneStats defines behaviors for comparing against another RuneStats.
type ComparableRuneStats interface {
	RuneStats

	// ContainsAll tests if this rune group contains all the same runes as the other passed
	// with at least the same counts.  This rune group may contain a super set of
	// extra runes (or higher counts) and still test true.
	ContainsAll(other RuneStats) bool
}

// Word adds CharStats functionality to a string.
type Word string

func (w Word) String() string {
	return string(w)
}

// Stats returns rune stats for the word.
func (w Word) Stats() RuneStats {
	return groupByRune([]rune(string(w)))
}

// Rack represents a group of playable characters.
type Rack struct {
	chars []rune
	runeCounts
}

func (r Rack) String() string {
	c := make([]rune, len(r.chars))
	copy(c, r.chars)
	sort.Sort(RuneSlice(c))

	return string(c)
}

// NewRack contructs a character rack.
func NewRack(chars []rune) Rack {
	return Rack{chars, groupByRune(chars)}
}

type runeCounts map[rune]int

func (rc runeCounts) Unique() []rune {
	u := make([]rune, 0, len(rc))
	for c := range rc {
		u = append(u, c)
	}

	return u
}

func (rc runeCounts) Count(r rune) int {
	return rc[r]
}

func (rc runeCounts) ContainsAll(other RuneStats) bool {
	otherRunes := other.Unique()
	for _, r := range otherRunes {
		if other.Count(r) > rc[r] {
			return false
		}
	}
	return true
}

func (rc runeCounts) String() string {
	u := rc.Unique()
	sort.Sort(RuneSlice(u))

	elems := make([]string, 0, len(u)+2)
	elems = append(elems, "{")
	for _, c := range u {
		elems = append(elems,
			"'"+string(c)+"':"+
				strconv.Itoa(rc[c]))
	}
	elems = append(elems, "}")

	return strings.Join(elems, " ")
}

func groupByRune(chars []rune) runeCounts {
	m := make(map[rune]int)
	for _, c := range chars {
		m[c]++
	}

	return runeCounts(m)
}
