package wp

import "strings"

// StringRack is a simple string implementation of the Rack interface.
type StringRack struct {
	chars string
}

// NewStringRack contructs a StringRack.
func NewStringRack(chars string) *StringRack {
	return &StringRack{chars}
}

func (r *StringRack) String() string {
	cs := strings.Split(r.chars, "")
	return strings.Join(cs, " ")
}

// Chars returns all the characters in the Rack with possible duplicates.
func (r *StringRack) Chars() []rune {
	return []rune(r.chars)
}

// CharsDistinct returns all the unique characters in the Rack.
func (r *StringRack) CharsDistinct() []rune {
	charCount := make(map[rune]int)
	for _, c := range r.chars {
		charCount[c]++
	}

	unique := make([]rune, len(charCount))
	for k := range charCount {
		unique = append(unique, k)
	}

	return unique
}

// CanMakeWord tests if a word can be formed using the Rack characters.
func (r *StringRack) CanMakeWord(word string) bool {
	availChars := r.chars

	for _, c := range word {
		if i := strings.IndexRune(availChars, c); i > -1 {
			availChars = availChars[:i] + availChars[i+1:] // remove char from avail pool
		} else {
			return false
		}
	}

	return true
}
