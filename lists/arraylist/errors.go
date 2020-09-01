package arraylist

import "fmt"

type OutOfBoundError struct {
	index uint64
}

func (e *OutOfBoundError) Error() string {
	return fmt.Sprintf("OutOfBoundError: element not found at index %d", e.index)
}