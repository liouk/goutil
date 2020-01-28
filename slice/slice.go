// Package goutil/slice provides slice related operations and utilities
package slice

// Contains checks if a specific element el exists in a string slice sli.
func Contains(sli []string, el string) bool {
	for _, e := range sli {
		if e == el {
			return true
		}
	}

	return false
}

// FindAndRemove removes a specific element el from a slice sli and returns the slice with/without the element and a boolean flag whether it removed it or not.
func FindAndRemove(sli []string, el string) ([]string, bool) {
	for i, e := range sli {
		if e == el {
			return RemoveAtIndex(sli, i), true
		}
	}

	return sli, false
}

// RemoveAtIndex removes an element from a slice sli at the specified index i, and returns the slice without the element, or nil if i was invalid (e.g. -1).
func RemoveAtIndex(sli []string, i int) []string {
	if i < 0 || i >= len(sli) {
		return nil
	}

	return append(sli[:i], sli[i+1:]...)
}
