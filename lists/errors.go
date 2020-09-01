package lists

import "fmt"

type RemovingInEmptinessError struct {}

func (e *RemovingInEmptinessError) Error() string {
	return "RemovingInEmptyError: trying to remove element in emptiness"
}

type FetchingInEmptinessError struct {}

func (e *FetchingInEmptinessError) Error() string {
	return "FetchingInEmptinessError: trying to fetch element in emptiness"
}

type ElementTypeError struct {
	Expected interface{}
	Actual   interface{}
}

func (e *ElementTypeError) Error() string {
	return fmt.Sprintf(
		"ElementTypeError: trying to insert element of different type (expect: %T, actual: %T)",
		e.Expected,
		e.Actual,
	)
}

type DifferentSizeError struct {
	Expected interface{}
	Actual interface{}
}

func (e *DifferentSizeError) Error() string {
	return fmt.Sprintf(
		"DifferentSizeError: collection has different size (expected: %d, actual: %d)",
		e.Expected,
		e.Actual,
	)
}
