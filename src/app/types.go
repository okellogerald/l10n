package app

// disregarded using map becuase it arranges keys in alphabetical order and not in the order in
// which they were inserted
type ARBData []ARBEntry
type ARBEntry struct {
	key   string
	value any
}
