package parse

import "sync"

var cache sync.Map

func Eval(input string, n int64) int64 {
	if val, ok := cache.Load(input); ok {
		if f, ok := val.(Fn); ok {
			return f(n)
		}
	}
	f, err := String(input)
	if err != nil {
		f = zero
	}
	cache.Store(input, f)
	return f(n)
}

type Fn func(n int64) int64

func zero(n int64) int64 { return 0 }

func String(input string) (Fn, error) {
	tokens, err := Lex(input)
	if err != nil {
		return zero, nil
	}
	return Tokens(tokens)
}

func Tokens(input []*Token) (Fn, error) {
	tree, err := ParseToken(input)
	if err != nil {
		return zero, nil
	}
	return Ast(tree)
}

func Ast(tree *Node) (Fn, error) {
	return zero, nil
}
