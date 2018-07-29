package runes

import (
	"sort"
	"strings"
)

// Multiset represents a collection of unordered runes with possible duplicates (a.k.a. Bag of runes).
type Multiset struct {
	counts map[rune]int
}

// NewMultiset is the constructor function for Multisets.
func NewMultiset(rs []rune) Multiset {
	return Multiset{counts: sliceToCounts(rs)}
}

// Counts returns the respective counts of each rune in the multiset.
func (ms Multiset) Counts() map[rune]int {
	// defensive copy
	copy := map[rune]int{}
	for r, cnt := range ms.counts {
		copy[r] = cnt
	}

	return copy
}

// Slice returns the slice representation of the multiset.
func (ms Multiset) Slice() []rune {
	rs := countsToSlice(ms.counts)
	sort.Sort(Slice(rs))
	return rs
}

// Unique returns the unique set of runes in the multiset.
func (ms Multiset) Unique() []rune {
	rs := make([]rune, 0, len(ms.counts))
	for r := range ms.counts {
		rs = append(rs, r)
	}
	return rs
}

// ContainsAll checks if this multiset contains all of the same runes with at least the same counts
// as the other supplied for comparison.
func (ms Multiset) ContainsAll(other Multiset) bool {
	for r, otherCnt := range other.counts {
		// verify this rune is contained and has at least the same count
		if cnt, ok := ms.counts[r]; !ok || (cnt < otherCnt) {
			return false
		}
	}
	return true
}

// Equals tests for equality with another Multiset.
// Runes and respective counts must match for the Multisets to be considered equal.
func (ms Multiset) Equals(other Multiset) bool {
	if len(ms.counts) != len(other.counts) {
		return false
	}
	for r, otherCnt := range other.counts {
		// verify this rune is contained and has at least the same count
		if cnt, ok := ms.counts[r]; !ok || (cnt != otherCnt) {
			return false
		}
	}
	return true
}

func (ms Multiset) String() string {
	rs := ms.Slice()
	ss := make([]string, 0, len(rs))
	for _, r := range rs {
		ss = append(ss, string(r))
	}

	return "[" + strings.Join(ss, ",") + "]"
}

func sliceToCounts(rs []rune) map[rune]int {
	counts := map[rune]int{}
	for _, r := range rs {
		counts[r] = counts[r] + 1
	}

	return counts
}

func countsToSlice(m map[rune]int) []rune {
	rs := make([]rune, 0, len(m)) // might be larger, but reasonable starting size
	for r, cnt := range m {
		for i := 0; i < cnt; i++ {
			rs = append(rs, r)
		}
	}

	return rs
}
