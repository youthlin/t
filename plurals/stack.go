package plurals

import "errors"

// A simple integer stack

// see antlr.IntStack
type Int64Stack []int64

var ErrEmptyStack = errors.New("Stack is empty")

func (s *Int64Stack) Pop() (int64, error) {
	l := len(*s) - 1
	if l < 0 {
		return 0, ErrEmptyStack
	}
	v := (*s)[l]
	*s = (*s)[0:l]
	return v, nil
}

func (s *Int64Stack) Push(e int64) {
	*s = append(*s, e)
}
