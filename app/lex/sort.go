package lex

// RuneSlice implements sort.Interface for a slice of runes.
type RuneSlice []rune

// Required by sort.Interface
func (rs RuneSlice) Len() int {
	return len(rs)
}

// Required by sort.Interface
func (rs RuneSlice) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

// Required by sort.Interface
func (rs RuneSlice) Less(i, j int) bool {
	return rs[i] < rs[j]
}

// ByScore implements sort.Interface for []*Match based on word score
type ByScore []*Match

// Required by sort.Interface
func (m ByScore) Len() int {
	return len(m)
}

// Required by sort.Interface
func (m ByScore) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// Required by sort.Interface
func (m ByScore) Less(i, j int) bool {
	return m[i].Score < m[j].Score
}
