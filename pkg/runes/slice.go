package runes

// Slice attaches the methods of sort.Interface to []rune, sorting in increasing order.
type Slice []rune

func (rs Slice) Len() int               { return len(rs) }
func (rs Slice) Less(i int, j int) bool { return rs[i] < rs[j] }
func (rs Slice) Swap(i int, j int)      { rs[i], rs[j] = rs[j], rs[i] }
