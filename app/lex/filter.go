package lex

// PartitioningFilter defines the interface for reducing the set of the Lexicon entries
// against which Word matches are made.  This also provides an opportunity to
// subdivide the Lexicon for processing with current goroutines.
type PartitioningFilter interface {
	Filter(l Lexicon, r Rack) ([]Lexicon, error)
}

// PartitioningFilterFunc is a type adapter to allow the use of an ordinary function
// as a PartitioningFilter
type PartitioningFilterFunc func(l Lexicon, r Rack) ([]Lexicon, error)

// Filter calls fn(l, r)
func (fn PartitioningFilterFunc) Filter(l Lexicon, r Rack) ([]Lexicon, error) {
	return fn(l, r)
}

// NoopFilter is PartitioningFilter that does no filtering or partitioning.
var NoopFilter = PartitioningFilterFunc(
	func(lex Lexicon, r Rack) ([]Lexicon, error) {
		return []Lexicon{lex}, nil
	})

// PrePartitionByFirstChar builds a Lexicon filter where the underlying
// Lexicon is partitioned into groups based on the first character of the word.
func PrePartitionByFirstChar(lex Lexicon) PartitioningFilter {
	byFirstChar := make(map[rune]Lexicon)
	for _, word := range lex {
		fc := []rune(word.String())[0]
		byFirstChar[fc] = append(byFirstChar[fc], word)
	}
	// fmt.Printf("byFirstChar(%d)\n", len(byFirstChar))

	fn := func(_ Lexicon, r Rack) ([]Lexicon, error) {
		ls := make([]Lexicon, 0, len(r.Unique()))

		for _, char := range r.Unique() {
			ls = append(ls, byFirstChar[char])
		}
		return ls, nil
	}

	return PartitioningFilterFunc(fn)
}
