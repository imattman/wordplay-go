package runes

import (
	"reflect"
	"testing"
)

func TestMultiset_Counts(t *testing.T) {
	tests := []struct {
		name   string
		source []rune
		want   map[rune]int
	}{
		{"zero", []rune(""), map[rune]int{}},
		{"one", []rune("a"), map[rune]int{'a': 1}},
		{"aaa", []rune("aaa"), map[rune]int{'a': 3}},
		{"abbccc", []rune("abbccc"), map[rune]int{'a': 1, 'b': 2, 'c': 3}},
		{"cabcbc", []rune("cabcbc"), map[rune]int{'a': 1, 'b': 2, 'c': 3}},
		{"aa bbb cccc", []rune("aa bbb cccc"), map[rune]int{' ': 2, 'a': 2, 'b': 3, 'c': 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := NewMultiset(tt.source)
			if got := ms.Counts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiset.Counts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiset_Slice(t *testing.T) {
	tests := []struct {
		name   string
		source []rune
		want   []rune
	}{
		{"zero", []rune(""), []rune{}},
		{"one", []rune("a"), []rune{'a'}},
		{"aaa", []rune("aaa"), []rune{'a', 'a', 'a'}},
		{"abbccc", []rune("abbccc"), []rune{'a', 'b', 'b', 'c', 'c', 'c'}},
		{"cabcbc", []rune("cabcbc"), []rune{'a', 'b', 'b', 'c', 'c', 'c'}},
		{"aa bbb cccc", []rune("aa bbb cccc"), []rune{' ', ' ', 'a', 'a', 'b', 'b', 'b', 'c', 'c', 'c', 'c'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := NewMultiset(tt.source)
			if got := ms.Slice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiset.Counts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiset_ContainsAll(t *testing.T) {
	tests := []struct {
		name  string
		first string
		other string
		want  bool
	}{
		{"compare empty", "", "", true},
		{"compare same", "aaa", "aaa", true},
		{"a vs empty", "a", "", true},
		{"empty vs a", "", "a", false},
		{"superset letters", "aabbccd", "aabbcc", true},
		{"superset count", "aaabbb", "aabb", true},
		{"insufficient letters", "abc", "abcz", false},
		{"insufficient count", "ab", "aabb", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := NewMultiset([]rune(tt.first))
			other := NewMultiset([]rune(tt.other))
			if got := ms.ContainsAll(other); got != tt.want {
				t.Errorf("Multiset.ContainsAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiset_Equals(t *testing.T) {
	tests := []struct {
		name  string
		first string
		other string
		want  bool
	}{
		{"compare empty", "", "", true},
		{"compare same", "aaa", "aaa", true},
		{"a vs empty", "a", "", false},
		{"empty vs a", "", "a", false},
		{"superset letters", "aabbccd", "aabbcc", false},
		{"superset count", "aaabbb", "aabb", false},
		{"insufficient letters", "abc", "abcz", false},
		{"insufficient count", "ab", "aabb", false},
		{"different order", "aabbccddee", "abcdeabcde", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := NewMultiset([]rune(tt.first))
			other := NewMultiset([]rune(tt.other))
			if got := ms.Equals(other); got != tt.want {
				t.Errorf("Multiset.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiset_String(t *testing.T) {
	tests := []struct {
		name   string
		source []rune
		want   string
	}{
		{"zero", []rune(""), "[]"},
		{"one", []rune("a"), "[a]"},
		{"aaa", []rune("aaa"), "[a,a,a]"},
		{"ababa", []rune("ababa"), "[a,a,a,b,b]"},
		{"aa zzz", []rune("aa zzz"), "[ ,a,a,z,z,z]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := NewMultiset(tt.source)
			if got := ms.String(); got != tt.want {
				t.Errorf("Multiset.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
