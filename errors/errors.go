package errors

import (
	"errors"
	"fmt"
)

func Errorf(format string, args ...any) error {
	return fmt.Errorf(format, args...)
}

func Wrapf(err error, format string, args ...any) error {
	args = append(args, err)
	return fmt.Errorf(format+": %w", args...)
}

func WithSecondaryError(err error, additionalErr error) error {
	return fmt.Errorf("%w: %w", err, additionalErr)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}
