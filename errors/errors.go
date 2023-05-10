package errors

import "fmt"

func Wrap(msg botError, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %w", msg, err)
	}
	return nil
}
