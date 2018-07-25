package runes

// Combinations generates combinations from a source slice of runes.
// Each combination is represented as a slice of runes.
// The null set (empty slice) is omitted from the results.
func Combinations(src []rune) [][]rune {
	// Start with a single entry representing the null set that provides an
	// intial collection to iterate over and build from.
	// This leading nil value is pruned from the final return
	combos := [][]rune{nil}

	for _, r := range src {
		// each added rune effectively doubles the full set of combos
		unchanged := make([][]rune, 0, len(combos))
		runeAdded := make([][]rune, 0, len(combos))
		for _, rs := range combos {
			// new combo with added rune appended
			newCombo := make([]rune, len(rs)+1)
			copy(newCombo, rs)
			newCombo[len(newCombo)-1] = r

			// Handle the case of a repeated rune in the src set.  The goal is to avoid
			// adding duplicate combos to the master result.
			// Identifying repeats of 'r' is straight forward assuming the 'src' is ordered making
			// repeats sequential.
			// The thing to note is a combo made in previous iteration will once again be
			// made in the current iteration in such a case of a repeat 'r' -- i.e. it's safe
			// to drop the old because it's replaced by a new duplicate.
			// So:
			// - identify the repeat case
			// - keep only one of the duplicate combos
			//
			if !(len(rs) > 0 && rs[len(rs)-1] == r) {
				unchanged = append(unchanged, rs)
			}
			runeAdded = append(runeAdded, newCombo)
		}
		combos = append(unchanged, runeAdded...)
	}

	// omit first value (nil)
	return combos[1:]
}
