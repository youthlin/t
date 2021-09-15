package plurals

import "errors"

// A simple integer stack

// Int64Stack see antlr.IntStack
type Int64Stack []int64

// ErrEmptyStack returned when call Pop on empty stack
var ErrEmptyStack = errors.New("stack is empty")

// Pop pop a number from stack
func (s *Int64Stack) Pop() (int64, error) {
	l := len(*s) - 1
	if l < 0 {
		return 0, ErrEmptyStack
	}
	v := (*s)[l]
	*s = (*s)[0:l]
	return v, nil
}

// Push push a number to stack
func (s *Int64Stack) Push(e int64) {
	*s = append(*s, e)
}
