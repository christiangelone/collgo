package maps

import "fmt"

type NotFoundKeyError struct {
	Key string
}

func (e *NotFoundKeyError) Error() string {
	return fmt.Sprintf("NotFoundKeyError: key not found %s", e.Key)
}
