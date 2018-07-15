package runes

import (
	"reflect"
	"sort"
	"testing"
)

func TestSlice_Sort(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   []rune
	}{
		{"empty", "", []rune{}},
		{"one", "a", []rune{'a'}},
		{"ordered", "abc", []rune{'a', 'b', 'c'}},
		{"dupes", "abacbcc", []rune{'a', 'a', 'b', 'b', 'c', 'c', 'c'}},
		{"spaces first", "abc abc abc", []rune{' ', ' ', 'a', 'a', 'a', 'b', 'b', 'b', 'c', 'c', 'c'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := []rune(tt.source)
			sort.Sort(Slice(got))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort(runes.Slice) = %v, want %v", got, tt.want)
			}
		})
	}
}
