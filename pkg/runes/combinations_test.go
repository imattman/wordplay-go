package runes

import (
	"reflect"
	"testing"
)

func TestCombinations(t *testing.T) {
	tests := []struct {
		name   string
		source []rune
		want   [][]rune
	}{
		{"zero", []rune(""), toCombo()},
		{"one", []rune("a"), toCombo("a")},
		{"two", []rune("ab"), toCombo("a", "b", "ab")},
		{"three", []rune("abc"), toCombo(
			"a",
			"b",
			"c",
			"ab",
			"ac",
			"bc",
			"abc",
		)},
		{"one dupe", []rune("aa"), toCombo("a", "aa")},
		{"one dupe plus", []rune("aab"), toCombo(
			"a",
			"b",
			"aa",
			"ab",
			"aab",
		)},
		{"many dupes", []rune("aaaa"), toCombo("a", "aa", "aaa", "aaaa")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Combinations(tt.source); !combosEqual(got, tt.want) {
				t.Errorf("Combinations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func toCombo(ss ...string) [][]rune {
	combos := [][]rune{}

	for _, rs := range ss {
		combos = append(combos, []rune(rs))
	}

	return combos
}

func combosEqual(one, two [][]rune) bool {
	if len(one) != len(two) {
		return false
	}

	return reflect.DeepEqual(toStringSet(one), toStringSet(two))
}

func toStringSet(rss [][]rune) map[string]struct{} {
	uniq := map[string]struct{}{}
	for _, rs := range rss {
		uniq[string(rs)] = struct{}{}
	}

	return uniq
}
